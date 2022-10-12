// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20180601

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Servers_Configuration_Spec_ARM struct {
	// Location: Location to deploy resource to
	Location *string `json:"location,omitempty"`

	// Name: The name of the server configuration.
	Name string `json:"name,omitempty"`

	// Properties: The properties of a configuration.
	Properties *ConfigurationProperties_ARM `json:"properties,omitempty"`

	// Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Servers_Configuration_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2018-06-01"
func (configuration Servers_Configuration_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (configuration *Servers_Configuration_Spec_ARM) GetName() string {
	return configuration.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforMariaDB/servers/configurations"
func (configuration *Servers_Configuration_Spec_ARM) GetType() string {
	return "Microsoft.DBforMariaDB/servers/configurations"
}

// Generated from: https://schema.management.azure.com/schemas/2018-06-01/Microsoft.DBforMariaDB.json#/definitions/ConfigurationProperties
type ConfigurationProperties_ARM struct {
	// Source: Source of the configuration.
	Source *string `json:"source,omitempty"`

	// Value: Value of the configuration.
	Value *string `json:"value,omitempty"`
}