// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210601storage

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

func Test_FlexibleServer_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServer via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServer, FlexibleServerGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServer runs a test to see if a specific instance of FlexibleServer round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServer(subject FlexibleServer) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServer
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

// Generator of FlexibleServer instances for property testing - lazily instantiated by FlexibleServerGenerator()
var flexibleServerGenerator gopter.Gen

// FlexibleServerGenerator returns a generator of FlexibleServer instances for property testing.
func FlexibleServerGenerator() gopter.Gen {
	if flexibleServerGenerator != nil {
		return flexibleServerGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForFlexibleServer(generators)
	flexibleServerGenerator = gen.Struct(reflect.TypeOf(FlexibleServer{}), generators)

	return flexibleServerGenerator
}

// AddRelatedPropertyGeneratorsForFlexibleServer is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFlexibleServer(gens map[string]gopter.Gen) {
	gens["Spec"] = FlexibleServersSpecGenerator()
	gens["Status"] = ServerStatusGenerator()
}

func Test_FlexibleServers_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServers_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServersSpec, FlexibleServersSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServersSpec runs a test to see if a specific instance of FlexibleServers_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServersSpec(subject FlexibleServers_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServers_Spec
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

// Generator of FlexibleServers_Spec instances for property testing - lazily instantiated by
//FlexibleServersSpecGenerator()
var flexibleServersSpecGenerator gopter.Gen

// FlexibleServersSpecGenerator returns a generator of FlexibleServers_Spec instances for property testing.
// We first initialize flexibleServersSpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func FlexibleServersSpecGenerator() gopter.Gen {
	if flexibleServersSpecGenerator != nil {
		return flexibleServersSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServersSpec(generators)
	flexibleServersSpecGenerator = gen.Struct(reflect.TypeOf(FlexibleServers_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServersSpec(generators)
	AddRelatedPropertyGeneratorsForFlexibleServersSpec(generators)
	flexibleServersSpecGenerator = gen.Struct(reflect.TypeOf(FlexibleServers_Spec{}), generators)

	return flexibleServersSpecGenerator
}

// AddIndependentPropertyGeneratorsForFlexibleServersSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFlexibleServersSpec(gens map[string]gopter.Gen) {
	gens["AdministratorLogin"] = gen.PtrOf(gen.AlphaString())
	gens["AvailabilityZone"] = gen.PtrOf(gen.AlphaString())
	gens["AzureName"] = gen.AlphaString()
	gens["CreateMode"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["PointInTimeUTC"] = gen.PtrOf(gen.AlphaString())
	gens["PropertiesTags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Version"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForFlexibleServersSpec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFlexibleServersSpec(gens map[string]gopter.Gen) {
	gens["Backup"] = gen.PtrOf(BackupGenerator())
	gens["HighAvailability"] = gen.PtrOf(HighAvailabilityGenerator())
	gens["MaintenanceWindow"] = gen.PtrOf(MaintenanceWindowGenerator())
	gens["Network"] = gen.PtrOf(NetworkGenerator())
	gens["Sku"] = gen.PtrOf(SkuGenerator())
	gens["Storage"] = gen.PtrOf(StorageGenerator())
}

func Test_Server_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Server_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServerStatus, ServerStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServerStatus runs a test to see if a specific instance of Server_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForServerStatus(subject Server_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Server_Status
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

// Generator of Server_Status instances for property testing - lazily instantiated by ServerStatusGenerator()
var serverStatusGenerator gopter.Gen

// ServerStatusGenerator returns a generator of Server_Status instances for property testing.
// We first initialize serverStatusGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ServerStatusGenerator() gopter.Gen {
	if serverStatusGenerator != nil {
		return serverStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServerStatus(generators)
	serverStatusGenerator = gen.Struct(reflect.TypeOf(Server_Status{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServerStatus(generators)
	AddRelatedPropertyGeneratorsForServerStatus(generators)
	serverStatusGenerator = gen.Struct(reflect.TypeOf(Server_Status{}), generators)

	return serverStatusGenerator
}

// AddIndependentPropertyGeneratorsForServerStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServerStatus(gens map[string]gopter.Gen) {
	gens["AdministratorLogin"] = gen.PtrOf(gen.AlphaString())
	gens["AvailabilityZone"] = gen.PtrOf(gen.AlphaString())
	gens["CreateMode"] = gen.PtrOf(gen.AlphaString())
	gens["FullyQualifiedDomainName"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["MinorVersion"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["PointInTimeUTC"] = gen.PtrOf(gen.AlphaString())
	gens["PropertiesTags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["SourceServerResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["State"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["Version"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForServerStatus is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServerStatus(gens map[string]gopter.Gen) {
	gens["Backup"] = gen.PtrOf(BackupStatusGenerator())
	gens["HighAvailability"] = gen.PtrOf(HighAvailabilityStatusGenerator())
	gens["MaintenanceWindow"] = gen.PtrOf(MaintenanceWindowStatusGenerator())
	gens["Network"] = gen.PtrOf(NetworkStatusGenerator())
	gens["Sku"] = gen.PtrOf(SkuStatusGenerator())
	gens["Storage"] = gen.PtrOf(StorageStatusGenerator())
	gens["SystemData"] = gen.PtrOf(SystemDataStatusGenerator())
}

func Test_Backup_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Backup via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBackup, BackupGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBackup runs a test to see if a specific instance of Backup round trips to JSON and back losslessly
func RunJSONSerializationTestForBackup(subject Backup) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Backup
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

// Generator of Backup instances for property testing - lazily instantiated by BackupGenerator()
var backupGenerator gopter.Gen

// BackupGenerator returns a generator of Backup instances for property testing.
func BackupGenerator() gopter.Gen {
	if backupGenerator != nil {
		return backupGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBackup(generators)
	backupGenerator = gen.Struct(reflect.TypeOf(Backup{}), generators)

	return backupGenerator
}

// AddIndependentPropertyGeneratorsForBackup is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBackup(gens map[string]gopter.Gen) {
	gens["BackupRetentionDays"] = gen.PtrOf(gen.Int())
	gens["GeoRedundantBackup"] = gen.PtrOf(gen.AlphaString())
}

func Test_Backup_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Backup_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBackupStatus, BackupStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBackupStatus runs a test to see if a specific instance of Backup_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForBackupStatus(subject Backup_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Backup_Status
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

// Generator of Backup_Status instances for property testing - lazily instantiated by BackupStatusGenerator()
var backupStatusGenerator gopter.Gen

// BackupStatusGenerator returns a generator of Backup_Status instances for property testing.
func BackupStatusGenerator() gopter.Gen {
	if backupStatusGenerator != nil {
		return backupStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBackupStatus(generators)
	backupStatusGenerator = gen.Struct(reflect.TypeOf(Backup_Status{}), generators)

	return backupStatusGenerator
}

// AddIndependentPropertyGeneratorsForBackupStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBackupStatus(gens map[string]gopter.Gen) {
	gens["BackupRetentionDays"] = gen.PtrOf(gen.Int())
	gens["EarliestRestoreDate"] = gen.PtrOf(gen.AlphaString())
	gens["GeoRedundantBackup"] = gen.PtrOf(gen.AlphaString())
}

func Test_HighAvailability_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of HighAvailability via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForHighAvailability, HighAvailabilityGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForHighAvailability runs a test to see if a specific instance of HighAvailability round trips to JSON and back losslessly
func RunJSONSerializationTestForHighAvailability(subject HighAvailability) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual HighAvailability
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

// Generator of HighAvailability instances for property testing - lazily instantiated by HighAvailabilityGenerator()
var highAvailabilityGenerator gopter.Gen

// HighAvailabilityGenerator returns a generator of HighAvailability instances for property testing.
func HighAvailabilityGenerator() gopter.Gen {
	if highAvailabilityGenerator != nil {
		return highAvailabilityGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForHighAvailability(generators)
	highAvailabilityGenerator = gen.Struct(reflect.TypeOf(HighAvailability{}), generators)

	return highAvailabilityGenerator
}

// AddIndependentPropertyGeneratorsForHighAvailability is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForHighAvailability(gens map[string]gopter.Gen) {
	gens["Mode"] = gen.PtrOf(gen.AlphaString())
	gens["StandbyAvailabilityZone"] = gen.PtrOf(gen.AlphaString())
}

func Test_HighAvailability_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of HighAvailability_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForHighAvailabilityStatus, HighAvailabilityStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForHighAvailabilityStatus runs a test to see if a specific instance of HighAvailability_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForHighAvailabilityStatus(subject HighAvailability_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual HighAvailability_Status
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

// Generator of HighAvailability_Status instances for property testing - lazily instantiated by
//HighAvailabilityStatusGenerator()
var highAvailabilityStatusGenerator gopter.Gen

// HighAvailabilityStatusGenerator returns a generator of HighAvailability_Status instances for property testing.
func HighAvailabilityStatusGenerator() gopter.Gen {
	if highAvailabilityStatusGenerator != nil {
		return highAvailabilityStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForHighAvailabilityStatus(generators)
	highAvailabilityStatusGenerator = gen.Struct(reflect.TypeOf(HighAvailability_Status{}), generators)

	return highAvailabilityStatusGenerator
}

// AddIndependentPropertyGeneratorsForHighAvailabilityStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForHighAvailabilityStatus(gens map[string]gopter.Gen) {
	gens["Mode"] = gen.PtrOf(gen.AlphaString())
	gens["StandbyAvailabilityZone"] = gen.PtrOf(gen.AlphaString())
	gens["State"] = gen.PtrOf(gen.AlphaString())
}

func Test_MaintenanceWindow_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MaintenanceWindow via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMaintenanceWindow, MaintenanceWindowGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMaintenanceWindow runs a test to see if a specific instance of MaintenanceWindow round trips to JSON and back losslessly
func RunJSONSerializationTestForMaintenanceWindow(subject MaintenanceWindow) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MaintenanceWindow
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

// Generator of MaintenanceWindow instances for property testing - lazily instantiated by MaintenanceWindowGenerator()
var maintenanceWindowGenerator gopter.Gen

// MaintenanceWindowGenerator returns a generator of MaintenanceWindow instances for property testing.
func MaintenanceWindowGenerator() gopter.Gen {
	if maintenanceWindowGenerator != nil {
		return maintenanceWindowGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMaintenanceWindow(generators)
	maintenanceWindowGenerator = gen.Struct(reflect.TypeOf(MaintenanceWindow{}), generators)

	return maintenanceWindowGenerator
}

// AddIndependentPropertyGeneratorsForMaintenanceWindow is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMaintenanceWindow(gens map[string]gopter.Gen) {
	gens["CustomWindow"] = gen.PtrOf(gen.AlphaString())
	gens["DayOfWeek"] = gen.PtrOf(gen.Int())
	gens["StartHour"] = gen.PtrOf(gen.Int())
	gens["StartMinute"] = gen.PtrOf(gen.Int())
}

func Test_MaintenanceWindow_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MaintenanceWindow_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMaintenanceWindowStatus, MaintenanceWindowStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMaintenanceWindowStatus runs a test to see if a specific instance of MaintenanceWindow_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForMaintenanceWindowStatus(subject MaintenanceWindow_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MaintenanceWindow_Status
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

// Generator of MaintenanceWindow_Status instances for property testing - lazily instantiated by
//MaintenanceWindowStatusGenerator()
var maintenanceWindowStatusGenerator gopter.Gen

// MaintenanceWindowStatusGenerator returns a generator of MaintenanceWindow_Status instances for property testing.
func MaintenanceWindowStatusGenerator() gopter.Gen {
	if maintenanceWindowStatusGenerator != nil {
		return maintenanceWindowStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMaintenanceWindowStatus(generators)
	maintenanceWindowStatusGenerator = gen.Struct(reflect.TypeOf(MaintenanceWindow_Status{}), generators)

	return maintenanceWindowStatusGenerator
}

// AddIndependentPropertyGeneratorsForMaintenanceWindowStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMaintenanceWindowStatus(gens map[string]gopter.Gen) {
	gens["CustomWindow"] = gen.PtrOf(gen.AlphaString())
	gens["DayOfWeek"] = gen.PtrOf(gen.Int())
	gens["StartHour"] = gen.PtrOf(gen.Int())
	gens["StartMinute"] = gen.PtrOf(gen.Int())
}

func Test_Network_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Network via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetwork, NetworkGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetwork runs a test to see if a specific instance of Network round trips to JSON and back losslessly
func RunJSONSerializationTestForNetwork(subject Network) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Network
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

// Generator of Network instances for property testing - lazily instantiated by NetworkGenerator()
var networkGenerator gopter.Gen

// NetworkGenerator returns a generator of Network instances for property testing.
func NetworkGenerator() gopter.Gen {
	if networkGenerator != nil {
		return networkGenerator
	}

	generators := make(map[string]gopter.Gen)
	networkGenerator = gen.Struct(reflect.TypeOf(Network{}), generators)

	return networkGenerator
}

func Test_Network_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Network_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkStatus, NetworkStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkStatus runs a test to see if a specific instance of Network_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkStatus(subject Network_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Network_Status
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

// Generator of Network_Status instances for property testing - lazily instantiated by NetworkStatusGenerator()
var networkStatusGenerator gopter.Gen

// NetworkStatusGenerator returns a generator of Network_Status instances for property testing.
func NetworkStatusGenerator() gopter.Gen {
	if networkStatusGenerator != nil {
		return networkStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkStatus(generators)
	networkStatusGenerator = gen.Struct(reflect.TypeOf(Network_Status{}), generators)

	return networkStatusGenerator
}

// AddIndependentPropertyGeneratorsForNetworkStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkStatus(gens map[string]gopter.Gen) {
	gens["DelegatedSubnetResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["PrivateDnsZoneArmResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["PublicNetworkAccess"] = gen.PtrOf(gen.AlphaString())
}

func Test_Sku_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Sku via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSku, SkuGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSku runs a test to see if a specific instance of Sku round trips to JSON and back losslessly
func RunJSONSerializationTestForSku(subject Sku) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Sku
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

// Generator of Sku instances for property testing - lazily instantiated by SkuGenerator()
var skuGenerator gopter.Gen

// SkuGenerator returns a generator of Sku instances for property testing.
func SkuGenerator() gopter.Gen {
	if skuGenerator != nil {
		return skuGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSku(generators)
	skuGenerator = gen.Struct(reflect.TypeOf(Sku{}), generators)

	return skuGenerator
}

// AddIndependentPropertyGeneratorsForSku is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSku(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tier"] = gen.PtrOf(gen.AlphaString())
}

func Test_Sku_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Sku_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSkuStatus, SkuStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSkuStatus runs a test to see if a specific instance of Sku_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForSkuStatus(subject Sku_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Sku_Status
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

// Generator of Sku_Status instances for property testing - lazily instantiated by SkuStatusGenerator()
var skuStatusGenerator gopter.Gen

// SkuStatusGenerator returns a generator of Sku_Status instances for property testing.
func SkuStatusGenerator() gopter.Gen {
	if skuStatusGenerator != nil {
		return skuStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSkuStatus(generators)
	skuStatusGenerator = gen.Struct(reflect.TypeOf(Sku_Status{}), generators)

	return skuStatusGenerator
}

// AddIndependentPropertyGeneratorsForSkuStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSkuStatus(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tier"] = gen.PtrOf(gen.AlphaString())
}

func Test_Storage_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Storage via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorage, StorageGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorage runs a test to see if a specific instance of Storage round trips to JSON and back losslessly
func RunJSONSerializationTestForStorage(subject Storage) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Storage
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

// Generator of Storage instances for property testing - lazily instantiated by StorageGenerator()
var storageGenerator gopter.Gen

// StorageGenerator returns a generator of Storage instances for property testing.
func StorageGenerator() gopter.Gen {
	if storageGenerator != nil {
		return storageGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorage(generators)
	storageGenerator = gen.Struct(reflect.TypeOf(Storage{}), generators)

	return storageGenerator
}

// AddIndependentPropertyGeneratorsForStorage is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorage(gens map[string]gopter.Gen) {
	gens["StorageSizeGB"] = gen.PtrOf(gen.Int())
}

func Test_Storage_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Storage_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageStatus, StorageStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageStatus runs a test to see if a specific instance of Storage_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageStatus(subject Storage_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Storage_Status
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

// Generator of Storage_Status instances for property testing - lazily instantiated by StorageStatusGenerator()
var storageStatusGenerator gopter.Gen

// StorageStatusGenerator returns a generator of Storage_Status instances for property testing.
func StorageStatusGenerator() gopter.Gen {
	if storageStatusGenerator != nil {
		return storageStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageStatus(generators)
	storageStatusGenerator = gen.Struct(reflect.TypeOf(Storage_Status{}), generators)

	return storageStatusGenerator
}

// AddIndependentPropertyGeneratorsForStorageStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageStatus(gens map[string]gopter.Gen) {
	gens["StorageSizeGB"] = gen.PtrOf(gen.Int())
}

func Test_SystemData_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SystemData_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSystemDataStatus, SystemDataStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSystemDataStatus runs a test to see if a specific instance of SystemData_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForSystemDataStatus(subject SystemData_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SystemData_Status
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

// Generator of SystemData_Status instances for property testing - lazily instantiated by SystemDataStatusGenerator()
var systemDataStatusGenerator gopter.Gen

// SystemDataStatusGenerator returns a generator of SystemData_Status instances for property testing.
func SystemDataStatusGenerator() gopter.Gen {
	if systemDataStatusGenerator != nil {
		return systemDataStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSystemDataStatus(generators)
	systemDataStatusGenerator = gen.Struct(reflect.TypeOf(SystemData_Status{}), generators)

	return systemDataStatusGenerator
}

// AddIndependentPropertyGeneratorsForSystemDataStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSystemDataStatus(gens map[string]gopter.Gen) {
	gens["CreatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedBy"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedByType"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedAt"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedBy"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedByType"] = gen.PtrOf(gen.AlphaString())
}