// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20211101

// Deprecated version of AuthorizationRule_STATUS. Use v1beta20211101.AuthorizationRule_STATUS instead
type AuthorizationRule_STATUS_ARM struct {
	Id         *string                                  `json:"id,omitempty"`
	Location   *string                                  `json:"location,omitempty"`
	Name       *string                                  `json:"name,omitempty"`
	Properties *AuthorizationRule_Properties_STATUS_ARM `json:"properties,omitempty"`
	SystemData *SystemData_STATUS_ARM                   `json:"systemData,omitempty"`
	Type       *string                                  `json:"type,omitempty"`
}

// Deprecated version of AuthorizationRule_Properties_STATUS. Use v1beta20211101.AuthorizationRule_Properties_STATUS instead
type AuthorizationRule_Properties_STATUS_ARM struct {
	Rights []AuthorizationRule_Properties_Rights_STATUS `json:"rights,omitempty"`
}

// Deprecated version of SystemData_STATUS. Use v1beta20211101.SystemData_STATUS instead
type SystemData_STATUS_ARM struct {
	CreatedAt          *string                               `json:"createdAt,omitempty"`
	CreatedBy          *string                               `json:"createdBy,omitempty"`
	CreatedByType      *SystemData_CreatedByType_STATUS      `json:"createdByType,omitempty"`
	LastModifiedAt     *string                               `json:"lastModifiedAt,omitempty"`
	LastModifiedBy     *string                               `json:"lastModifiedBy,omitempty"`
	LastModifiedByType *SystemData_LastModifiedByType_STATUS `json:"lastModifiedByType,omitempty"`
}

// Deprecated version of AuthorizationRule_Properties_Rights_STATUS. Use
// v1beta20211101.AuthorizationRule_Properties_Rights_STATUS instead
type AuthorizationRule_Properties_Rights_STATUS string

const (
	AuthorizationRule_Properties_Rights_STATUS_Listen = AuthorizationRule_Properties_Rights_STATUS("Listen")
	AuthorizationRule_Properties_Rights_STATUS_Manage = AuthorizationRule_Properties_Rights_STATUS("Manage")
	AuthorizationRule_Properties_Rights_STATUS_Send   = AuthorizationRule_Properties_Rights_STATUS("Send")
)

// Deprecated version of SystemData_CreatedByType_STATUS. Use v1beta20211101.SystemData_CreatedByType_STATUS instead
type SystemData_CreatedByType_STATUS string

const (
	SystemData_CreatedByType_STATUS_Application     = SystemData_CreatedByType_STATUS("Application")
	SystemData_CreatedByType_STATUS_Key             = SystemData_CreatedByType_STATUS("Key")
	SystemData_CreatedByType_STATUS_ManagedIdentity = SystemData_CreatedByType_STATUS("ManagedIdentity")
	SystemData_CreatedByType_STATUS_User            = SystemData_CreatedByType_STATUS("User")
)

// Deprecated version of SystemData_LastModifiedByType_STATUS. Use v1beta20211101.SystemData_LastModifiedByType_STATUS
// instead
type SystemData_LastModifiedByType_STATUS string

const (
	SystemData_LastModifiedByType_STATUS_Application     = SystemData_LastModifiedByType_STATUS("Application")
	SystemData_LastModifiedByType_STATUS_Key             = SystemData_LastModifiedByType_STATUS("Key")
	SystemData_LastModifiedByType_STATUS_ManagedIdentity = SystemData_LastModifiedByType_STATUS("ManagedIdentity")
	SystemData_LastModifiedByType_STATUS_User            = SystemData_LastModifiedByType_STATUS("User")
)