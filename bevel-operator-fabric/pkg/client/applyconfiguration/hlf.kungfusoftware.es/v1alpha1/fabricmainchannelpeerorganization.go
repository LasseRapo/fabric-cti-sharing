/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// FabricMainChannelPeerOrganizationApplyConfiguration represents a declarative configuration of the FabricMainChannelPeerOrganization type for use
// with apply.
type FabricMainChannelPeerOrganizationApplyConfiguration struct {
	MSPID       *string `json:"mspID,omitempty"`
	CAName      *string `json:"caName,omitempty"`
	CANamespace *string `json:"caNamespace,omitempty"`
}

// FabricMainChannelPeerOrganizationApplyConfiguration constructs a declarative configuration of the FabricMainChannelPeerOrganization type for use with
// apply.
func FabricMainChannelPeerOrganization() *FabricMainChannelPeerOrganizationApplyConfiguration {
	return &FabricMainChannelPeerOrganizationApplyConfiguration{}
}

// WithMSPID sets the MSPID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MSPID field is set to the value of the last call.
func (b *FabricMainChannelPeerOrganizationApplyConfiguration) WithMSPID(value string) *FabricMainChannelPeerOrganizationApplyConfiguration {
	b.MSPID = &value
	return b
}

// WithCAName sets the CAName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CAName field is set to the value of the last call.
func (b *FabricMainChannelPeerOrganizationApplyConfiguration) WithCAName(value string) *FabricMainChannelPeerOrganizationApplyConfiguration {
	b.CAName = &value
	return b
}

// WithCANamespace sets the CANamespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CANamespace field is set to the value of the last call.
func (b *FabricMainChannelPeerOrganizationApplyConfiguration) WithCANamespace(value string) *FabricMainChannelPeerOrganizationApplyConfiguration {
	b.CANamespace = &value
	return b
}
