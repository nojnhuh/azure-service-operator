// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20211101storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=eventhub.azure.com,resources=namespaceseventhubsauthorizationrules,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=eventhub.azure.com,resources={namespaceseventhubsauthorizationrules/status,namespaceseventhubsauthorizationrules/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Storage version of v1alpha1api20211101.NamespacesEventhubsAuthorizationRule
type NamespacesEventhubsAuthorizationRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NamespacesEventhubsAuthorizationRules_SPEC `json:"spec,omitempty"`
	Status            AuthorizationRule_Status                   `json:"status,omitempty"`
}

var _ conditions.Conditioner = &NamespacesEventhubsAuthorizationRule{}

// GetConditions returns the conditions of the resource
func (rule *NamespacesEventhubsAuthorizationRule) GetConditions() conditions.Conditions {
	return rule.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (rule *NamespacesEventhubsAuthorizationRule) SetConditions(conditions conditions.Conditions) {
	rule.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &NamespacesEventhubsAuthorizationRule{}

// AzureName returns the Azure name of the resource
func (rule *NamespacesEventhubsAuthorizationRule) AzureName() string {
	return rule.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (rule NamespacesEventhubsAuthorizationRule) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetResourceKind returns the kind of the resource
func (rule *NamespacesEventhubsAuthorizationRule) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (rule *NamespacesEventhubsAuthorizationRule) GetSpec() genruntime.ConvertibleSpec {
	return &rule.Spec
}

// GetStatus returns the status of this resource
func (rule *NamespacesEventhubsAuthorizationRule) GetStatus() genruntime.ConvertibleStatus {
	return &rule.Status
}

// GetType returns the ARM Type of the resource. This is always ""
func (rule *NamespacesEventhubsAuthorizationRule) GetType() string {
	return ""
}

// NewEmptyStatus returns a new empty (blank) status
func (rule *NamespacesEventhubsAuthorizationRule) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &AuthorizationRule_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (rule *NamespacesEventhubsAuthorizationRule) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(rule.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  rule.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (rule *NamespacesEventhubsAuthorizationRule) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*AuthorizationRule_Status); ok {
		rule.Status = *st
		return nil
	}

	// Convert status to required version
	var st AuthorizationRule_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	rule.Status = st
	return nil
}

// Hub marks that this NamespacesEventhubsAuthorizationRule is the hub type for conversion
func (rule *NamespacesEventhubsAuthorizationRule) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (rule *NamespacesEventhubsAuthorizationRule) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: rule.Spec.OriginalVersion,
		Kind:    "NamespacesEventhubsAuthorizationRule",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1alpha1api20211101.NamespacesEventhubsAuthorizationRule
type NamespacesEventhubsAuthorizationRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NamespacesEventhubsAuthorizationRule `json:"items"`
}

//Storage version of v1alpha1api20211101.NamespacesEventhubsAuthorizationRules_SPEC
type NamespacesEventhubsAuthorizationRules_SPEC struct {
	//AzureName: The name of the resource in Azure. This is often the same as the name
	//of the resource in Kubernetes but it doesn't have to be.
	AzureName       string `json:"azureName"`
	OriginalVersion string `json:"originalVersion"`

	// +kubebuilder:validation:Required
	Owner       genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner" kind:"ResourceGroup"`
	PropertyBag genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	Rights      []string                          `json:"rights,omitempty"`
}

var _ genruntime.ConvertibleSpec = &NamespacesEventhubsAuthorizationRules_SPEC{}

// ConvertSpecFrom populates our NamespacesEventhubsAuthorizationRules_SPEC from the provided source
func (spec *NamespacesEventhubsAuthorizationRules_SPEC) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == spec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(spec)
}

// ConvertSpecTo populates the provided destination from our NamespacesEventhubsAuthorizationRules_SPEC
func (spec *NamespacesEventhubsAuthorizationRules_SPEC) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == spec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(spec)
}

func init() {
	SchemeBuilder.Register(&NamespacesEventhubsAuthorizationRule{}, &NamespacesEventhubsAuthorizationRuleList{})
}
