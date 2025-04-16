/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1alpha1 "github.com/kfsoftware/hlf-operator/pkg/apis/hlf.kungfusoftware.es/v1alpha1"
	hlfkungfusoftwareesv1alpha1 "github.com/kfsoftware/hlf-operator/pkg/client/applyconfiguration/hlf.kungfusoftware.es/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeFabricPeers implements FabricPeerInterface
type FakeFabricPeers struct {
	Fake *FakeHlfV1alpha1
	ns   string
}

var fabricpeersResource = v1alpha1.SchemeGroupVersion.WithResource("fabricpeers")

var fabricpeersKind = v1alpha1.SchemeGroupVersion.WithKind("FabricPeer")

// Get takes name of the fabricPeer, and returns the corresponding fabricPeer object, and an error if there is any.
func (c *FakeFabricPeers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.FabricPeer, err error) {
	emptyResult := &v1alpha1.FabricPeer{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(fabricpeersResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.FabricPeer), err
}

// List takes label and field selectors, and returns the list of FabricPeers that match those selectors.
func (c *FakeFabricPeers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.FabricPeerList, err error) {
	emptyResult := &v1alpha1.FabricPeerList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(fabricpeersResource, fabricpeersKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FabricPeerList{ListMeta: obj.(*v1alpha1.FabricPeerList).ListMeta}
	for _, item := range obj.(*v1alpha1.FabricPeerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested fabricPeers.
func (c *FakeFabricPeers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(fabricpeersResource, c.ns, opts))

}

// Create takes the representation of a fabricPeer and creates it.  Returns the server's representation of the fabricPeer, and an error, if there is any.
func (c *FakeFabricPeers) Create(ctx context.Context, fabricPeer *v1alpha1.FabricPeer, opts v1.CreateOptions) (result *v1alpha1.FabricPeer, err error) {
	emptyResult := &v1alpha1.FabricPeer{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(fabricpeersResource, c.ns, fabricPeer, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.FabricPeer), err
}

// Update takes the representation of a fabricPeer and updates it. Returns the server's representation of the fabricPeer, and an error, if there is any.
func (c *FakeFabricPeers) Update(ctx context.Context, fabricPeer *v1alpha1.FabricPeer, opts v1.UpdateOptions) (result *v1alpha1.FabricPeer, err error) {
	emptyResult := &v1alpha1.FabricPeer{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(fabricpeersResource, c.ns, fabricPeer, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.FabricPeer), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFabricPeers) UpdateStatus(ctx context.Context, fabricPeer *v1alpha1.FabricPeer, opts v1.UpdateOptions) (result *v1alpha1.FabricPeer, err error) {
	emptyResult := &v1alpha1.FabricPeer{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(fabricpeersResource, "status", c.ns, fabricPeer, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.FabricPeer), err
}

// Delete takes name of the fabricPeer and deletes it. Returns an error if one occurs.
func (c *FakeFabricPeers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(fabricpeersResource, c.ns, name, opts), &v1alpha1.FabricPeer{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFabricPeers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(fabricpeersResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.FabricPeerList{})
	return err
}

// Patch applies the patch and returns the patched fabricPeer.
func (c *FakeFabricPeers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.FabricPeer, err error) {
	emptyResult := &v1alpha1.FabricPeer{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(fabricpeersResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.FabricPeer), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied fabricPeer.
func (c *FakeFabricPeers) Apply(ctx context.Context, fabricPeer *hlfkungfusoftwareesv1alpha1.FabricPeerApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.FabricPeer, err error) {
	if fabricPeer == nil {
		return nil, fmt.Errorf("fabricPeer provided to Apply must not be nil")
	}
	data, err := json.Marshal(fabricPeer)
	if err != nil {
		return nil, err
	}
	name := fabricPeer.Name
	if name == nil {
		return nil, fmt.Errorf("fabricPeer.Name must be provided to Apply")
	}
	emptyResult := &v1alpha1.FabricPeer{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(fabricpeersResource, c.ns, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.FabricPeer), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeFabricPeers) ApplyStatus(ctx context.Context, fabricPeer *hlfkungfusoftwareesv1alpha1.FabricPeerApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.FabricPeer, err error) {
	if fabricPeer == nil {
		return nil, fmt.Errorf("fabricPeer provided to Apply must not be nil")
	}
	data, err := json.Marshal(fabricPeer)
	if err != nil {
		return nil, err
	}
	name := fabricPeer.Name
	if name == nil {
		return nil, fmt.Errorf("fabricPeer.Name must be provided to Apply")
	}
	emptyResult := &v1alpha1.FabricPeer{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(fabricpeersResource, c.ns, *name, types.ApplyPatchType, data, opts.ToPatchOptions(), "status"), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.FabricPeer), err
}
