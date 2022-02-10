// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210301

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type RedisEnterpriseDatabases_SPECARM struct {
	AzureName string `json:"azureName"`
	Name      string `json:"name"`

	//Properties: Other properties of the database.
	Properties *DatabaseProperties_SpecARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &RedisEnterpriseDatabases_SPECARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-03-01"
func (specarm RedisEnterpriseDatabases_SPECARM) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetName returns the Name of the resource
func (specarm RedisEnterpriseDatabases_SPECARM) GetName() string {
	return specarm.Name
}

// GetType returns the ARM Type of the resource. This is always ""
func (specarm RedisEnterpriseDatabases_SPECARM) GetType() string {
	return ""
}

type DatabaseProperties_SpecARM struct {
	//ClientProtocol: Specifies whether redis clients can connect using TLS-encrypted
	//or plaintext redis protocols. Default is TLS-encrypted.
	ClientProtocol *DatabasePropertiesSpecClientProtocol `json:"clientProtocol,omitempty"`

	//ClusteringPolicy: Clustering policy - default is OSSCluster. Specified at create
	//time.
	ClusteringPolicy *DatabasePropertiesSpecClusteringPolicy `json:"clusteringPolicy,omitempty"`

	//EvictionPolicy: Redis eviction policy - default is VolatileLRU
	EvictionPolicy *DatabasePropertiesSpecEvictionPolicy `json:"evictionPolicy,omitempty"`

	//Modules: Optional set of redis modules to enable in this database - modules can
	//only be added at creation time.
	Modules []Module_SpecARM `json:"modules,omitempty"`

	//Persistence: Persistence settings
	Persistence *Persistence_SpecARM `json:"persistence,omitempty"`

	//Port: TCP port of the database endpoint. Specified at create time. Defaults to
	//an available port.
	Port *int `json:"port,omitempty"`
}

type Module_SpecARM struct {
	//Args: Configuration options for the module, e.g. 'ERROR_RATE 0.00 INITIAL_SIZE
	//400'.
	Args *string `json:"args,omitempty"`

	//Name: The name of the module, e.g. 'RedisBloom', 'RediSearch', 'RedisTimeSeries'
	Name string `json:"name"`
}

type Persistence_SpecARM struct {
	//AofEnabled: Sets whether AOF is enabled.
	AofEnabled *bool `json:"aofEnabled,omitempty"`

	//AofFrequency: Sets the frequency at which data is written to disk.
	AofFrequency *PersistenceSpecAofFrequency `json:"aofFrequency,omitempty"`

	//RdbEnabled: Sets whether RDB is enabled.
	RdbEnabled *bool `json:"rdbEnabled,omitempty"`

	//RdbFrequency: Sets the frequency at which a snapshot of the database is created.
	RdbFrequency *PersistenceSpecRdbFrequency `json:"rdbFrequency,omitempty"`
}
