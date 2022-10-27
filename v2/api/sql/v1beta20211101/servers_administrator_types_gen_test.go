// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211101

import (
	"encoding/json"
	v20211101s "github.com/Azure/azure-service-operator/v2/api/sql/v1beta20211101storage"
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

func Test_ServersAdministrator_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ServersAdministrator to hub returns original",
		prop.ForAll(RunResourceConversionTestForServersAdministrator, ServersAdministratorGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForServersAdministrator tests if a specific instance of ServersAdministrator round trips to the hub storage version and back losslessly
func RunResourceConversionTestForServersAdministrator(subject ServersAdministrator) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20211101s.ServersAdministrator
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual ServersAdministrator
	err = actual.ConvertFrom(&hub)
	if err != nil {
		return err.Error()
	}

	// Compare actual with what we started with
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_ServersAdministrator_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ServersAdministrator to ServersAdministrator via AssignProperties_To_ServersAdministrator & AssignProperties_From_ServersAdministrator returns original",
		prop.ForAll(RunPropertyAssignmentTestForServersAdministrator, ServersAdministratorGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForServersAdministrator tests if a specific instance of ServersAdministrator can be assigned to v1beta20211101storage and back losslessly
func RunPropertyAssignmentTestForServersAdministrator(subject ServersAdministrator) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20211101s.ServersAdministrator
	err := copied.AssignProperties_To_ServersAdministrator(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual ServersAdministrator
	err = actual.AssignProperties_From_ServersAdministrator(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_ServersAdministrator_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServersAdministrator via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServersAdministrator, ServersAdministratorGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServersAdministrator runs a test to see if a specific instance of ServersAdministrator round trips to JSON and back losslessly
func RunJSONSerializationTestForServersAdministrator(subject ServersAdministrator) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServersAdministrator
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

// Generator of ServersAdministrator instances for property testing - lazily instantiated by
// ServersAdministratorGenerator()
var serversAdministratorGenerator gopter.Gen

// ServersAdministratorGenerator returns a generator of ServersAdministrator instances for property testing.
func ServersAdministratorGenerator() gopter.Gen {
	if serversAdministratorGenerator != nil {
		return serversAdministratorGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForServersAdministrator(generators)
	serversAdministratorGenerator = gen.Struct(reflect.TypeOf(ServersAdministrator{}), generators)

	return serversAdministratorGenerator
}

// AddRelatedPropertyGeneratorsForServersAdministrator is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServersAdministrator(gens map[string]gopter.Gen) {
	gens["Spec"] = Servers_Administrator_SpecGenerator()
	gens["Status"] = Servers_Administrator_STATUSGenerator()
}

func Test_Servers_Administrator_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Servers_Administrator_Spec to Servers_Administrator_Spec via AssignProperties_To_Servers_Administrator_Spec & AssignProperties_From_Servers_Administrator_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForServers_Administrator_Spec, Servers_Administrator_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForServers_Administrator_Spec tests if a specific instance of Servers_Administrator_Spec can be assigned to v1beta20211101storage and back losslessly
func RunPropertyAssignmentTestForServers_Administrator_Spec(subject Servers_Administrator_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20211101s.Servers_Administrator_Spec
	err := copied.AssignProperties_To_Servers_Administrator_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Servers_Administrator_Spec
	err = actual.AssignProperties_From_Servers_Administrator_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Servers_Administrator_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Servers_Administrator_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServers_Administrator_Spec, Servers_Administrator_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServers_Administrator_Spec runs a test to see if a specific instance of Servers_Administrator_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForServers_Administrator_Spec(subject Servers_Administrator_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Servers_Administrator_Spec
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

// Generator of Servers_Administrator_Spec instances for property testing - lazily instantiated by
// Servers_Administrator_SpecGenerator()
var servers_Administrator_SpecGenerator gopter.Gen

// Servers_Administrator_SpecGenerator returns a generator of Servers_Administrator_Spec instances for property testing.
func Servers_Administrator_SpecGenerator() gopter.Gen {
	if servers_Administrator_SpecGenerator != nil {
		return servers_Administrator_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServers_Administrator_Spec(generators)
	servers_Administrator_SpecGenerator = gen.Struct(reflect.TypeOf(Servers_Administrator_Spec{}), generators)

	return servers_Administrator_SpecGenerator
}

// AddIndependentPropertyGeneratorsForServers_Administrator_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServers_Administrator_Spec(gens map[string]gopter.Gen) {
	gens["AdministratorType"] = gen.PtrOf(gen.OneConstOf(AdministratorProperties_AdministratorType_ActiveDirectory))
	gens["Login"] = gen.PtrOf(gen.AlphaString())
	gens["Sid"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
}

func Test_Servers_Administrator_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Servers_Administrator_STATUS to Servers_Administrator_STATUS via AssignProperties_To_Servers_Administrator_STATUS & AssignProperties_From_Servers_Administrator_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForServers_Administrator_STATUS, Servers_Administrator_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForServers_Administrator_STATUS tests if a specific instance of Servers_Administrator_STATUS can be assigned to v1beta20211101storage and back losslessly
func RunPropertyAssignmentTestForServers_Administrator_STATUS(subject Servers_Administrator_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20211101s.Servers_Administrator_STATUS
	err := copied.AssignProperties_To_Servers_Administrator_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Servers_Administrator_STATUS
	err = actual.AssignProperties_From_Servers_Administrator_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Servers_Administrator_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Servers_Administrator_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServers_Administrator_STATUS, Servers_Administrator_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServers_Administrator_STATUS runs a test to see if a specific instance of Servers_Administrator_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForServers_Administrator_STATUS(subject Servers_Administrator_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Servers_Administrator_STATUS
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

// Generator of Servers_Administrator_STATUS instances for property testing - lazily instantiated by
// Servers_Administrator_STATUSGenerator()
var servers_Administrator_STATUSGenerator gopter.Gen

// Servers_Administrator_STATUSGenerator returns a generator of Servers_Administrator_STATUS instances for property testing.
func Servers_Administrator_STATUSGenerator() gopter.Gen {
	if servers_Administrator_STATUSGenerator != nil {
		return servers_Administrator_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServers_Administrator_STATUS(generators)
	servers_Administrator_STATUSGenerator = gen.Struct(reflect.TypeOf(Servers_Administrator_STATUS{}), generators)

	return servers_Administrator_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForServers_Administrator_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServers_Administrator_STATUS(gens map[string]gopter.Gen) {
	gens["AdministratorType"] = gen.PtrOf(gen.OneConstOf(AdministratorProperties_AdministratorType_STATUS_ActiveDirectory))
	gens["AzureADOnlyAuthentication"] = gen.PtrOf(gen.Bool())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Login"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Sid"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}
