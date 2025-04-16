/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// SecretRefApplyConfiguration represents a declarative configuration of the SecretRef type for use
// with apply.
type SecretRefApplyConfiguration struct {
	Name *string `json:"name,omitempty"`
}

// SecretRefApplyConfiguration constructs a declarative configuration of the SecretRef type for use with
// apply.
func SecretRef() *SecretRefApplyConfiguration {
	return &SecretRefApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *SecretRefApplyConfiguration) WithName(value string) *SecretRefApplyConfiguration {
	b.Name = &value
	return b
}
