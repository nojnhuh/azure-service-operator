// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20211101

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

func Test_NamespacesAuthorizationRules_SPECARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesAuthorizationRules_SPECARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesAuthorizationRules_SPECARM, NamespacesAuthorizationRules_SPECARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesAuthorizationRules_SPECARM runs a test to see if a specific instance of NamespacesAuthorizationRules_SPECARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesAuthorizationRules_SPECARM(subject NamespacesAuthorizationRules_SPECARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesAuthorizationRules_SPECARM
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

// Generator of NamespacesAuthorizationRules_SPECARM instances for property testing - lazily instantiated by
//NamespacesAuthorizationRules_SPECARMGenerator()
var namespacesAuthorizationRules_specarmGenerator gopter.Gen

// NamespacesAuthorizationRules_SPECARMGenerator returns a generator of NamespacesAuthorizationRules_SPECARM instances for property testing.
// We first initialize namespacesAuthorizationRules_specarmGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NamespacesAuthorizationRules_SPECARMGenerator() gopter.Gen {
	if namespacesAuthorizationRules_specarmGenerator != nil {
		return namespacesAuthorizationRules_specarmGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesAuthorizationRules_SPECARM(generators)
	namespacesAuthorizationRules_specarmGenerator = gen.Struct(reflect.TypeOf(NamespacesAuthorizationRules_SPECARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesAuthorizationRules_SPECARM(generators)
	AddRelatedPropertyGeneratorsForNamespacesAuthorizationRules_SPECARM(generators)
	namespacesAuthorizationRules_specarmGenerator = gen.Struct(reflect.TypeOf(NamespacesAuthorizationRules_SPECARM{}), generators)

	return namespacesAuthorizationRules_specarmGenerator
}

// AddIndependentPropertyGeneratorsForNamespacesAuthorizationRules_SPECARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespacesAuthorizationRules_SPECARM(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForNamespacesAuthorizationRules_SPECARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesAuthorizationRules_SPECARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(NamespacesAuthorizationRules_Properties_SPECARMGenerator())
}

func Test_NamespacesAuthorizationRules_Properties_SPECARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesAuthorizationRules_Properties_SPECARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesAuthorizationRules_Properties_SPECARM, NamespacesAuthorizationRules_Properties_SPECARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesAuthorizationRules_Properties_SPECARM runs a test to see if a specific instance of NamespacesAuthorizationRules_Properties_SPECARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesAuthorizationRules_Properties_SPECARM(subject NamespacesAuthorizationRules_Properties_SPECARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesAuthorizationRules_Properties_SPECARM
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

// Generator of NamespacesAuthorizationRules_Properties_SPECARM instances for property testing - lazily instantiated by
//NamespacesAuthorizationRules_Properties_SPECARMGenerator()
var namespacesAuthorizationRules_properties_specarmGenerator gopter.Gen

// NamespacesAuthorizationRules_Properties_SPECARMGenerator returns a generator of NamespacesAuthorizationRules_Properties_SPECARM instances for property testing.
func NamespacesAuthorizationRules_Properties_SPECARMGenerator() gopter.Gen {
	if namespacesAuthorizationRules_properties_specarmGenerator != nil {
		return namespacesAuthorizationRules_properties_specarmGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesAuthorizationRules_Properties_SPECARM(generators)
	namespacesAuthorizationRules_properties_specarmGenerator = gen.Struct(reflect.TypeOf(NamespacesAuthorizationRules_Properties_SPECARM{}), generators)

	return namespacesAuthorizationRules_properties_specarmGenerator
}

// AddIndependentPropertyGeneratorsForNamespacesAuthorizationRules_Properties_SPECARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespacesAuthorizationRules_Properties_SPECARM(gens map[string]gopter.Gen) {
	gens["Rights"] = gen.SliceOf(gen.OneConstOf(NamespacesAuthorizationRules_Properties_Rights_SPECListen, NamespacesAuthorizationRules_Properties_Rights_SPECManage, NamespacesAuthorizationRules_Properties_Rights_SPECSend))
}