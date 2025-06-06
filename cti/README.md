## Create virtual environment
```
python3 -m venv ctivenv
```

## Activate virtual environment
```
source ctivenv/bin/activate
```

## Deactivate virtual environment
```
deactivate
```

## Install dependecies
```
pip install -r requirements.txt
```

## Update dependecies
```
pip freeze > requirements.txt
```

## Delete all dependencies
```
pip freeze | xargs pip uninstall -y
```

## Remove the virtual environment
```
sudo rm -rf ctivenv
```

## Chaincode Usage

For instructions on deploying and invoking the chaincode for CTI transfer, see [CHAINCODE_USAGE.md](./CHAINCODE_USAGE.md).

## Organization Access List

| **Function**    | **Head<br>of<br>Operations** | **Intelligence<br>Unit** | **Tactical<br>Unit** | **Special<br>Operations<br>Unit** |
| -------------- | ---------------------------- | ------------------------ | -------------------- | --------------------------------- |
| Create CTI     | ✓                            | ✓                        | ✗                    | ✓                                 |
| Read CTI | ✓                            | (✓)                      | ✓                    | ✓                                 |
| Read all CTI | ✓                            | (✓)                      | (✓)                  | ✓                                 |
| Update CTI    | ✓                            | (✓)                      | ✗                    | (✓)                               |
| Delete CTI     | ✓                            | ✗                        | ✗                    | ✗                                 |
