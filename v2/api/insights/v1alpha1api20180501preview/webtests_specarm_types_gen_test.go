// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20180501preview

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

func Test_Webtests_SPECARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Webtests_SPECARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWebtests_SPECARM, Webtests_SPECARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWebtests_SPECARM runs a test to see if a specific instance of Webtests_SPECARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWebtests_SPECARM(subject Webtests_SPECARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Webtests_SPECARM
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

// Generator of Webtests_SPECARM instances for property testing - lazily instantiated by Webtests_SPECARMGenerator()
var webtests_specarmGenerator gopter.Gen

// Webtests_SPECARMGenerator returns a generator of Webtests_SPECARM instances for property testing.
// We first initialize webtests_specarmGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Webtests_SPECARMGenerator() gopter.Gen {
	if webtests_specarmGenerator != nil {
		return webtests_specarmGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWebtests_SPECARM(generators)
	webtests_specarmGenerator = gen.Struct(reflect.TypeOf(Webtests_SPECARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWebtests_SPECARM(generators)
	AddRelatedPropertyGeneratorsForWebtests_SPECARM(generators)
	webtests_specarmGenerator = gen.Struct(reflect.TypeOf(Webtests_SPECARM{}), generators)

	return webtests_specarmGenerator
}

// AddIndependentPropertyGeneratorsForWebtests_SPECARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWebtests_SPECARM(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.AlphaString()
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForWebtests_SPECARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForWebtests_SPECARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(WebTestProperties_SpecARMGenerator())
}

func Test_WebTestProperties_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WebTestProperties_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWebTestProperties_SpecARM, WebTestProperties_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWebTestProperties_SpecARM runs a test to see if a specific instance of WebTestProperties_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWebTestProperties_SpecARM(subject WebTestProperties_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WebTestProperties_SpecARM
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

// Generator of WebTestProperties_SpecARM instances for property testing - lazily instantiated by
//WebTestProperties_SpecARMGenerator()
var webTestProperties_specARMGenerator gopter.Gen

// WebTestProperties_SpecARMGenerator returns a generator of WebTestProperties_SpecARM instances for property testing.
// We first initialize webTestProperties_specARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func WebTestProperties_SpecARMGenerator() gopter.Gen {
	if webTestProperties_specARMGenerator != nil {
		return webTestProperties_specARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWebTestProperties_SpecARM(generators)
	webTestProperties_specARMGenerator = gen.Struct(reflect.TypeOf(WebTestProperties_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWebTestProperties_SpecARM(generators)
	AddRelatedPropertyGeneratorsForWebTestProperties_SpecARM(generators)
	webTestProperties_specARMGenerator = gen.Struct(reflect.TypeOf(WebTestProperties_SpecARM{}), generators)

	return webTestProperties_specARMGenerator
}

// AddIndependentPropertyGeneratorsForWebTestProperties_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWebTestProperties_SpecARM(gens map[string]gopter.Gen) {
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["Frequency"] = gen.PtrOf(gen.Int())
	gens["Kind"] = gen.OneConstOf(
		WebTestProperties_Kind_SpecBasic,
		WebTestProperties_Kind_SpecMultistep,
		WebTestProperties_Kind_SpecPing,
		WebTestProperties_Kind_SpecStandard)
	gens["Name"] = gen.AlphaString()
	gens["RetryEnabled"] = gen.PtrOf(gen.Bool())
	gens["SyntheticMonitorId"] = gen.AlphaString()
	gens["Timeout"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForWebTestProperties_SpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForWebTestProperties_SpecARM(gens map[string]gopter.Gen) {
	gens["Configuration"] = gen.PtrOf(WebTestProperties_Configuration_SpecARMGenerator())
	gens["Locations"] = gen.SliceOf(WebTestGeolocation_SpecARMGenerator())
	gens["Request"] = gen.PtrOf(WebTestProperties_Request_SpecARMGenerator())
	gens["ValidationRules"] = gen.PtrOf(WebTestProperties_ValidationRules_SpecARMGenerator())
}

func Test_WebTestGeolocation_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WebTestGeolocation_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWebTestGeolocation_SpecARM, WebTestGeolocation_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWebTestGeolocation_SpecARM runs a test to see if a specific instance of WebTestGeolocation_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWebTestGeolocation_SpecARM(subject WebTestGeolocation_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WebTestGeolocation_SpecARM
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

// Generator of WebTestGeolocation_SpecARM instances for property testing - lazily instantiated by
//WebTestGeolocation_SpecARMGenerator()
var webTestGeolocation_specARMGenerator gopter.Gen

// WebTestGeolocation_SpecARMGenerator returns a generator of WebTestGeolocation_SpecARM instances for property testing.
func WebTestGeolocation_SpecARMGenerator() gopter.Gen {
	if webTestGeolocation_specARMGenerator != nil {
		return webTestGeolocation_specARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWebTestGeolocation_SpecARM(generators)
	webTestGeolocation_specARMGenerator = gen.Struct(reflect.TypeOf(WebTestGeolocation_SpecARM{}), generators)

	return webTestGeolocation_specARMGenerator
}

// AddIndependentPropertyGeneratorsForWebTestGeolocation_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWebTestGeolocation_SpecARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_WebTestProperties_Configuration_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WebTestProperties_Configuration_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWebTestProperties_Configuration_SpecARM, WebTestProperties_Configuration_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWebTestProperties_Configuration_SpecARM runs a test to see if a specific instance of WebTestProperties_Configuration_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWebTestProperties_Configuration_SpecARM(subject WebTestProperties_Configuration_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WebTestProperties_Configuration_SpecARM
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

// Generator of WebTestProperties_Configuration_SpecARM instances for property testing - lazily instantiated by
//WebTestProperties_Configuration_SpecARMGenerator()
var webTestProperties_configuration_specARMGenerator gopter.Gen

// WebTestProperties_Configuration_SpecARMGenerator returns a generator of WebTestProperties_Configuration_SpecARM instances for property testing.
func WebTestProperties_Configuration_SpecARMGenerator() gopter.Gen {
	if webTestProperties_configuration_specARMGenerator != nil {
		return webTestProperties_configuration_specARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWebTestProperties_Configuration_SpecARM(generators)
	webTestProperties_configuration_specARMGenerator = gen.Struct(reflect.TypeOf(WebTestProperties_Configuration_SpecARM{}), generators)

	return webTestProperties_configuration_specARMGenerator
}

// AddIndependentPropertyGeneratorsForWebTestProperties_Configuration_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWebTestProperties_Configuration_SpecARM(gens map[string]gopter.Gen) {
	gens["WebTest"] = gen.PtrOf(gen.AlphaString())
}

func Test_WebTestProperties_Request_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WebTestProperties_Request_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWebTestProperties_Request_SpecARM, WebTestProperties_Request_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWebTestProperties_Request_SpecARM runs a test to see if a specific instance of WebTestProperties_Request_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWebTestProperties_Request_SpecARM(subject WebTestProperties_Request_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WebTestProperties_Request_SpecARM
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

// Generator of WebTestProperties_Request_SpecARM instances for property testing - lazily instantiated by
//WebTestProperties_Request_SpecARMGenerator()
var webTestProperties_request_specARMGenerator gopter.Gen

// WebTestProperties_Request_SpecARMGenerator returns a generator of WebTestProperties_Request_SpecARM instances for property testing.
// We first initialize webTestProperties_request_specARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func WebTestProperties_Request_SpecARMGenerator() gopter.Gen {
	if webTestProperties_request_specARMGenerator != nil {
		return webTestProperties_request_specARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWebTestProperties_Request_SpecARM(generators)
	webTestProperties_request_specARMGenerator = gen.Struct(reflect.TypeOf(WebTestProperties_Request_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWebTestProperties_Request_SpecARM(generators)
	AddRelatedPropertyGeneratorsForWebTestProperties_Request_SpecARM(generators)
	webTestProperties_request_specARMGenerator = gen.Struct(reflect.TypeOf(WebTestProperties_Request_SpecARM{}), generators)

	return webTestProperties_request_specARMGenerator
}

// AddIndependentPropertyGeneratorsForWebTestProperties_Request_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWebTestProperties_Request_SpecARM(gens map[string]gopter.Gen) {
	gens["FollowRedirects"] = gen.PtrOf(gen.Bool())
	gens["HttpVerb"] = gen.PtrOf(gen.AlphaString())
	gens["ParseDependentRequests"] = gen.PtrOf(gen.Bool())
	gens["RequestBody"] = gen.PtrOf(gen.AlphaString())
	gens["RequestUrl"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForWebTestProperties_Request_SpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForWebTestProperties_Request_SpecARM(gens map[string]gopter.Gen) {
	gens["Headers"] = gen.SliceOf(HeaderField_SpecARMGenerator())
}

func Test_WebTestProperties_ValidationRules_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WebTestProperties_ValidationRules_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWebTestProperties_ValidationRules_SpecARM, WebTestProperties_ValidationRules_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWebTestProperties_ValidationRules_SpecARM runs a test to see if a specific instance of WebTestProperties_ValidationRules_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWebTestProperties_ValidationRules_SpecARM(subject WebTestProperties_ValidationRules_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WebTestProperties_ValidationRules_SpecARM
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

// Generator of WebTestProperties_ValidationRules_SpecARM instances for property testing - lazily instantiated by
//WebTestProperties_ValidationRules_SpecARMGenerator()
var webTestProperties_validationRules_specARMGenerator gopter.Gen

// WebTestProperties_ValidationRules_SpecARMGenerator returns a generator of WebTestProperties_ValidationRules_SpecARM instances for property testing.
// We first initialize webTestProperties_validationRules_specARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func WebTestProperties_ValidationRules_SpecARMGenerator() gopter.Gen {
	if webTestProperties_validationRules_specARMGenerator != nil {
		return webTestProperties_validationRules_specARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWebTestProperties_ValidationRules_SpecARM(generators)
	webTestProperties_validationRules_specARMGenerator = gen.Struct(reflect.TypeOf(WebTestProperties_ValidationRules_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWebTestProperties_ValidationRules_SpecARM(generators)
	AddRelatedPropertyGeneratorsForWebTestProperties_ValidationRules_SpecARM(generators)
	webTestProperties_validationRules_specARMGenerator = gen.Struct(reflect.TypeOf(WebTestProperties_ValidationRules_SpecARM{}), generators)

	return webTestProperties_validationRules_specARMGenerator
}

// AddIndependentPropertyGeneratorsForWebTestProperties_ValidationRules_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWebTestProperties_ValidationRules_SpecARM(gens map[string]gopter.Gen) {
	gens["ExpectedHttpStatusCode"] = gen.PtrOf(gen.Int())
	gens["IgnoreHttpsStatusCode"] = gen.PtrOf(gen.Bool())
	gens["SSLCertRemainingLifetimeCheck"] = gen.PtrOf(gen.Int())
	gens["SSLCheck"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForWebTestProperties_ValidationRules_SpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForWebTestProperties_ValidationRules_SpecARM(gens map[string]gopter.Gen) {
	gens["ContentValidation"] = gen.PtrOf(WebTestProperties_ValidationRules_ContentValidation_SpecARMGenerator())
}

func Test_HeaderField_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of HeaderField_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForHeaderField_SpecARM, HeaderField_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForHeaderField_SpecARM runs a test to see if a specific instance of HeaderField_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForHeaderField_SpecARM(subject HeaderField_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual HeaderField_SpecARM
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

// Generator of HeaderField_SpecARM instances for property testing - lazily instantiated by
//HeaderField_SpecARMGenerator()
var headerField_specARMGenerator gopter.Gen

// HeaderField_SpecARMGenerator returns a generator of HeaderField_SpecARM instances for property testing.
func HeaderField_SpecARMGenerator() gopter.Gen {
	if headerField_specARMGenerator != nil {
		return headerField_specARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForHeaderField_SpecARM(generators)
	headerField_specARMGenerator = gen.Struct(reflect.TypeOf(HeaderField_SpecARM{}), generators)

	return headerField_specARMGenerator
}

// AddIndependentPropertyGeneratorsForHeaderField_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForHeaderField_SpecARM(gens map[string]gopter.Gen) {
	gens["Key"] = gen.PtrOf(gen.AlphaString())
	gens["Value"] = gen.PtrOf(gen.AlphaString())
}

func Test_WebTestProperties_ValidationRules_ContentValidation_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WebTestProperties_ValidationRules_ContentValidation_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWebTestProperties_ValidationRules_ContentValidation_SpecARM, WebTestProperties_ValidationRules_ContentValidation_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWebTestProperties_ValidationRules_ContentValidation_SpecARM runs a test to see if a specific instance of WebTestProperties_ValidationRules_ContentValidation_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWebTestProperties_ValidationRules_ContentValidation_SpecARM(subject WebTestProperties_ValidationRules_ContentValidation_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WebTestProperties_ValidationRules_ContentValidation_SpecARM
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

// Generator of WebTestProperties_ValidationRules_ContentValidation_SpecARM instances for property testing - lazily
//instantiated by WebTestProperties_ValidationRules_ContentValidation_SpecARMGenerator()
var webTestProperties_validationRules_contentValidation_specARMGenerator gopter.Gen

// WebTestProperties_ValidationRules_ContentValidation_SpecARMGenerator returns a generator of WebTestProperties_ValidationRules_ContentValidation_SpecARM instances for property testing.
func WebTestProperties_ValidationRules_ContentValidation_SpecARMGenerator() gopter.Gen {
	if webTestProperties_validationRules_contentValidation_specARMGenerator != nil {
		return webTestProperties_validationRules_contentValidation_specARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWebTestProperties_ValidationRules_ContentValidation_SpecARM(generators)
	webTestProperties_validationRules_contentValidation_specARMGenerator = gen.Struct(reflect.TypeOf(WebTestProperties_ValidationRules_ContentValidation_SpecARM{}), generators)

	return webTestProperties_validationRules_contentValidation_specARMGenerator
}

// AddIndependentPropertyGeneratorsForWebTestProperties_ValidationRules_ContentValidation_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWebTestProperties_ValidationRules_ContentValidation_SpecARM(gens map[string]gopter.Gen) {
	gens["ContentMatch"] = gen.PtrOf(gen.AlphaString())
	gens["IgnoreCase"] = gen.PtrOf(gen.Bool())
	gens["PassIfTextFound"] = gen.PtrOf(gen.Bool())
}