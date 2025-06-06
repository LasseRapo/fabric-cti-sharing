/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// FabricNetworkConfigIdentityApplyConfiguration represents a declarative configuration of the FabricNetworkConfigIdentity type for use
// with apply.
type FabricNetworkConfigIdentityApplyConfiguration struct {
	Name      *string `json:"name,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
}

// FabricNetworkConfigIdentityApplyConfiguration constructs a declarative configuration of the FabricNetworkConfigIdentity type for use with
// apply.
func FabricNetworkConfigIdentity() *FabricNetworkConfigIdentityApplyConfiguration {
	return &FabricNetworkConfigIdentityApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *FabricNetworkConfigIdentityApplyConfiguration) WithName(value string) *FabricNetworkConfigIdentityApplyConfiguration {
	b.Name = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *FabricNetworkConfigIdentityApplyConfiguration) WithNamespace(value string) *FabricNetworkConfigIdentityApplyConfiguration {
	b.Namespace = &value
	return b
}
