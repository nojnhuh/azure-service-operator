// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210401storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=storage.azure.com,resources=storageaccountsqueueservices,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=storage.azure.com,resources={storageaccountsqueueservices/status,storageaccountsqueueservices/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Storage version of v1alpha1api20210401.StorageAccountsQueueService
type StorageAccountsQueueService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageAccountsQueueServices_SPEC `json:"spec,omitempty"`
	Status            QueueServiceProperties_Status     `json:"status,omitempty"`
}

var _ conditions.Conditioner = &StorageAccountsQueueService{}

// GetConditions returns the conditions of the resource
func (service *StorageAccountsQueueService) GetConditions() conditions.Conditions {
	return service.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (service *StorageAccountsQueueService) SetConditions(conditions conditions.Conditions) {
	service.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &StorageAccountsQueueService{}

// AzureName returns the Azure name of the resource
func (service *StorageAccountsQueueService) AzureName() string {
	return service.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-04-01"
func (service StorageAccountsQueueService) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetResourceKind returns the kind of the resource
func (service *StorageAccountsQueueService) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (service *StorageAccountsQueueService) GetSpec() genruntime.ConvertibleSpec {
	return &service.Spec
}

// GetStatus returns the status of this resource
func (service *StorageAccountsQueueService) GetStatus() genruntime.ConvertibleStatus {
	return &service.Status
}

// GetType returns the ARM Type of the resource. This is always ""
func (service *StorageAccountsQueueService) GetType() string {
	return ""
}

// NewEmptyStatus returns a new empty (blank) status
func (service *StorageAccountsQueueService) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &QueueServiceProperties_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (service *StorageAccountsQueueService) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(service.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  service.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (service *StorageAccountsQueueService) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*QueueServiceProperties_Status); ok {
		service.Status = *st
		return nil
	}

	// Convert status to required version
	var st QueueServiceProperties_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	service.Status = st
	return nil
}

// Hub marks that this StorageAccountsQueueService is the hub type for conversion
func (service *StorageAccountsQueueService) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (service *StorageAccountsQueueService) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: service.Spec.OriginalVersion,
		Kind:    "StorageAccountsQueueService",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1alpha1api20210401.StorageAccountsQueueService
type StorageAccountsQueueServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageAccountsQueueService `json:"items"`
}

//Storage version of v1alpha1api20210401.QueueServiceProperties_Status
type QueueServiceProperties_Status struct {
	Conditions  []conditions.Condition `json:"conditions,omitempty"`
	Cors        *CorsRules_Status      `json:"cors,omitempty"`
	Id          *string                `json:"id,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &QueueServiceProperties_Status{}

// ConvertStatusFrom populates our QueueServiceProperties_Status from the provided source
func (properties *QueueServiceProperties_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == properties {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(properties)
}

// ConvertStatusTo populates the provided destination from our QueueServiceProperties_Status
func (properties *QueueServiceProperties_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == properties {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(properties)
}

//Storage version of v1alpha1api20210401.StorageAccountsQueueServices_SPEC
type StorageAccountsQueueServices_SPEC struct {
	//AzureName: The name of the resource in Azure. This is often the same as the name
	//of the resource in Kubernetes but it doesn't have to be.
	AzureName       string          `json:"azureName"`
	Cors            *CorsRules_Spec `json:"cors,omitempty"`
	OriginalVersion string          `json:"originalVersion"`

	// +kubebuilder:validation:Required
	Owner       genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner" kind:"ResourceGroup"`
	PropertyBag genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
}

var _ genruntime.ConvertibleSpec = &StorageAccountsQueueServices_SPEC{}

// ConvertSpecFrom populates our StorageAccountsQueueServices_SPEC from the provided source
func (spec *StorageAccountsQueueServices_SPEC) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == spec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(spec)
}

// ConvertSpecTo populates the provided destination from our StorageAccountsQueueServices_SPEC
func (spec *StorageAccountsQueueServices_SPEC) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == spec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(spec)
}

func init() {
	SchemeBuilder.Register(&StorageAccountsQueueService{}, &StorageAccountsQueueServiceList{})
}
