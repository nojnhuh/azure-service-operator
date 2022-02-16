// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201201

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type RedisLinkedServer_SpecARM struct {
	AzureName string `json:"azureName"`
	Name      string `json:"name"`

	//Properties: Properties required to create a linked server.
	Properties RedisLinkedServerPropertiesARM `json:"properties"`
}

var _ genruntime.ARMResourceSpec = &RedisLinkedServer_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-12-01"
func (server RedisLinkedServer_SpecARM) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetName returns the Name of the resource
func (server RedisLinkedServer_SpecARM) GetName() string {
	return server.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cache/redis/linkedServers"
func (server RedisLinkedServer_SpecARM) GetType() string {
	return "Microsoft.Cache/redis/linkedServers"
}

type RedisLinkedServerPropertiesARM struct {
	LinkedRedisCacheId string `json:"linkedRedisCacheId"`

	//LinkedRedisCacheLocation: Location of the linked redis cache.
	LinkedRedisCacheLocation string `json:"linkedRedisCacheLocation"`

	//ServerRole: Role of the linked server.
	ServerRole RedisLinkedServerPropertiesServerRole `json:"serverRole"`
}