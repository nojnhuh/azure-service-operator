// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200801preview

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kr/pretty"
	"github.com/kylelemons/godebug/diff"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"os"
	"reflect"
	"testing"
)

func Test_RoleAssignment_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RoleAssignment_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRoleAssignment_SpecARM, RoleAssignment_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRoleAssignment_SpecARM runs a test to see if a specific instance of RoleAssignment_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRoleAssignment_SpecARM(subject RoleAssignment_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RoleAssignment_SpecARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of RoleAssignment_SpecARM instances for property testing - lazily instantiated by
//RoleAssignment_SpecARMGenerator()
var roleAssignment_specARMGenerator gopter.Gen

// RoleAssignment_SpecARMGenerator returns a generator of RoleAssignment_SpecARM instances for property testing.
// We first initialize roleAssignment_specARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RoleAssignment_SpecARMGenerator() gopter.Gen {
	if roleAssignment_specARMGenerator != nil {
		return roleAssignment_specARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRoleAssignment_SpecARM(generators)
	roleAssignment_specARMGenerator = gen.Struct(reflect.TypeOf(RoleAssignment_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRoleAssignment_SpecARM(generators)
	AddRelatedPropertyGeneratorsForRoleAssignment_SpecARM(generators)
	roleAssignment_specARMGenerator = gen.Struct(reflect.TypeOf(RoleAssignment_SpecARM{}), generators)

	return roleAssignment_specARMGenerator
}

// AddIndependentPropertyGeneratorsForRoleAssignment_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRoleAssignment_SpecARM(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForRoleAssignment_SpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRoleAssignment_SpecARM(gens map[string]gopter.Gen) {
	gens["Properties"] = RoleAssignmentPropertiesARMGenerator()
}

func Test_RoleAssignmentPropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RoleAssignmentPropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRoleAssignmentPropertiesARM, RoleAssignmentPropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRoleAssignmentPropertiesARM runs a test to see if a specific instance of RoleAssignmentPropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRoleAssignmentPropertiesARM(subject RoleAssignmentPropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RoleAssignmentPropertiesARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of RoleAssignmentPropertiesARM instances for property testing - lazily instantiated by
//RoleAssignmentPropertiesARMGenerator()
var roleAssignmentPropertiesARMGenerator gopter.Gen

// RoleAssignmentPropertiesARMGenerator returns a generator of RoleAssignmentPropertiesARM instances for property testing.
func RoleAssignmentPropertiesARMGenerator() gopter.Gen {
	if roleAssignmentPropertiesARMGenerator != nil {
		return roleAssignmentPropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRoleAssignmentPropertiesARM(generators)
	roleAssignmentPropertiesARMGenerator = gen.Struct(reflect.TypeOf(RoleAssignmentPropertiesARM{}), generators)

	return roleAssignmentPropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForRoleAssignmentPropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRoleAssignmentPropertiesARM(gens map[string]gopter.Gen) {
	gens["Condition"] = gen.PtrOf(gen.AlphaString())
	gens["ConditionVersion"] = gen.PtrOf(gen.AlphaString())
	gens["DelegatedManagedIdentityResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["PrincipalId"] = gen.AlphaString()
	gens["PrincipalType"] = gen.PtrOf(gen.OneConstOf(
		RoleAssignmentPropertiesPrincipalTypeForeignGroup,
		RoleAssignmentPropertiesPrincipalTypeGroup,
		RoleAssignmentPropertiesPrincipalTypeServicePrincipal,
		RoleAssignmentPropertiesPrincipalTypeUser))
	gens["RoleDefinitionId"] = gen.AlphaString()
}