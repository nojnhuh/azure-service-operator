// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20211101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type NamespacesEventhubsConsumergroups_SPECARM struct {
	AzureName string `json:"azureName"`
	Name      string `json:"name"`

	//Properties: Single item in List or Get Consumer group operation
	Properties *NamespacesEventhubsConsumergroups_SPEC_PropertiesARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &NamespacesEventhubsConsumergroups_SPECARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (specarm NamespacesEventhubsConsumergroups_SPECARM) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetName returns the Name of the resource
func (specarm NamespacesEventhubsConsumergroups_SPECARM) GetName() string {
	return specarm.Name
}

// GetType returns the ARM Type of the resource. This is always ""
func (specarm NamespacesEventhubsConsumergroups_SPECARM) GetType() string {
	return ""
}

type NamespacesEventhubsConsumergroups_SPEC_PropertiesARM struct {
	//UserMetadata: User Metadata is a placeholder to store user-defined string data
	//with maximum length 1024. e.g. it can be used to store descriptive data, such as
	//list of teams and their contact information also user-defined configuration
	//settings can be stored.
	UserMetadata *string `json:"userMetadata,omitempty"`
}
