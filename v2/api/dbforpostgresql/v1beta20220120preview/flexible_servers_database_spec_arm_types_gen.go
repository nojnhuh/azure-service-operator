// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20220120preview

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of FlexibleServers_Database_Spec. Use v1api20220120preview.FlexibleServers_Database_Spec instead
type FlexibleServers_Database_Spec_ARM struct {
	Name       string                  `json:"name,omitempty"`
	Properties *DatabaseProperties_ARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &FlexibleServers_Database_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-01-20-preview"
func (database FlexibleServers_Database_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (database *FlexibleServers_Database_Spec_ARM) GetName() string {
	return database.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforPostgreSQL/flexibleServers/databases"
func (database *FlexibleServers_Database_Spec_ARM) GetType() string {
	return "Microsoft.DBforPostgreSQL/flexibleServers/databases"
}

// Deprecated version of DatabaseProperties. Use v1api20220120preview.DatabaseProperties instead
type DatabaseProperties_ARM struct {
	Charset   *string `json:"charset,omitempty"`
	Collation *string `json:"collation,omitempty"`
}