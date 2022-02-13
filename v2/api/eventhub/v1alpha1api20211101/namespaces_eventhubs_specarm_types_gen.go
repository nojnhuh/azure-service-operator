// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20211101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type NamespacesEventhubs_SPECARM struct {
	AzureName string `json:"azureName"`
	Name      string `json:"name"`

	//Properties: Properties supplied to the Create Or Update Event Hub operation.
	Properties *NamespacesEventhubs_Properties_SPECARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &NamespacesEventhubs_SPECARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (specarm NamespacesEventhubs_SPECARM) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetName returns the Name of the resource
func (specarm NamespacesEventhubs_SPECARM) GetName() string {
	return specarm.Name
}

// GetType returns the ARM Type of the resource. This is always ""
func (specarm NamespacesEventhubs_SPECARM) GetType() string {
	return ""
}

type NamespacesEventhubs_Properties_SPECARM struct {
	//CaptureDescription: Properties of capture description
	CaptureDescription *CaptureDescription_SpecARM `json:"captureDescription,omitempty"`

	//MessageRetentionInDays: Number of days to retain the events for this Event Hub,
	//value should be 1 to 7 days
	MessageRetentionInDays *int `json:"messageRetentionInDays,omitempty"`

	//PartitionCount: Number of partitions created for the Event Hub, allowed values
	//are from 1 to 32 partitions.
	PartitionCount *int `json:"partitionCount,omitempty"`

	//Status: Enumerates the possible values for the status of the Event Hub.
	Status *NamespacesEventhubs_Properties_Status_SPEC `json:"status,omitempty"`
}

type CaptureDescription_SpecARM struct {
	//Destination: Properties of Destination where capture will be stored. (Storage
	//Account, Blob Names)
	Destination *Destination_SpecARM `json:"destination,omitempty"`

	//Enabled: A value that indicates whether capture description is enabled.
	Enabled *bool `json:"enabled,omitempty"`

	//Encoding: Enumerates the possible values for the encoding format of capture
	//description. Note: 'AvroDeflate' will be deprecated in New API Version
	Encoding *CaptureDescription_Encoding_Spec `json:"encoding,omitempty"`

	//IntervalInSeconds: The time window allows you to set the frequency with which
	//the capture to Azure Blobs will happen, value should between 60 to 900 seconds
	IntervalInSeconds *int `json:"intervalInSeconds,omitempty"`

	//SizeLimitInBytes: The size window defines the amount of data built up in your
	//Event Hub before an capture operation, value should be between 10485760 to
	//524288000 bytes
	SizeLimitInBytes *int `json:"sizeLimitInBytes,omitempty"`

	//SkipEmptyArchives: A value that indicates whether to Skip Empty Archives
	SkipEmptyArchives *bool `json:"skipEmptyArchives,omitempty"`
}

type Destination_SpecARM struct {
	//Name: Name for capture destination
	Name *string `json:"name,omitempty"`

	//Properties: Properties describing the storage account, blob container and
	//archive name format for capture destination
	Properties *Destination_Properties_SpecARM `json:"properties,omitempty"`
}

type Destination_Properties_SpecARM struct {
	//ArchiveNameFormat: Blob naming convention for archive, e.g.
	//{Namespace}/{EventHub}/{PartitionId}/{Year}/{Month}/{Day}/{Hour}/{Minute}/{Second}.
	//Here all the parameters (Namespace,EventHub .. etc) are mandatory irrespective
	//of order
	ArchiveNameFormat *string `json:"archiveNameFormat,omitempty"`

	//BlobContainer: Blob container Name
	BlobContainer *string `json:"blobContainer,omitempty"`

	//DataLakeAccountName: The Azure Data Lake Store name for the captured events
	DataLakeAccountName *string `json:"dataLakeAccountName,omitempty"`

	//DataLakeFolderPath: The destination folder path for the captured events
	DataLakeFolderPath *string `json:"dataLakeFolderPath,omitempty"`

	//DataLakeSubscriptionId: Subscription Id of Azure Data Lake Store
	DataLakeSubscriptionId   *string `json:"dataLakeSubscriptionId,omitempty"`
	StorageAccountResourceId *string `json:"storageAccountResourceId,omitempty"`
}