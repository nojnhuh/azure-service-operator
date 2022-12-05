// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210515

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type DatabaseAccounts_MongodbDatabases_Collection_Spec_ARM struct {
	// Location: The location of the resource group to which the resource belongs.
	Location *string `json:"location,omitempty"`

	// Name: Cosmos DB collection name.
	Name string `json:"name,omitempty"`

	// Properties: Properties to create and update Azure Cosmos DB MongoDB collection.
	Properties *MongoDBCollectionCreateUpdateProperties_ARM `json:"properties,omitempty"`

	// Tags: Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this
	// resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no
	// greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template
	// type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph",
	// "DocumentDB", and "MongoDB".
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &DatabaseAccounts_MongodbDatabases_Collection_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (collection DatabaseAccounts_MongodbDatabases_Collection_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (collection *DatabaseAccounts_MongodbDatabases_Collection_Spec_ARM) GetName() string {
	return collection.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/mongodbDatabases/collections"
func (collection *DatabaseAccounts_MongodbDatabases_Collection_Spec_ARM) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/mongodbDatabases/collections"
}

// Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/MongoDBCollectionCreateUpdateProperties
type MongoDBCollectionCreateUpdateProperties_ARM struct {
	// Options: CreateUpdateOptions are a list of key-value pairs that describe the resource. Supported keys are "If-Match",
	// "If-None-Match", "Session-Token" and "Throughput"
	Options *CreateUpdateOptions_ARM `json:"options,omitempty"`

	// Resource: Cosmos DB MongoDB collection resource object
	Resource *MongoDBCollectionResource_ARM `json:"resource,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/MongoDBCollectionResource
type MongoDBCollectionResource_ARM struct {
	// AnalyticalStorageTtl: Analytical TTL.
	AnalyticalStorageTtl *int `json:"analyticalStorageTtl,omitempty"`

	// Id: Name of the Cosmos DB MongoDB collection
	Id *string `json:"id,omitempty"`

	// Indexes: List of index keys
	Indexes []MongoIndex_ARM `json:"indexes,omitempty"`

	// ShardKey: The shard key and partition kind pair, only support "Hash" partition kind
	ShardKey map[string]string `json:"shardKey,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/MongoIndex
type MongoIndex_ARM struct {
	// Key: Cosmos DB MongoDB collection resource object
	Key *MongoIndexKeys_ARM `json:"key,omitempty"`

	// Options: Cosmos DB MongoDB collection index options
	Options *MongoIndexOptions_ARM `json:"options,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/MongoIndexKeys
type MongoIndexKeys_ARM struct {
	// Keys: List of keys for each MongoDB collection in the Azure Cosmos DB service
	Keys []string `json:"keys,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/MongoIndexOptions
type MongoIndexOptions_ARM struct {
	// ExpireAfterSeconds: Expire after seconds
	ExpireAfterSeconds *int `json:"expireAfterSeconds,omitempty"`

	// Unique: Is unique or not
	Unique *bool `json:"unique,omitempty"`
}