// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20200601

import "encoding/json"

// EventGrid Domain.
type Domain_STATUS_ARM struct {
	// Id: Fully qualified identifier of the resource.
	Id *string `json:"id,omitempty"`

	// Location: Location of the resource.
	Location *string `json:"location,omitempty"`

	// Name: Name of the resource.
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the domain.
	Properties *DomainProperties_STATUS_ARM `json:"properties,omitempty"`

	// SystemData: The system metadata relating to Domain resource.
	SystemData *SystemData_STATUS_ARM `json:"systemData,omitempty"`

	// Tags: Tags of the resource.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: Type of the resource.
	Type *string `json:"type,omitempty"`
}

// Properties of the Domain.
type DomainProperties_STATUS_ARM struct {
	// Endpoint: Endpoint for the domain.
	Endpoint *string `json:"endpoint,omitempty"`

	// InboundIpRules: This can be used to restrict traffic from specific IPs instead of all IPs. Note: These are considered
	// only if PublicNetworkAccess is enabled.
	InboundIpRules []InboundIpRule_STATUS_ARM `json:"inboundIpRules,omitempty"`

	// InputSchema: This determines the format that Event Grid should expect for incoming events published to the domain.
	InputSchema *DomainProperties_InputSchema_STATUS `json:"inputSchema,omitempty"`

	// InputSchemaMapping: Information about the InputSchemaMapping which specified the info about mapping event payload.
	InputSchemaMapping *InputSchemaMapping_STATUS_ARM `json:"inputSchemaMapping,omitempty"`

	// MetricResourceId: Metric resource id for the domain.
	MetricResourceId *string `json:"metricResourceId,omitempty"`

	// PrivateEndpointConnections: List of private endpoint connections.
	PrivateEndpointConnections []PrivateEndpointConnection_STATUS_Domain_SubResourceEmbedded_ARM `json:"privateEndpointConnections,omitempty"`

	// ProvisioningState: Provisioning state of the domain.
	ProvisioningState *DomainProperties_ProvisioningState_STATUS `json:"provisioningState,omitempty"`

	// PublicNetworkAccess: This determines if traffic is allowed over public network. By default it is enabled.
	// You can further restrict to specific IPs by configuring <seealso
	// cref="P:Microsoft.Azure.Events.ResourceProvider.Common.Contracts.DomainProperties.InboundIpRules" />
	PublicNetworkAccess *DomainProperties_PublicNetworkAccess_STATUS `json:"publicNetworkAccess,omitempty"`
}

// Metadata pertaining to creation and last modification of the resource.
type SystemData_STATUS_ARM struct {
	// CreatedAt: The timestamp of resource creation (UTC).
	CreatedAt *string `json:"createdAt,omitempty"`

	// CreatedBy: The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// CreatedByType: The type of identity that created the resource.
	CreatedByType *SystemData_CreatedByType_STATUS `json:"createdByType,omitempty"`

	// LastModifiedAt: The timestamp of resource last modification (UTC)
	LastModifiedAt *string `json:"lastModifiedAt,omitempty"`

	// LastModifiedBy: The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// LastModifiedByType: The type of identity that last modified the resource.
	LastModifiedByType *SystemData_LastModifiedByType_STATUS `json:"lastModifiedByType,omitempty"`
}

type InboundIpRule_STATUS_ARM struct {
	// Action: Action to perform based on the match or no match of the IpMask.
	Action *InboundIpRule_Action_STATUS `json:"action,omitempty"`

	// IpMask: IP Address in CIDR notation e.g., 10.0.0.0/8.
	IpMask *string `json:"ipMask,omitempty"`
}

type InputSchemaMapping_STATUS_ARM struct {
	// Json: Mutually exclusive with all other properties
	Json *JsonInputSchemaMapping_STATUS_ARM `json:"json,omitempty"`
}

// MarshalJSON defers JSON marshaling to the first non-nil property, because InputSchemaMapping_STATUS_ARM represents a discriminated union (JSON OneOf)
func (mapping InputSchemaMapping_STATUS_ARM) MarshalJSON() ([]byte, error) {
	if mapping.Json != nil {
		return json.Marshal(mapping.Json)
	}
	return nil, nil
}

// UnmarshalJSON unmarshals the InputSchemaMapping_STATUS_ARM
func (mapping *InputSchemaMapping_STATUS_ARM) UnmarshalJSON(data []byte) error {
	var rawJson map[string]interface{}
	err := json.Unmarshal(data, &rawJson)
	if err != nil {
		return err
	}
	discriminator := rawJson["inputSchemaMappingType"]
	if discriminator == "Json" {
		mapping.Json = &JsonInputSchemaMapping_STATUS_ARM{}
		return json.Unmarshal(data, mapping.Json)
	}

	// No error
	return nil
}

type PrivateEndpointConnection_STATUS_Domain_SubResourceEmbedded_ARM struct {
	// Id: Fully qualified identifier of the resource.
	Id *string `json:"id,omitempty"`
}

type SystemData_CreatedByType_STATUS string

const (
	SystemData_CreatedByType_STATUS_Application     = SystemData_CreatedByType_STATUS("Application")
	SystemData_CreatedByType_STATUS_Key             = SystemData_CreatedByType_STATUS("Key")
	SystemData_CreatedByType_STATUS_ManagedIdentity = SystemData_CreatedByType_STATUS("ManagedIdentity")
	SystemData_CreatedByType_STATUS_User            = SystemData_CreatedByType_STATUS("User")
)

type SystemData_LastModifiedByType_STATUS string

const (
	SystemData_LastModifiedByType_STATUS_Application     = SystemData_LastModifiedByType_STATUS("Application")
	SystemData_LastModifiedByType_STATUS_Key             = SystemData_LastModifiedByType_STATUS("Key")
	SystemData_LastModifiedByType_STATUS_ManagedIdentity = SystemData_LastModifiedByType_STATUS("ManagedIdentity")
	SystemData_LastModifiedByType_STATUS_User            = SystemData_LastModifiedByType_STATUS("User")
)

type JsonInputSchemaMapping_STATUS_ARM struct {
	// InputSchemaMappingType: Type of the custom mapping
	InputSchemaMappingType JsonInputSchemaMapping_InputSchemaMappingType_STATUS `json:"inputSchemaMappingType,omitempty"`

	// Properties: JSON Properties of the input schema mapping
	Properties *JsonInputSchemaMappingProperties_STATUS_ARM `json:"properties,omitempty"`
}

// This can be used to map properties of a source schema (or default values, for certain supported properties) to
// properties of the EventGridEvent schema.
type JsonInputSchemaMappingProperties_STATUS_ARM struct {
	// DataVersion: The mapping information for the DataVersion property of the Event Grid Event.
	DataVersion *JsonFieldWithDefault_STATUS_ARM `json:"dataVersion,omitempty"`

	// EventTime: The mapping information for the EventTime property of the Event Grid Event.
	EventTime *JsonField_STATUS_ARM `json:"eventTime,omitempty"`

	// EventType: The mapping information for the EventType property of the Event Grid Event.
	EventType *JsonFieldWithDefault_STATUS_ARM `json:"eventType,omitempty"`

	// Id: The mapping information for the Id property of the Event Grid Event.
	Id *JsonField_STATUS_ARM `json:"id,omitempty"`

	// Subject: The mapping information for the Subject property of the Event Grid Event.
	Subject *JsonFieldWithDefault_STATUS_ARM `json:"subject,omitempty"`

	// Topic: The mapping information for the Topic property of the Event Grid Event.
	Topic *JsonField_STATUS_ARM `json:"topic,omitempty"`
}

// This is used to express the source of an input schema mapping for a single target field in the Event Grid Event schema.
// This is currently used in the mappings for the 'id', 'topic' and 'eventtime' properties. This represents a field in the
// input event schema.
type JsonField_STATUS_ARM struct {
	// SourceField: Name of a field in the input event schema that's to be used as the source of a mapping.
	SourceField *string `json:"sourceField,omitempty"`
}

// This is used to express the source of an input schema mapping for a single target field
// in the Event Grid Event schema.
// This is currently used in the mappings for the 'subject',
// 'eventtype' and 'dataversion' properties. This represents a
// field in the input event schema
// along with a default value to be used, and at least one of these two properties should
// be provided.
type JsonFieldWithDefault_STATUS_ARM struct {
	// DefaultValue: The default value to be used for mapping when a SourceField is not provided or if there's no property with
	// the specified name in the published JSON event payload.
	DefaultValue *string `json:"defaultValue,omitempty"`

	// SourceField: Name of a field in the input event schema that's to be used as the source of a mapping.
	SourceField *string `json:"sourceField,omitempty"`
}
