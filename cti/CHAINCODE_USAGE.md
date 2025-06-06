## Configure internal DNS
```bash
kubectl apply -f - <<EOF
kind: ConfigMap
apiVersion: v1
metadata:
  name: coredns
  namespace: kube-system
data:
  Corefile: |
    .:53 {
        errors
        health {
           lameduck 5s
        }
        rewrite name regex (.*)\.localho\.st istio-ingressgateway.istio-system.svc.cluster.local
        hosts {
          fallthrough
        }
        ready
        kubernetes cluster.local in-addr.arpa ip6.arpa {
           pods insecure
           fallthrough in-addr.arpa ip6.arpa
           ttl 30
        }
        prometheus :9153
        forward . /etc/resolv.conf {
           max_concurrent 1000
        }
        cache 30
        loop
        reload
        loadbalance
    }
EOF
```


## Generate network config
```bash
kubectl hlf inspect -c=main --output resources/network.yaml -o Org1MSP -o Org2MSP -o Org3MSP -o Org4MSP -o OrdererMSP
```
```bash
kubectl hlf inspect -c=private --output resources/privatenetwork.yaml -o Org1MSP -o Org4MSP -o OrdererMSP
```


```bash
kubectl hlf utils adduser --userPath=resources/org1msp.yaml --config=resources/network.yaml --username=admin --mspid=Org1MSP
kubectl hlf utils adduser --userPath=resources/org2msp.yaml --config=resources/network.yaml --username=admin --mspid=Org2MSP
kubectl hlf utils adduser --userPath=resources/org3msp.yaml --config=resources/network.yaml --username=admin --mspid=Org3MSP
kubectl hlf utils adduser --userPath=resources/org4msp.yaml --config=resources/network.yaml --username=admin --mspid=Org4MSP
```

```bash
kubectl hlf utils adduser --userPath=resources/org1msp.yaml --config=resources/privatenetwork.yaml --username=admin --mspid=Org1MSP
kubectl hlf utils adduser --userPath=resources/org4msp.yaml --config=resources/privatenetwork.yaml --username=admin --mspid=Org4MSP
```


## Create metadata file
```bash
rm code.tar.gz chaincode.tgz
export CHAINCODE_NAME=ctitransfer111
export CHAINCODE_LABEL=ctitransfer111
cat << METADATA-EOF > "metadata.json"
{
    "type": "ccaas",
    "label": "${CHAINCODE_LABEL}"
}
METADATA-EOF
```


## Prepare connection file
```bash
cat > "connection.json" <<CONN_EOF
{
  "address": "${CHAINCODE_NAME}:7052",
  "dial_timeout": "10s",
  "tls_required": false
}
CONN_EOF

tar cfz code.tar.gz connection.json
tar cfz chaincode.tgz metadata.json code.tar.gz
export PACKAGE_ID=$(kubectl hlf chaincode calculatepackageid --path=chaincode.tgz --language=node --label=$CHAINCODE_LABEL)
echo "PACKAGE_ID=$PACKAGE_ID"

kubectl hlf chaincode install --path=./chaincode.tgz \
    --config=resources/network.yaml --language=golang --label=$CHAINCODE_LABEL --user=admin --peer=org1-peer0.default

kubectl hlf chaincode install --path=./chaincode.tgz \
    --config=resources/network.yaml --language=golang --label=$CHAINCODE_LABEL --user=admin --peer=org1-peer1.default

kubectl hlf chaincode install --path=./chaincode.tgz \
    --config=resources/network.yaml --language=golang --label=$CHAINCODE_LABEL --user=admin --peer=org2-peer0.default

kubectl hlf chaincode install --path=./chaincode.tgz \
    --config=resources/network.yaml --language=golang --label=$CHAINCODE_LABEL --user=admin --peer=org2-peer1.default

kubectl hlf chaincode install --path=./chaincode.tgz \
    --config=resources/network.yaml --language=golang --label=$CHAINCODE_LABEL --user=admin --peer=org3-peer0.default

kubectl hlf chaincode install --path=./chaincode.tgz \
    --config=resources/network.yaml --language=golang --label=$CHAINCODE_LABEL --user=admin --peer=org3-peer1.default

kubectl hlf chaincode install --path=./chaincode.tgz \
    --config=resources/network.yaml --language=golang --label=$CHAINCODE_LABEL --user=admin --peer=org4-peer0.default

kubectl hlf chaincode install --path=./chaincode.tgz \
    --config=resources/network.yaml --language=golang --label=$CHAINCODE_LABEL --user=admin --peer=org4-peer1.default

kubectl hlf chaincode install --path=./chaincode.tgz \
    --config=resources/privatenetwork.yaml --language=golang --label=$CHAINCODE_LABEL --user=admin --peer=org1-peer0.default

kubectl hlf chaincode install --path=./chaincode.tgz \
    --config=resources/privatenetwork.yaml --language=golang --label=$CHAINCODE_LABEL --user=admin --peer=org1-peer1.default

kubectl hlf chaincode install --path=./chaincode.tgz \
    --config=resources/privatenetwork.yaml --language=golang --label=$CHAINCODE_LABEL --user=admin --peer=org4-peer0.default

kubectl hlf chaincode install --path=./chaincode.tgz \
    --config=resources/privatenetwork.yaml --language=golang --label=$CHAINCODE_LABEL --user=admin --peer=org4-peer1.default
```


## Check if the chaincode is installed
```bash
kubectl hlf chaincode queryinstalled --config=resources/network.yaml --user=admin --peer=org1-peer0.default
```


## Deploy chaincode container on cluster
```bash
kubectl hlf externalchaincode sync --image=betoni/cti-transfer:v1.1.1 \
    --name=$CHAINCODE_NAME \
    --namespace=default \
    --package-id=$PACKAGE_ID \
    --tls-required=false \
    --replicas=1
```


## Approve chaincode Org1MSP
```bash
export SEQUENCE=1
export VERSION="1.0"
kubectl hlf chaincode approveformyorg --config=resources/network.yaml --user=admin --peer=org1-peer0.default \
    --package-id=$PACKAGE_ID \
    --version "$VERSION" --sequence "$SEQUENCE" --name=$CHAINCODE_NAME \
    --policy="AND('Org1MSP.member', 'Org4MSP.member')" --channel=main
```

## Approve chaincode Org2MSP
```bash
kubectl hlf chaincode approveformyorg --config=resources/network.yaml --user=admin --peer=org2-peer0.default \
    --package-id=$PACKAGE_ID \
    --version "$VERSION" --sequence "$SEQUENCE" --name=$CHAINCODE_NAME \
    --policy="AND('Org1MSP.member', 'Org4MSP.member')" --channel=main
```

## Approve chaincode Org3MSP
```bash
kubectl hlf chaincode approveformyorg --config=resources/network.yaml --user=admin --peer=org3-peer0.default \
    --package-id=$PACKAGE_ID \
    --version "$VERSION" --sequence "$SEQUENCE" --name=$CHAINCODE_NAME \
    --policy="AND('Org1MSP.member', 'Org4MSP.member')" --channel=main
```

## Approve chaincode Org4MSP
```bash
kubectl hlf chaincode approveformyorg --config=resources/network.yaml --user=admin --peer=org4-peer0.default \
    --package-id=$PACKAGE_ID \
    --version "$VERSION" --sequence "$SEQUENCE" --name=$CHAINCODE_NAME \
    --policy="AND('Org1MSP.member', 'Org4MSP.member')" --channel=main
```

## Approve chaincode Org1MSP - Private
```bash
export SEQUENCE=1
export VERSION="1.0"
kubectl hlf chaincode approveformyorg --config=resources/privatenetwork.yaml --user=admin --peer=org1-peer0.default \
    --package-id=$PACKAGE_ID \
    --version "$VERSION" --sequence "$SEQUENCE" --name=$CHAINCODE_NAME \
    --policy="AND('Org1MSP.member', 'Org4MSP.member')" --channel=private
```

## Approve chaincode Org4MSP - Private
```bash
kubectl hlf chaincode approveformyorg --config=resources/privatenetwork.yaml --user=admin --peer=org4-peer0.default \
    --package-id=$PACKAGE_ID \
    --version "$VERSION" --sequence "$SEQUENCE" --name=$CHAINCODE_NAME \
    --policy="AND('Org1MSP.member', 'Org4MSP.member')" --channel=private
```


## Commit chaincode
```bash
kubectl hlf chaincode commit --config=resources/network.yaml --user=admin --mspid=Org1MSP \
    --version "$VERSION" --sequence "$SEQUENCE" --name=$CHAINCODE_NAME \
    --policy="AND('Org1MSP.member', 'Org4MSP.member')" --channel=main
```

```bash
kubectl hlf chaincode commit --config=resources/privatenetwork.yaml --user=admin --mspid=Org1MSP \
    --version "$VERSION" --sequence "$SEQUENCE" --name=$CHAINCODE_NAME \
    --policy="AND('Org1MSP.member', 'Org4MSP.member')" --channel=private
```


## Initialize the chaincode with sample CTI metadata
```bash
kubectl hlf chaincode invoke --config=resources/network.yaml \
    --user=admin --peer=org1-peer0.default \
    --chaincode=$CHAINCODE_NAME --channel=main \
    --fcn=initLedger
```

## Query all CTI metadata
```bash
kubectl hlf chaincode invoke --config=resources/network.yaml \
    --user=admin --peer=org1-peer0.default \
    --chaincode=$CHAINCODE_NAME --channel=main \
    --fcn=GetAllCTI \
    --args='2000' --args='""'
```


## Add new CTI metadata
```bash
kubectl hlf chaincode invoke --config=resources/network.yaml \
    --user=admin --peer=org2-peer0.default \
    --chaincode=$CHAINCODE_NAME --channel=main \
    --fcn=CreateCTIMetadata \
    --args="{\"UUID\":\"55555\",\"Description\":\"Example metadata 5555\",\"Timestamp\":\"2025-04-30T12:00:00Z\",\"SenderIdentity\":\"IntelligenceUnit\",\"CID\":\"CID5555\",\"VaultKey\":\"vaultKey5555\",\"SHA256Hash\":\"sha256hash5555\",\"AccessList\":[\"HeadOfOperations\",\"IntelligenceUnit\"]}"
```


## Query CTI metadata by its UUID
```bash
kubectl hlf chaincode invoke --config=resources/network.yaml \
    --user=admin --peer=org1-peer0.default \
    --chaincode=$CHAINCODE_NAME --channel=main \
    --fcn=ReadCTIMetadata \
    --args="12345"
```


## Update an existing CTI metadata entry
```bash
kubectl hlf chaincode invoke --config=resources/network.yaml \
    --user=admin --peer=org1-peer0.default \
    --chaincode=$CHAINCODE_NAME --channel=main \
    --fcn=UpdateCTIMetadata \
    --args="{\"UUID\":\"12345\",\"Description\":\"Updated metadata entry 12345\",\"Timestamp\":\"2025-05-01T12:00:00Z\",\"SenderIdentity\":\"HeadOfOperations\",\"CID\":\"CID12345-updated\",\"VaultKey\":\"vaultKey12345-updated\",\"SHA256Hash\":\"sha256hash12345-updated\",\"AccessList\":[\"HeadOfOperations\",\"IntelligenceUnit\",\"TacticalUnit\"]}"
```


## Delete CTI metadata by its UUID
```bash
kubectl hlf chaincode invoke --config=resources/network.yaml \
    --user=admin --peer=org1-peer0.default \
    --chaincode=$CHAINCODE_NAME --channel=main \
    --fcn=DeleteCTIMetadata \
    --args="12345"
```