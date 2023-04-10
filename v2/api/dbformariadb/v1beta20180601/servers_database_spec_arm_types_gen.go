// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20180601

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of Servers_Database_Spec. Use v1api20180601.Servers_Database_Spec instead
type Servers_Database_Spec_ARM struct {
	Name       string                  `json:"name,omitempty"`
	Properties *DatabaseProperties_ARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Servers_Database_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2018-06-01"
func (database Servers_Database_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (database *Servers_Database_Spec_ARM) GetName() string {
	return database.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforMariaDB/servers/databases"
func (database *Servers_Database_Spec_ARM) GetType() string {
	return "Microsoft.DBforMariaDB/servers/databases"
}

// Deprecated version of DatabaseProperties. Use v1api20180601.DatabaseProperties instead
type DatabaseProperties_ARM struct {
	Charset   *string `json:"charset,omitempty"`
	Collation *string `json:"collation,omitempty"`
}