/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// ChaincodePackageApplyConfiguration represents a declarative configuration of the ChaincodePackage type for use
// with apply.
type ChaincodePackageApplyConfiguration struct {
	Name        *string                                `json:"name,omitempty"`
	Address     *string                                `json:"address,omitempty"`
	Type        *string                                `json:"type,omitempty"`
	DialTimeout *string                                `json:"dialTimeout,omitempty"`
	TLS         *ChaincodePackageTLSApplyConfiguration `json:"tls,omitempty"`
}

// ChaincodePackageApplyConfiguration constructs a declarative configuration of the ChaincodePackage type for use with
// apply.
func ChaincodePackage() *ChaincodePackageApplyConfiguration {
	return &ChaincodePackageApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *ChaincodePackageApplyConfiguration) WithName(value string) *ChaincodePackageApplyConfiguration {
	b.Name = &value
	return b
}

// WithAddress sets the Address field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Address field is set to the value of the last call.
func (b *ChaincodePackageApplyConfiguration) WithAddress(value string) *ChaincodePackageApplyConfiguration {
	b.Address = &value
	return b
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *ChaincodePackageApplyConfiguration) WithType(value string) *ChaincodePackageApplyConfiguration {
	b.Type = &value
	return b
}

// WithDialTimeout sets the DialTimeout field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DialTimeout field is set to the value of the last call.
func (b *ChaincodePackageApplyConfiguration) WithDialTimeout(value string) *ChaincodePackageApplyConfiguration {
	b.DialTimeout = &value
	return b
}

// WithTLS sets the TLS field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TLS field is set to the value of the last call.
func (b *ChaincodePackageApplyConfiguration) WithTLS(value *ChaincodePackageTLSApplyConfiguration) *ChaincodePackageApplyConfiguration {
	b.TLS = value
	return b
}
