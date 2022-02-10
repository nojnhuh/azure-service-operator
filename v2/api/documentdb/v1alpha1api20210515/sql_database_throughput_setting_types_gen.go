// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515

import (
	"fmt"
	"github.com/Azure/azure-service-operator/v2/api/documentdb/v1alpha1api20210515storage"
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
type SqlDatabaseThroughputSetting struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseAccountsSqlDatabasesThroughputSettings_SPEC `json:"spec,omitempty"`
	Status            ThroughputSettingsUpdateParameters_Status           `json:"status,omitempty"`
}

var _ conditions.Conditioner = &SqlDatabaseThroughputSetting{}

// GetConditions returns the conditions of the resource
func (setting *SqlDatabaseThroughputSetting) GetConditions() conditions.Conditions {
	return setting.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (setting *SqlDatabaseThroughputSetting) SetConditions(conditions conditions.Conditions) {
	setting.Status.Conditions = conditions
}

var _ conversion.Convertible = &SqlDatabaseThroughputSetting{}

// ConvertFrom populates our SqlDatabaseThroughputSetting from the provided hub SqlDatabaseThroughputSetting
func (setting *SqlDatabaseThroughputSetting) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v1alpha1api20210515storage.SqlDatabaseThroughputSetting)
	if !ok {
		return fmt.Errorf("expected storage:documentdb/v1alpha1api20210515storage/SqlDatabaseThroughputSetting but received %T instead", hub)
	}

	return setting.AssignPropertiesFromSqlDatabaseThroughputSetting(source)
}

// ConvertTo populates the provided hub SqlDatabaseThroughputSetting from our SqlDatabaseThroughputSetting
func (setting *SqlDatabaseThroughputSetting) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v1alpha1api20210515storage.SqlDatabaseThroughputSetting)
	if !ok {
		return fmt.Errorf("expected storage:documentdb/v1alpha1api20210515storage/SqlDatabaseThroughputSetting but received %T instead", hub)
	}

	return setting.AssignPropertiesToSqlDatabaseThroughputSetting(destination)
}

// +kubebuilder:webhook:path=/mutate-documentdb-azure-com-v1alpha1api20210515-sqldatabasethroughputsetting,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=documentdb.azure.com,resources=sqldatabasethroughputsettings,verbs=create;update,versions=v1alpha1api20210515,name=default.v1alpha1api20210515.sqldatabasethroughputsettings.documentdb.azure.com,admissionReviewVersions=v1beta1

var _ admission.Defaulter = &SqlDatabaseThroughputSetting{}

// Default applies defaults to the SqlDatabaseThroughputSetting resource
func (setting *SqlDatabaseThroughputSetting) Default() {
	setting.defaultImpl()
	var temp interface{} = setting
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (setting *SqlDatabaseThroughputSetting) defaultAzureName() {
	if setting.Spec.AzureName == "" {
		setting.Spec.AzureName = setting.Name
	}
}

// defaultImpl applies the code generated defaults to the SqlDatabaseThroughputSetting resource
func (setting *SqlDatabaseThroughputSetting) defaultImpl() { setting.defaultAzureName() }

var _ genruntime.KubernetesResource = &SqlDatabaseThroughputSetting{}

// AzureName returns the Azure name of the resource
func (setting *SqlDatabaseThroughputSetting) AzureName() string {
	return setting.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (setting SqlDatabaseThroughputSetting) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetResourceKind returns the kind of the resource
func (setting *SqlDatabaseThroughputSetting) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (setting *SqlDatabaseThroughputSetting) GetSpec() genruntime.ConvertibleSpec {
	return &setting.Spec
}

// GetStatus returns the status of this resource
func (setting *SqlDatabaseThroughputSetting) GetStatus() genruntime.ConvertibleStatus {
	return &setting.Status
}

// GetType returns the ARM Type of the resource. This is always ""
func (setting *SqlDatabaseThroughputSetting) GetType() string {
	return ""
}

// NewEmptyStatus returns a new empty (blank) status
func (setting *SqlDatabaseThroughputSetting) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &ThroughputSettingsUpdateParameters_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (setting *SqlDatabaseThroughputSetting) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(setting.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  setting.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (setting *SqlDatabaseThroughputSetting) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*ThroughputSettingsUpdateParameters_Status); ok {
		setting.Status = *st
		return nil
	}

	// Convert status to required version
	var st ThroughputSettingsUpdateParameters_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	setting.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-documentdb-azure-com-v1alpha1api20210515-sqldatabasethroughputsetting,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=documentdb.azure.com,resources=sqldatabasethroughputsettings,verbs=create;update,versions=v1alpha1api20210515,name=validate.v1alpha1api20210515.sqldatabasethroughputsettings.documentdb.azure.com,admissionReviewVersions=v1beta1

var _ admission.Validator = &SqlDatabaseThroughputSetting{}

// ValidateCreate validates the creation of the resource
func (setting *SqlDatabaseThroughputSetting) ValidateCreate() error {
	validations := setting.createValidations()
	var temp interface{} = setting
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
func (setting *SqlDatabaseThroughputSetting) ValidateDelete() error {
	validations := setting.deleteValidations()
	var temp interface{} = setting
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
func (setting *SqlDatabaseThroughputSetting) ValidateUpdate(old runtime.Object) error {
	validations := setting.updateValidations()
	var temp interface{} = setting
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
func (setting *SqlDatabaseThroughputSetting) createValidations() []func() error {
	return []func() error{setting.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (setting *SqlDatabaseThroughputSetting) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (setting *SqlDatabaseThroughputSetting) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return setting.validateResourceReferences()
		},
	}
}

// validateResourceReferences validates all resource references
func (setting *SqlDatabaseThroughputSetting) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&setting.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// AssignPropertiesFromSqlDatabaseThroughputSetting populates our SqlDatabaseThroughputSetting from the provided source SqlDatabaseThroughputSetting
func (setting *SqlDatabaseThroughputSetting) AssignPropertiesFromSqlDatabaseThroughputSetting(source *v1alpha1api20210515storage.SqlDatabaseThroughputSetting) error {

	// ObjectMeta
	setting.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec DatabaseAccountsSqlDatabasesThroughputSettings_SPEC
	err := spec.AssignPropertiesFromDatabaseAccountsSqlDatabasesThroughputSettingsSPEC(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromDatabaseAccountsSqlDatabasesThroughputSettingsSPEC() to populate field Spec")
	}
	setting.Spec = spec

	// Status
	var status ThroughputSettingsUpdateParameters_Status
	err = status.AssignPropertiesFromThroughputSettingsUpdateParametersStatus(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromThroughputSettingsUpdateParametersStatus() to populate field Status")
	}
	setting.Status = status

	// No error
	return nil
}

// AssignPropertiesToSqlDatabaseThroughputSetting populates the provided destination SqlDatabaseThroughputSetting from our SqlDatabaseThroughputSetting
func (setting *SqlDatabaseThroughputSetting) AssignPropertiesToSqlDatabaseThroughputSetting(destination *v1alpha1api20210515storage.SqlDatabaseThroughputSetting) error {

	// ObjectMeta
	destination.ObjectMeta = *setting.ObjectMeta.DeepCopy()

	// Spec
	var spec v1alpha1api20210515storage.DatabaseAccountsSqlDatabasesThroughputSettings_SPEC
	err := setting.Spec.AssignPropertiesToDatabaseAccountsSqlDatabasesThroughputSettingsSPEC(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToDatabaseAccountsSqlDatabasesThroughputSettingsSPEC() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v1alpha1api20210515storage.ThroughputSettingsUpdateParameters_Status
	err = setting.Status.AssignPropertiesToThroughputSettingsUpdateParametersStatus(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToThroughputSettingsUpdateParametersStatus() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (setting *SqlDatabaseThroughputSetting) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: setting.Spec.OriginalVersion(),
		Kind:    "SqlDatabaseThroughputSetting",
	}
}

// +kubebuilder:object:root=true
type SqlDatabaseThroughputSettingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SqlDatabaseThroughputSetting `json:"items"`
}

type DatabaseAccountsSqlDatabasesThroughputSettings_SPEC struct {
	//AzureName: The name of the resource in Azure. This is often the same as the name
	//of the resource in Kubernetes but it doesn't have to be.
	AzureName string `json:"azureName"`

	//Location: The location of the resource group to which the resource belongs.
	Location *string `json:"location,omitempty"`

	// +kubebuilder:validation:Required
	Owner genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner" kind:"ResourceGroup"`

	// +kubebuilder:validation:Required
	//Resource: The standard JSON format of a resource throughput
	Resource ThroughputSettingsResource_Spec `json:"resource"`
	Tags     map[string]string               `json:"tags,omitempty"`
}

var _ genruntime.ARMTransformer = &DatabaseAccountsSqlDatabasesThroughputSettings_SPEC{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (spec *DatabaseAccountsSqlDatabasesThroughputSettings_SPEC) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if spec == nil {
		return nil, nil
	}
	var result DatabaseAccountsSqlDatabasesThroughputSettings_SPECARM

	// Set property ‘AzureName’:
	result.AzureName = spec.AzureName

	// Set property ‘Location’:
	if spec.Location != nil {
		location := *spec.Location
		result.Location = &location
	}

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	resourceARM, err := spec.Resource.ConvertToARM(resolved)
	if err != nil {
		return nil, err
	}
	result.Properties.Resource = resourceARM.(ThroughputSettingsResource_SpecARM)

	// Set property ‘Tags’:
	if spec.Tags != nil {
		result.Tags = make(map[string]string)
		for key, value := range spec.Tags {
			result.Tags[key] = value
		}
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (spec *DatabaseAccountsSqlDatabasesThroughputSettings_SPEC) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &DatabaseAccountsSqlDatabasesThroughputSettings_SPECARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (spec *DatabaseAccountsSqlDatabasesThroughputSettings_SPEC) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(DatabaseAccountsSqlDatabasesThroughputSettings_SPECARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected DatabaseAccountsSqlDatabasesThroughputSettings_SPECARM, got %T", armInput)
	}

	// Set property ‘AzureName’:
	spec.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		spec.Location = &location
	}

	// Set property ‘Owner’:
	spec.Owner = genruntime.KnownResourceReference{
		Name: owner.Name,
	}

	// Set property ‘Resource’:
	// copying flattened property:
	var resource ThroughputSettingsResource_Spec
	err := resource.PopulateFromARM(owner, typedInput.Properties.Resource)
	if err != nil {
		return err
	}
	spec.Resource = resource

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		spec.Tags = make(map[string]string)
		for key, value := range typedInput.Tags {
			spec.Tags[key] = value
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &DatabaseAccountsSqlDatabasesThroughputSettings_SPEC{}

// ConvertSpecFrom populates our DatabaseAccountsSqlDatabasesThroughputSettings_SPEC from the provided source
func (spec *DatabaseAccountsSqlDatabasesThroughputSettings_SPEC) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v1alpha1api20210515storage.DatabaseAccountsSqlDatabasesThroughputSettings_SPEC)
	if ok {
		// Populate our instance from source
		return spec.AssignPropertiesFromDatabaseAccountsSqlDatabasesThroughputSettingsSPEC(src)
	}

	// Convert to an intermediate form
	src = &v1alpha1api20210515storage.DatabaseAccountsSqlDatabasesThroughputSettings_SPEC{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = spec.AssignPropertiesFromDatabaseAccountsSqlDatabasesThroughputSettingsSPEC(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our DatabaseAccountsSqlDatabasesThroughputSettings_SPEC
func (spec *DatabaseAccountsSqlDatabasesThroughputSettings_SPEC) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v1alpha1api20210515storage.DatabaseAccountsSqlDatabasesThroughputSettings_SPEC)
	if ok {
		// Populate destination from our instance
		return spec.AssignPropertiesToDatabaseAccountsSqlDatabasesThroughputSettingsSPEC(dst)
	}

	// Convert to an intermediate form
	dst = &v1alpha1api20210515storage.DatabaseAccountsSqlDatabasesThroughputSettings_SPEC{}
	err := spec.AssignPropertiesToDatabaseAccountsSqlDatabasesThroughputSettingsSPEC(dst)
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

// AssignPropertiesFromDatabaseAccountsSqlDatabasesThroughputSettingsSPEC populates our DatabaseAccountsSqlDatabasesThroughputSettings_SPEC from the provided source DatabaseAccountsSqlDatabasesThroughputSettings_SPEC
func (spec *DatabaseAccountsSqlDatabasesThroughputSettings_SPEC) AssignPropertiesFromDatabaseAccountsSqlDatabasesThroughputSettingsSPEC(source *v1alpha1api20210515storage.DatabaseAccountsSqlDatabasesThroughputSettings_SPEC) error {

	// AzureName
	spec.AzureName = source.AzureName

	// Location
	spec.Location = genruntime.ClonePointerToString(source.Location)

	// Owner
	spec.Owner = source.Owner.Copy()

	// Resource
	if source.Resource != nil {
		var resource ThroughputSettingsResource_Spec
		err := resource.AssignPropertiesFromThroughputSettingsResourceSpec(source.Resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromThroughputSettingsResourceSpec() to populate field Resource")
		}
		spec.Resource = resource
	} else {
		spec.Resource = ThroughputSettingsResource_Spec{}
	}

	// Tags
	spec.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// No error
	return nil
}

// AssignPropertiesToDatabaseAccountsSqlDatabasesThroughputSettingsSPEC populates the provided destination DatabaseAccountsSqlDatabasesThroughputSettings_SPEC from our DatabaseAccountsSqlDatabasesThroughputSettings_SPEC
func (spec *DatabaseAccountsSqlDatabasesThroughputSettings_SPEC) AssignPropertiesToDatabaseAccountsSqlDatabasesThroughputSettingsSPEC(destination *v1alpha1api20210515storage.DatabaseAccountsSqlDatabasesThroughputSettings_SPEC) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = spec.AzureName

	// Location
	destination.Location = genruntime.ClonePointerToString(spec.Location)

	// OriginalVersion
	destination.OriginalVersion = spec.OriginalVersion()

	// Owner
	destination.Owner = spec.Owner.Copy()

	// Resource
	var resource v1alpha1api20210515storage.ThroughputSettingsResource_Spec
	err := spec.Resource.AssignPropertiesToThroughputSettingsResourceSpec(&resource)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToThroughputSettingsResourceSpec() to populate field Resource")
	}
	destination.Resource = &resource

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(spec.Tags)

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
func (spec *DatabaseAccountsSqlDatabasesThroughputSettings_SPEC) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (spec *DatabaseAccountsSqlDatabasesThroughputSettings_SPEC) SetAzureName(azureName string) {
	spec.AzureName = azureName
}

func init() {
	SchemeBuilder.Register(&SqlDatabaseThroughputSetting{}, &SqlDatabaseThroughputSettingList{})
}
