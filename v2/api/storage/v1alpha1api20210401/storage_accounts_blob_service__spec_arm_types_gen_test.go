// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210401

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

func Test_StorageAccountsBlobService_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccountsBlobService_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccountsBlobService_SpecARM, StorageAccountsBlobService_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccountsBlobService_SpecARM runs a test to see if a specific instance of StorageAccountsBlobService_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccountsBlobService_SpecARM(subject StorageAccountsBlobService_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccountsBlobService_SpecARM
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

// Generator of StorageAccountsBlobService_SpecARM instances for property testing - lazily instantiated by
//StorageAccountsBlobService_SpecARMGenerator()
var storageAccountsBlobService_specARMGenerator gopter.Gen

// StorageAccountsBlobService_SpecARMGenerator returns a generator of StorageAccountsBlobService_SpecARM instances for property testing.
// We first initialize storageAccountsBlobService_specARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func StorageAccountsBlobService_SpecARMGenerator() gopter.Gen {
	if storageAccountsBlobService_specARMGenerator != nil {
		return storageAccountsBlobService_specARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccountsBlobService_SpecARM(generators)
	storageAccountsBlobService_specARMGenerator = gen.Struct(reflect.TypeOf(StorageAccountsBlobService_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccountsBlobService_SpecARM(generators)
	AddRelatedPropertyGeneratorsForStorageAccountsBlobService_SpecARM(generators)
	storageAccountsBlobService_specARMGenerator = gen.Struct(reflect.TypeOf(StorageAccountsBlobService_SpecARM{}), generators)

	return storageAccountsBlobService_specARMGenerator
}

// AddIndependentPropertyGeneratorsForStorageAccountsBlobService_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageAccountsBlobService_SpecARM(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForStorageAccountsBlobService_SpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccountsBlobService_SpecARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(StorageAccountsBlobService_SpecPropertiesARMGenerator())
}

func Test_StorageAccountsBlobService_SpecPropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccountsBlobService_SpecPropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccountsBlobService_SpecPropertiesARM, StorageAccountsBlobService_SpecPropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccountsBlobService_SpecPropertiesARM runs a test to see if a specific instance of StorageAccountsBlobService_SpecPropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccountsBlobService_SpecPropertiesARM(subject StorageAccountsBlobService_SpecPropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccountsBlobService_SpecPropertiesARM
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

// Generator of StorageAccountsBlobService_SpecPropertiesARM instances for property testing - lazily instantiated by
//StorageAccountsBlobService_SpecPropertiesARMGenerator()
var storageAccountsBlobService_specPropertiesARMGenerator gopter.Gen

// StorageAccountsBlobService_SpecPropertiesARMGenerator returns a generator of StorageAccountsBlobService_SpecPropertiesARM instances for property testing.
// We first initialize storageAccountsBlobService_specPropertiesARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func StorageAccountsBlobService_SpecPropertiesARMGenerator() gopter.Gen {
	if storageAccountsBlobService_specPropertiesARMGenerator != nil {
		return storageAccountsBlobService_specPropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccountsBlobService_SpecPropertiesARM(generators)
	storageAccountsBlobService_specPropertiesARMGenerator = gen.Struct(reflect.TypeOf(StorageAccountsBlobService_SpecPropertiesARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccountsBlobService_SpecPropertiesARM(generators)
	AddRelatedPropertyGeneratorsForStorageAccountsBlobService_SpecPropertiesARM(generators)
	storageAccountsBlobService_specPropertiesARMGenerator = gen.Struct(reflect.TypeOf(StorageAccountsBlobService_SpecPropertiesARM{}), generators)

	return storageAccountsBlobService_specPropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForStorageAccountsBlobService_SpecPropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageAccountsBlobService_SpecPropertiesARM(gens map[string]gopter.Gen) {
	gens["AutomaticSnapshotPolicyEnabled"] = gen.PtrOf(gen.Bool())
	gens["DefaultServiceVersion"] = gen.PtrOf(gen.AlphaString())
	gens["IsVersioningEnabled"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForStorageAccountsBlobService_SpecPropertiesARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccountsBlobService_SpecPropertiesARM(gens map[string]gopter.Gen) {
	gens["ChangeFeed"] = gen.PtrOf(ChangeFeedARMGenerator())
	gens["ContainerDeleteRetentionPolicy"] = gen.PtrOf(DeleteRetentionPolicyARMGenerator())
	gens["Cors"] = gen.PtrOf(CorsRulesARMGenerator())
	gens["DeleteRetentionPolicy"] = gen.PtrOf(DeleteRetentionPolicyARMGenerator())
	gens["LastAccessTimeTrackingPolicy"] = gen.PtrOf(LastAccessTimeTrackingPolicyARMGenerator())
	gens["RestorePolicy"] = gen.PtrOf(RestorePolicyPropertiesARMGenerator())
}

func Test_ChangeFeedARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ChangeFeedARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForChangeFeedARM, ChangeFeedARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForChangeFeedARM runs a test to see if a specific instance of ChangeFeedARM round trips to JSON and back losslessly
func RunJSONSerializationTestForChangeFeedARM(subject ChangeFeedARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ChangeFeedARM
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

// Generator of ChangeFeedARM instances for property testing - lazily instantiated by ChangeFeedARMGenerator()
var changeFeedARMGenerator gopter.Gen

// ChangeFeedARMGenerator returns a generator of ChangeFeedARM instances for property testing.
func ChangeFeedARMGenerator() gopter.Gen {
	if changeFeedARMGenerator != nil {
		return changeFeedARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForChangeFeedARM(generators)
	changeFeedARMGenerator = gen.Struct(reflect.TypeOf(ChangeFeedARM{}), generators)

	return changeFeedARMGenerator
}

// AddIndependentPropertyGeneratorsForChangeFeedARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForChangeFeedARM(gens map[string]gopter.Gen) {
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["RetentionInDays"] = gen.PtrOf(gen.Int())
}

func Test_CorsRulesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CorsRulesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCorsRulesARM, CorsRulesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCorsRulesARM runs a test to see if a specific instance of CorsRulesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForCorsRulesARM(subject CorsRulesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CorsRulesARM
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

// Generator of CorsRulesARM instances for property testing - lazily instantiated by CorsRulesARMGenerator()
var corsRulesARMGenerator gopter.Gen

// CorsRulesARMGenerator returns a generator of CorsRulesARM instances for property testing.
func CorsRulesARMGenerator() gopter.Gen {
	if corsRulesARMGenerator != nil {
		return corsRulesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForCorsRulesARM(generators)
	corsRulesARMGenerator = gen.Struct(reflect.TypeOf(CorsRulesARM{}), generators)

	return corsRulesARMGenerator
}

// AddRelatedPropertyGeneratorsForCorsRulesARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForCorsRulesARM(gens map[string]gopter.Gen) {
	gens["CorsRules"] = gen.SliceOf(CorsRuleARMGenerator())
}

func Test_DeleteRetentionPolicyARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DeleteRetentionPolicyARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDeleteRetentionPolicyARM, DeleteRetentionPolicyARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDeleteRetentionPolicyARM runs a test to see if a specific instance of DeleteRetentionPolicyARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDeleteRetentionPolicyARM(subject DeleteRetentionPolicyARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DeleteRetentionPolicyARM
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

// Generator of DeleteRetentionPolicyARM instances for property testing - lazily instantiated by
//DeleteRetentionPolicyARMGenerator()
var deleteRetentionPolicyARMGenerator gopter.Gen

// DeleteRetentionPolicyARMGenerator returns a generator of DeleteRetentionPolicyARM instances for property testing.
func DeleteRetentionPolicyARMGenerator() gopter.Gen {
	if deleteRetentionPolicyARMGenerator != nil {
		return deleteRetentionPolicyARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDeleteRetentionPolicyARM(generators)
	deleteRetentionPolicyARMGenerator = gen.Struct(reflect.TypeOf(DeleteRetentionPolicyARM{}), generators)

	return deleteRetentionPolicyARMGenerator
}

// AddIndependentPropertyGeneratorsForDeleteRetentionPolicyARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDeleteRetentionPolicyARM(gens map[string]gopter.Gen) {
	gens["Days"] = gen.PtrOf(gen.Int())
	gens["Enabled"] = gen.PtrOf(gen.Bool())
}

func Test_LastAccessTimeTrackingPolicyARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of LastAccessTimeTrackingPolicyARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForLastAccessTimeTrackingPolicyARM, LastAccessTimeTrackingPolicyARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForLastAccessTimeTrackingPolicyARM runs a test to see if a specific instance of LastAccessTimeTrackingPolicyARM round trips to JSON and back losslessly
func RunJSONSerializationTestForLastAccessTimeTrackingPolicyARM(subject LastAccessTimeTrackingPolicyARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual LastAccessTimeTrackingPolicyARM
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

// Generator of LastAccessTimeTrackingPolicyARM instances for property testing - lazily instantiated by
//LastAccessTimeTrackingPolicyARMGenerator()
var lastAccessTimeTrackingPolicyARMGenerator gopter.Gen

// LastAccessTimeTrackingPolicyARMGenerator returns a generator of LastAccessTimeTrackingPolicyARM instances for property testing.
func LastAccessTimeTrackingPolicyARMGenerator() gopter.Gen {
	if lastAccessTimeTrackingPolicyARMGenerator != nil {
		return lastAccessTimeTrackingPolicyARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLastAccessTimeTrackingPolicyARM(generators)
	lastAccessTimeTrackingPolicyARMGenerator = gen.Struct(reflect.TypeOf(LastAccessTimeTrackingPolicyARM{}), generators)

	return lastAccessTimeTrackingPolicyARMGenerator
}

// AddIndependentPropertyGeneratorsForLastAccessTimeTrackingPolicyARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForLastAccessTimeTrackingPolicyARM(gens map[string]gopter.Gen) {
	gens["BlobType"] = gen.SliceOf(gen.AlphaString())
	gens["Enable"] = gen.Bool()
	gens["Name"] = gen.PtrOf(gen.OneConstOf(LastAccessTimeTrackingPolicyNameAccessTimeTracking))
	gens["TrackingGranularityInDays"] = gen.PtrOf(gen.Int())
}

func Test_RestorePolicyPropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RestorePolicyPropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRestorePolicyPropertiesARM, RestorePolicyPropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRestorePolicyPropertiesARM runs a test to see if a specific instance of RestorePolicyPropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRestorePolicyPropertiesARM(subject RestorePolicyPropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RestorePolicyPropertiesARM
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

// Generator of RestorePolicyPropertiesARM instances for property testing - lazily instantiated by
//RestorePolicyPropertiesARMGenerator()
var restorePolicyPropertiesARMGenerator gopter.Gen

// RestorePolicyPropertiesARMGenerator returns a generator of RestorePolicyPropertiesARM instances for property testing.
func RestorePolicyPropertiesARMGenerator() gopter.Gen {
	if restorePolicyPropertiesARMGenerator != nil {
		return restorePolicyPropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRestorePolicyPropertiesARM(generators)
	restorePolicyPropertiesARMGenerator = gen.Struct(reflect.TypeOf(RestorePolicyPropertiesARM{}), generators)

	return restorePolicyPropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForRestorePolicyPropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRestorePolicyPropertiesARM(gens map[string]gopter.Gen) {
	gens["Days"] = gen.PtrOf(gen.Int())
	gens["Enabled"] = gen.Bool()
}

func Test_CorsRuleARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CorsRuleARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCorsRuleARM, CorsRuleARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCorsRuleARM runs a test to see if a specific instance of CorsRuleARM round trips to JSON and back losslessly
func RunJSONSerializationTestForCorsRuleARM(subject CorsRuleARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CorsRuleARM
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

// Generator of CorsRuleARM instances for property testing - lazily instantiated by CorsRuleARMGenerator()
var corsRuleARMGenerator gopter.Gen

// CorsRuleARMGenerator returns a generator of CorsRuleARM instances for property testing.
func CorsRuleARMGenerator() gopter.Gen {
	if corsRuleARMGenerator != nil {
		return corsRuleARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCorsRuleARM(generators)
	corsRuleARMGenerator = gen.Struct(reflect.TypeOf(CorsRuleARM{}), generators)

	return corsRuleARMGenerator
}

// AddIndependentPropertyGeneratorsForCorsRuleARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCorsRuleARM(gens map[string]gopter.Gen) {
	gens["AllowedHeaders"] = gen.SliceOf(gen.AlphaString())
	gens["AllowedMethods"] = gen.SliceOf(gen.OneConstOf(
		CorsRuleAllowedMethodsDELETE,
		CorsRuleAllowedMethodsGET,
		CorsRuleAllowedMethodsHEAD,
		CorsRuleAllowedMethodsMERGE,
		CorsRuleAllowedMethodsOPTIONS,
		CorsRuleAllowedMethodsPOST,
		CorsRuleAllowedMethodsPUT))
	gens["AllowedOrigins"] = gen.SliceOf(gen.AlphaString())
	gens["ExposedHeaders"] = gen.SliceOf(gen.AlphaString())
	gens["MaxAgeInSeconds"] = gen.Int()
}