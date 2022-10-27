// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211101

import (
	"fmt"
	v20211101s "github.com/Azure/azure-service-operator/v2/api/sql/v1beta20211101storage"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Generator information:
// - Generated from: /sql/resource-manager/Microsoft.Sql/stable/2021-11-01/VirtualNetworkRules.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/virtualNetworkRules/{virtualNetworkRuleName}
type ServersVirtualNetworkRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Servers_VirtualNetworkRule_Spec   `json:"spec,omitempty"`
	Status            Servers_VirtualNetworkRule_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &ServersVirtualNetworkRule{}

// GetConditions returns the conditions of the resource
func (rule *ServersVirtualNetworkRule) GetConditions() conditions.Conditions {
	return rule.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (rule *ServersVirtualNetworkRule) SetConditions(conditions conditions.Conditions) {
	rule.Status.Conditions = conditions
}

var _ conversion.Convertible = &ServersVirtualNetworkRule{}

// ConvertFrom populates our ServersVirtualNetworkRule from the provided hub ServersVirtualNetworkRule
func (rule *ServersVirtualNetworkRule) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v20211101s.ServersVirtualNetworkRule)
	if !ok {
		return fmt.Errorf("expected sql/v1beta20211101storage/ServersVirtualNetworkRule but received %T instead", hub)
	}

	return rule.AssignProperties_From_ServersVirtualNetworkRule(source)
}

// ConvertTo populates the provided hub ServersVirtualNetworkRule from our ServersVirtualNetworkRule
func (rule *ServersVirtualNetworkRule) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20211101s.ServersVirtualNetworkRule)
	if !ok {
		return fmt.Errorf("expected sql/v1beta20211101storage/ServersVirtualNetworkRule but received %T instead", hub)
	}

	return rule.AssignProperties_To_ServersVirtualNetworkRule(destination)
}

// +kubebuilder:webhook:path=/mutate-sql-azure-com-v1beta20211101-serversvirtualnetworkrule,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=sql.azure.com,resources=serversvirtualnetworkrules,verbs=create;update,versions=v1beta20211101,name=default.v1beta20211101.serversvirtualnetworkrules.sql.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &ServersVirtualNetworkRule{}

// Default applies defaults to the ServersVirtualNetworkRule resource
func (rule *ServersVirtualNetworkRule) Default() {
	rule.defaultImpl()
	var temp any = rule
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (rule *ServersVirtualNetworkRule) defaultAzureName() {
	if rule.Spec.AzureName == "" {
		rule.Spec.AzureName = rule.Name
	}
}

// defaultImpl applies the code generated defaults to the ServersVirtualNetworkRule resource
func (rule *ServersVirtualNetworkRule) defaultImpl() { rule.defaultAzureName() }

var _ genruntime.KubernetesResource = &ServersVirtualNetworkRule{}

// AzureName returns the Azure name of the resource
func (rule *ServersVirtualNetworkRule) AzureName() string {
	return rule.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (rule ServersVirtualNetworkRule) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (rule *ServersVirtualNetworkRule) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (rule *ServersVirtualNetworkRule) GetSpec() genruntime.ConvertibleSpec {
	return &rule.Spec
}

// GetStatus returns the status of this resource
func (rule *ServersVirtualNetworkRule) GetStatus() genruntime.ConvertibleStatus {
	return &rule.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Sql/servers/virtualNetworkRules"
func (rule *ServersVirtualNetworkRule) GetType() string {
	return "Microsoft.Sql/servers/virtualNetworkRules"
}

// NewEmptyStatus returns a new empty (blank) status
func (rule *ServersVirtualNetworkRule) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Servers_VirtualNetworkRule_STATUS{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (rule *ServersVirtualNetworkRule) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(rule.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  rule.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (rule *ServersVirtualNetworkRule) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Servers_VirtualNetworkRule_STATUS); ok {
		rule.Status = *st
		return nil
	}

	// Convert status to required version
	var st Servers_VirtualNetworkRule_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	rule.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-sql-azure-com-v1beta20211101-serversvirtualnetworkrule,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=sql.azure.com,resources=serversvirtualnetworkrules,verbs=create;update,versions=v1beta20211101,name=validate.v1beta20211101.serversvirtualnetworkrules.sql.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &ServersVirtualNetworkRule{}

// ValidateCreate validates the creation of the resource
func (rule *ServersVirtualNetworkRule) ValidateCreate() error {
	validations := rule.createValidations()
	var temp any = rule
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// ValidateDelete validates the deletion of the resource
func (rule *ServersVirtualNetworkRule) ValidateDelete() error {
	validations := rule.deleteValidations()
	var temp any = rule
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// ValidateUpdate validates an update of the resource
func (rule *ServersVirtualNetworkRule) ValidateUpdate(old runtime.Object) error {
	validations := rule.updateValidations()
	var temp any = rule
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation(old)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// createValidations validates the creation of the resource
func (rule *ServersVirtualNetworkRule) createValidations() []func() error {
	return []func() error{rule.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (rule *ServersVirtualNetworkRule) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (rule *ServersVirtualNetworkRule) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return rule.validateResourceReferences()
		},
		rule.validateWriteOnceProperties}
}

// validateResourceReferences validates all resource references
func (rule *ServersVirtualNetworkRule) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&rule.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (rule *ServersVirtualNetworkRule) validateWriteOnceProperties(old runtime.Object) error {
	oldObj, ok := old.(*ServersVirtualNetworkRule)
	if !ok {
		return nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, rule)
}

// AssignProperties_From_ServersVirtualNetworkRule populates our ServersVirtualNetworkRule from the provided source ServersVirtualNetworkRule
func (rule *ServersVirtualNetworkRule) AssignProperties_From_ServersVirtualNetworkRule(source *v20211101s.ServersVirtualNetworkRule) error {

	// ObjectMeta
	rule.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec Servers_VirtualNetworkRule_Spec
	err := spec.AssignProperties_From_Servers_VirtualNetworkRule_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Servers_VirtualNetworkRule_Spec() to populate field Spec")
	}
	rule.Spec = spec

	// Status
	var status Servers_VirtualNetworkRule_STATUS
	err = status.AssignProperties_From_Servers_VirtualNetworkRule_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Servers_VirtualNetworkRule_STATUS() to populate field Status")
	}
	rule.Status = status

	// No error
	return nil
}

// AssignProperties_To_ServersVirtualNetworkRule populates the provided destination ServersVirtualNetworkRule from our ServersVirtualNetworkRule
func (rule *ServersVirtualNetworkRule) AssignProperties_To_ServersVirtualNetworkRule(destination *v20211101s.ServersVirtualNetworkRule) error {

	// ObjectMeta
	destination.ObjectMeta = *rule.ObjectMeta.DeepCopy()

	// Spec
	var spec v20211101s.Servers_VirtualNetworkRule_Spec
	err := rule.Spec.AssignProperties_To_Servers_VirtualNetworkRule_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Servers_VirtualNetworkRule_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20211101s.Servers_VirtualNetworkRule_STATUS
	err = rule.Status.AssignProperties_To_Servers_VirtualNetworkRule_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Servers_VirtualNetworkRule_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (rule *ServersVirtualNetworkRule) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: rule.Spec.OriginalVersion(),
		Kind:    "ServersVirtualNetworkRule",
	}
}

// +kubebuilder:object:root=true
// Generator information:
// - Generated from: /sql/resource-manager/Microsoft.Sql/stable/2021-11-01/VirtualNetworkRules.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/virtualNetworkRules/{virtualNetworkRuleName}
type ServersVirtualNetworkRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServersVirtualNetworkRule `json:"items"`
}

type Servers_VirtualNetworkRule_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string `json:"azureName,omitempty"`

	// IgnoreMissingVnetServiceEndpoint: Create firewall rule before the virtual network has vnet service endpoint enabled.
	IgnoreMissingVnetServiceEndpoint *bool `json:"ignoreMissingVnetServiceEndpoint,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a sql.azure.com/Server resource
	Owner *genruntime.KnownResourceReference `group:"sql.azure.com" json:"owner,omitempty" kind:"Server"`

	// +kubebuilder:validation:Required
	// VirtualNetworkSubnetReference: The ARM resource id of the virtual network subnet.
	VirtualNetworkSubnetReference *genruntime.ResourceReference `armReference:"VirtualNetworkSubnetId" json:"virtualNetworkSubnetReference,omitempty"`
}

var _ genruntime.ARMTransformer = &Servers_VirtualNetworkRule_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (rule *Servers_VirtualNetworkRule_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if rule == nil {
		return nil, nil
	}
	result := &Servers_VirtualNetworkRule_Spec_ARM{}

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	if rule.IgnoreMissingVnetServiceEndpoint != nil || rule.VirtualNetworkSubnetReference != nil {
		result.Properties = &VirtualNetworkRuleProperties_ARM{}
	}
	if rule.IgnoreMissingVnetServiceEndpoint != nil {
		ignoreMissingVnetServiceEndpoint := *rule.IgnoreMissingVnetServiceEndpoint
		result.Properties.IgnoreMissingVnetServiceEndpoint = &ignoreMissingVnetServiceEndpoint
	}
	if rule.VirtualNetworkSubnetReference != nil {
		virtualNetworkSubnetIdARMID, err := resolved.ResolvedReferences.Lookup(*rule.VirtualNetworkSubnetReference)
		if err != nil {
			return nil, err
		}
		virtualNetworkSubnetId := virtualNetworkSubnetIdARMID
		result.Properties.VirtualNetworkSubnetId = &virtualNetworkSubnetId
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (rule *Servers_VirtualNetworkRule_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Servers_VirtualNetworkRule_Spec_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (rule *Servers_VirtualNetworkRule_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Servers_VirtualNetworkRule_Spec_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Servers_VirtualNetworkRule_Spec_ARM, got %T", armInput)
	}

	// Set property ‘AzureName’:
	rule.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property ‘IgnoreMissingVnetServiceEndpoint’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.IgnoreMissingVnetServiceEndpoint != nil {
			ignoreMissingVnetServiceEndpoint := *typedInput.Properties.IgnoreMissingVnetServiceEndpoint
			rule.IgnoreMissingVnetServiceEndpoint = &ignoreMissingVnetServiceEndpoint
		}
	}

	// Set property ‘Owner’:
	rule.Owner = &genruntime.KnownResourceReference{Name: owner.Name}

	// no assignment for property ‘VirtualNetworkSubnetReference’

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &Servers_VirtualNetworkRule_Spec{}

// ConvertSpecFrom populates our Servers_VirtualNetworkRule_Spec from the provided source
func (rule *Servers_VirtualNetworkRule_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20211101s.Servers_VirtualNetworkRule_Spec)
	if ok {
		// Populate our instance from source
		return rule.AssignProperties_From_Servers_VirtualNetworkRule_Spec(src)
	}

	// Convert to an intermediate form
	src = &v20211101s.Servers_VirtualNetworkRule_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = rule.AssignProperties_From_Servers_VirtualNetworkRule_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our Servers_VirtualNetworkRule_Spec
func (rule *Servers_VirtualNetworkRule_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20211101s.Servers_VirtualNetworkRule_Spec)
	if ok {
		// Populate destination from our instance
		return rule.AssignProperties_To_Servers_VirtualNetworkRule_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v20211101s.Servers_VirtualNetworkRule_Spec{}
	err := rule.AssignProperties_To_Servers_VirtualNetworkRule_Spec(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecTo()")
	}

	// Update dst from our instance
	err = dst.ConvertSpecTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecTo()")
	}

	return nil
}

// AssignProperties_From_Servers_VirtualNetworkRule_Spec populates our Servers_VirtualNetworkRule_Spec from the provided source Servers_VirtualNetworkRule_Spec
func (rule *Servers_VirtualNetworkRule_Spec) AssignProperties_From_Servers_VirtualNetworkRule_Spec(source *v20211101s.Servers_VirtualNetworkRule_Spec) error {

	// AzureName
	rule.AzureName = source.AzureName

	// IgnoreMissingVnetServiceEndpoint
	if source.IgnoreMissingVnetServiceEndpoint != nil {
		ignoreMissingVnetServiceEndpoint := *source.IgnoreMissingVnetServiceEndpoint
		rule.IgnoreMissingVnetServiceEndpoint = &ignoreMissingVnetServiceEndpoint
	} else {
		rule.IgnoreMissingVnetServiceEndpoint = nil
	}

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		rule.Owner = &owner
	} else {
		rule.Owner = nil
	}

	// VirtualNetworkSubnetReference
	if source.VirtualNetworkSubnetReference != nil {
		virtualNetworkSubnetReference := source.VirtualNetworkSubnetReference.Copy()
		rule.VirtualNetworkSubnetReference = &virtualNetworkSubnetReference
	} else {
		rule.VirtualNetworkSubnetReference = nil
	}

	// No error
	return nil
}

// AssignProperties_To_Servers_VirtualNetworkRule_Spec populates the provided destination Servers_VirtualNetworkRule_Spec from our Servers_VirtualNetworkRule_Spec
func (rule *Servers_VirtualNetworkRule_Spec) AssignProperties_To_Servers_VirtualNetworkRule_Spec(destination *v20211101s.Servers_VirtualNetworkRule_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = rule.AzureName

	// IgnoreMissingVnetServiceEndpoint
	if rule.IgnoreMissingVnetServiceEndpoint != nil {
		ignoreMissingVnetServiceEndpoint := *rule.IgnoreMissingVnetServiceEndpoint
		destination.IgnoreMissingVnetServiceEndpoint = &ignoreMissingVnetServiceEndpoint
	} else {
		destination.IgnoreMissingVnetServiceEndpoint = nil
	}

	// OriginalVersion
	destination.OriginalVersion = rule.OriginalVersion()

	// Owner
	if rule.Owner != nil {
		owner := rule.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// VirtualNetworkSubnetReference
	if rule.VirtualNetworkSubnetReference != nil {
		virtualNetworkSubnetReference := rule.VirtualNetworkSubnetReference.Copy()
		destination.VirtualNetworkSubnetReference = &virtualNetworkSubnetReference
	} else {
		destination.VirtualNetworkSubnetReference = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// OriginalVersion returns the original API version used to create the resource.
func (rule *Servers_VirtualNetworkRule_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (rule *Servers_VirtualNetworkRule_Spec) SetAzureName(azureName string) {
	rule.AzureName = azureName
}

type Servers_VirtualNetworkRule_STATUS struct {
	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// IgnoreMissingVnetServiceEndpoint: Create firewall rule before the virtual network has vnet service endpoint enabled.
	IgnoreMissingVnetServiceEndpoint *bool `json:"ignoreMissingVnetServiceEndpoint,omitempty"`

	// Name: Resource name.
	Name *string `json:"name,omitempty"`

	// State: Virtual Network Rule State
	State *VirtualNetworkRuleProperties_State_STATUS `json:"state,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`

	// VirtualNetworkSubnetId: The ARM resource id of the virtual network subnet.
	VirtualNetworkSubnetId *string `json:"virtualNetworkSubnetId,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Servers_VirtualNetworkRule_STATUS{}

// ConvertStatusFrom populates our Servers_VirtualNetworkRule_STATUS from the provided source
func (rule *Servers_VirtualNetworkRule_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20211101s.Servers_VirtualNetworkRule_STATUS)
	if ok {
		// Populate our instance from source
		return rule.AssignProperties_From_Servers_VirtualNetworkRule_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v20211101s.Servers_VirtualNetworkRule_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = rule.AssignProperties_From_Servers_VirtualNetworkRule_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our Servers_VirtualNetworkRule_STATUS
func (rule *Servers_VirtualNetworkRule_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20211101s.Servers_VirtualNetworkRule_STATUS)
	if ok {
		// Populate destination from our instance
		return rule.AssignProperties_To_Servers_VirtualNetworkRule_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v20211101s.Servers_VirtualNetworkRule_STATUS{}
	err := rule.AssignProperties_To_Servers_VirtualNetworkRule_STATUS(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusTo()")
	}

	// Update dst from our instance
	err = dst.ConvertStatusTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusTo()")
	}

	return nil
}

var _ genruntime.FromARMConverter = &Servers_VirtualNetworkRule_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (rule *Servers_VirtualNetworkRule_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Servers_VirtualNetworkRule_STATUS_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (rule *Servers_VirtualNetworkRule_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Servers_VirtualNetworkRule_STATUS_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Servers_VirtualNetworkRule_STATUS_ARM, got %T", armInput)
	}

	// no assignment for property ‘Conditions’

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		rule.Id = &id
	}

	// Set property ‘IgnoreMissingVnetServiceEndpoint’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.IgnoreMissingVnetServiceEndpoint != nil {
			ignoreMissingVnetServiceEndpoint := *typedInput.Properties.IgnoreMissingVnetServiceEndpoint
			rule.IgnoreMissingVnetServiceEndpoint = &ignoreMissingVnetServiceEndpoint
		}
	}

	// Set property ‘Name’:
	if typedInput.Name != nil {
		name := *typedInput.Name
		rule.Name = &name
	}

	// Set property ‘State’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.State != nil {
			state := *typedInput.Properties.State
			rule.State = &state
		}
	}

	// Set property ‘Type’:
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		rule.Type = &typeVar
	}

	// Set property ‘VirtualNetworkSubnetId’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.VirtualNetworkSubnetId != nil {
			virtualNetworkSubnetId := *typedInput.Properties.VirtualNetworkSubnetId
			rule.VirtualNetworkSubnetId = &virtualNetworkSubnetId
		}
	}

	// No error
	return nil
}

// AssignProperties_From_Servers_VirtualNetworkRule_STATUS populates our Servers_VirtualNetworkRule_STATUS from the provided source Servers_VirtualNetworkRule_STATUS
func (rule *Servers_VirtualNetworkRule_STATUS) AssignProperties_From_Servers_VirtualNetworkRule_STATUS(source *v20211101s.Servers_VirtualNetworkRule_STATUS) error {

	// Conditions
	rule.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	rule.Id = genruntime.ClonePointerToString(source.Id)

	// IgnoreMissingVnetServiceEndpoint
	if source.IgnoreMissingVnetServiceEndpoint != nil {
		ignoreMissingVnetServiceEndpoint := *source.IgnoreMissingVnetServiceEndpoint
		rule.IgnoreMissingVnetServiceEndpoint = &ignoreMissingVnetServiceEndpoint
	} else {
		rule.IgnoreMissingVnetServiceEndpoint = nil
	}

	// Name
	rule.Name = genruntime.ClonePointerToString(source.Name)

	// State
	if source.State != nil {
		state := VirtualNetworkRuleProperties_State_STATUS(*source.State)
		rule.State = &state
	} else {
		rule.State = nil
	}

	// Type
	rule.Type = genruntime.ClonePointerToString(source.Type)

	// VirtualNetworkSubnetId
	rule.VirtualNetworkSubnetId = genruntime.ClonePointerToString(source.VirtualNetworkSubnetId)

	// No error
	return nil
}

// AssignProperties_To_Servers_VirtualNetworkRule_STATUS populates the provided destination Servers_VirtualNetworkRule_STATUS from our Servers_VirtualNetworkRule_STATUS
func (rule *Servers_VirtualNetworkRule_STATUS) AssignProperties_To_Servers_VirtualNetworkRule_STATUS(destination *v20211101s.Servers_VirtualNetworkRule_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(rule.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(rule.Id)

	// IgnoreMissingVnetServiceEndpoint
	if rule.IgnoreMissingVnetServiceEndpoint != nil {
		ignoreMissingVnetServiceEndpoint := *rule.IgnoreMissingVnetServiceEndpoint
		destination.IgnoreMissingVnetServiceEndpoint = &ignoreMissingVnetServiceEndpoint
	} else {
		destination.IgnoreMissingVnetServiceEndpoint = nil
	}

	// Name
	destination.Name = genruntime.ClonePointerToString(rule.Name)

	// State
	if rule.State != nil {
		state := string(*rule.State)
		destination.State = &state
	} else {
		destination.State = nil
	}

	// Type
	destination.Type = genruntime.ClonePointerToString(rule.Type)

	// VirtualNetworkSubnetId
	destination.VirtualNetworkSubnetId = genruntime.ClonePointerToString(rule.VirtualNetworkSubnetId)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

type VirtualNetworkRuleProperties_State_STATUS string

const (
	VirtualNetworkRuleProperties_State_STATUS_Deleting     = VirtualNetworkRuleProperties_State_STATUS("Deleting")
	VirtualNetworkRuleProperties_State_STATUS_Failed       = VirtualNetworkRuleProperties_State_STATUS("Failed")
	VirtualNetworkRuleProperties_State_STATUS_InProgress   = VirtualNetworkRuleProperties_State_STATUS("InProgress")
	VirtualNetworkRuleProperties_State_STATUS_Initializing = VirtualNetworkRuleProperties_State_STATUS("Initializing")
	VirtualNetworkRuleProperties_State_STATUS_Ready        = VirtualNetworkRuleProperties_State_STATUS("Ready")
	VirtualNetworkRuleProperties_State_STATUS_Unknown      = VirtualNetworkRuleProperties_State_STATUS("Unknown")
)

func init() {
	SchemeBuilder.Register(&ServersVirtualNetworkRule{}, &ServersVirtualNetworkRuleList{})
}
