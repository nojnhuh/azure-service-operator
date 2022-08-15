// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210601

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

func Test_FlexibleServers_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServers_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServersSpecARM, FlexibleServersSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServersSpecARM runs a test to see if a specific instance of FlexibleServers_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServersSpecARM(subject FlexibleServers_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServers_SpecARM
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

// Generator of FlexibleServers_SpecARM instances for property testing - lazily instantiated by
// FlexibleServersSpecARMGenerator()
var flexibleServersSpecARMGenerator gopter.Gen

// FlexibleServersSpecARMGenerator returns a generator of FlexibleServers_SpecARM instances for property testing.
// We first initialize flexibleServersSpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func FlexibleServersSpecARMGenerator() gopter.Gen {
	if flexibleServersSpecARMGenerator != nil {
		return flexibleServersSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServersSpecARM(generators)
	flexibleServersSpecARMGenerator = gen.Struct(reflect.TypeOf(FlexibleServers_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServersSpecARM(generators)
	AddRelatedPropertyGeneratorsForFlexibleServersSpecARM(generators)
	flexibleServersSpecARMGenerator = gen.Struct(reflect.TypeOf(FlexibleServers_SpecARM{}), generators)

	return flexibleServersSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForFlexibleServersSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFlexibleServersSpecARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForFlexibleServersSpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFlexibleServersSpecARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ServerPropertiesARMGenerator())
	gens["Sku"] = gen.PtrOf(SkuARMGenerator())
}

func Test_ServerPropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServerPropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServerPropertiesARM, ServerPropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServerPropertiesARM runs a test to see if a specific instance of ServerPropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForServerPropertiesARM(subject ServerPropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServerPropertiesARM
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

// Generator of ServerPropertiesARM instances for property testing - lazily instantiated by
// ServerPropertiesARMGenerator()
var serverPropertiesARMGenerator gopter.Gen

// ServerPropertiesARMGenerator returns a generator of ServerPropertiesARM instances for property testing.
// We first initialize serverPropertiesARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ServerPropertiesARMGenerator() gopter.Gen {
	if serverPropertiesARMGenerator != nil {
		return serverPropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServerPropertiesARM(generators)
	serverPropertiesARMGenerator = gen.Struct(reflect.TypeOf(ServerPropertiesARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServerPropertiesARM(generators)
	AddRelatedPropertyGeneratorsForServerPropertiesARM(generators)
	serverPropertiesARMGenerator = gen.Struct(reflect.TypeOf(ServerPropertiesARM{}), generators)

	return serverPropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForServerPropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServerPropertiesARM(gens map[string]gopter.Gen) {
	gens["AdministratorLogin"] = gen.PtrOf(gen.AlphaString())
	gens["AdministratorLoginPassword"] = gen.PtrOf(gen.AlphaString())
	gens["AvailabilityZone"] = gen.PtrOf(gen.AlphaString())
	gens["CreateMode"] = gen.PtrOf(gen.OneConstOf(
		ServerPropertiesCreateModeCreate,
		ServerPropertiesCreateModeDefault,
		ServerPropertiesCreateModePointInTimeRestore,
		ServerPropertiesCreateModeUpdate))
	gens["PointInTimeUTC"] = gen.PtrOf(gen.AlphaString())
	gens["SourceServerResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["Version"] = gen.PtrOf(gen.OneConstOf(ServerPropertiesVersion11, ServerPropertiesVersion12, ServerPropertiesVersion13))
}

// AddRelatedPropertyGeneratorsForServerPropertiesARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServerPropertiesARM(gens map[string]gopter.Gen) {
	gens["Backup"] = gen.PtrOf(BackupARMGenerator())
	gens["HighAvailability"] = gen.PtrOf(HighAvailabilityARMGenerator())
	gens["MaintenanceWindow"] = gen.PtrOf(MaintenanceWindowARMGenerator())
	gens["Network"] = gen.PtrOf(NetworkARMGenerator())
	gens["Storage"] = gen.PtrOf(StorageARMGenerator())
}

func Test_SkuARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SkuARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSkuARM, SkuARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSkuARM runs a test to see if a specific instance of SkuARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSkuARM(subject SkuARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SkuARM
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

// Generator of SkuARM instances for property testing - lazily instantiated by SkuARMGenerator()
var skuARMGenerator gopter.Gen

// SkuARMGenerator returns a generator of SkuARM instances for property testing.
func SkuARMGenerator() gopter.Gen {
	if skuARMGenerator != nil {
		return skuARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSkuARM(generators)
	skuARMGenerator = gen.Struct(reflect.TypeOf(SkuARM{}), generators)

	return skuARMGenerator
}

// AddIndependentPropertyGeneratorsForSkuARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSkuARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tier"] = gen.PtrOf(gen.OneConstOf(SkuTierBurstable, SkuTierGeneralPurpose, SkuTierMemoryOptimized))
}

func Test_BackupARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BackupARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBackupARM, BackupARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBackupARM runs a test to see if a specific instance of BackupARM round trips to JSON and back losslessly
func RunJSONSerializationTestForBackupARM(subject BackupARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BackupARM
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

// Generator of BackupARM instances for property testing - lazily instantiated by BackupARMGenerator()
var backupARMGenerator gopter.Gen

// BackupARMGenerator returns a generator of BackupARM instances for property testing.
func BackupARMGenerator() gopter.Gen {
	if backupARMGenerator != nil {
		return backupARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBackupARM(generators)
	backupARMGenerator = gen.Struct(reflect.TypeOf(BackupARM{}), generators)

	return backupARMGenerator
}

// AddIndependentPropertyGeneratorsForBackupARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBackupARM(gens map[string]gopter.Gen) {
	gens["BackupRetentionDays"] = gen.PtrOf(gen.Int())
	gens["GeoRedundantBackup"] = gen.PtrOf(gen.OneConstOf(BackupGeoRedundantBackupDisabled, BackupGeoRedundantBackupEnabled))
}

func Test_HighAvailabilityARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of HighAvailabilityARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForHighAvailabilityARM, HighAvailabilityARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForHighAvailabilityARM runs a test to see if a specific instance of HighAvailabilityARM round trips to JSON and back losslessly
func RunJSONSerializationTestForHighAvailabilityARM(subject HighAvailabilityARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual HighAvailabilityARM
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

// Generator of HighAvailabilityARM instances for property testing - lazily instantiated by
// HighAvailabilityARMGenerator()
var highAvailabilityARMGenerator gopter.Gen

// HighAvailabilityARMGenerator returns a generator of HighAvailabilityARM instances for property testing.
func HighAvailabilityARMGenerator() gopter.Gen {
	if highAvailabilityARMGenerator != nil {
		return highAvailabilityARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForHighAvailabilityARM(generators)
	highAvailabilityARMGenerator = gen.Struct(reflect.TypeOf(HighAvailabilityARM{}), generators)

	return highAvailabilityARMGenerator
}

// AddIndependentPropertyGeneratorsForHighAvailabilityARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForHighAvailabilityARM(gens map[string]gopter.Gen) {
	gens["Mode"] = gen.PtrOf(gen.OneConstOf(HighAvailabilityModeDisabled, HighAvailabilityModeZoneRedundant))
	gens["StandbyAvailabilityZone"] = gen.PtrOf(gen.AlphaString())
}

func Test_MaintenanceWindowARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MaintenanceWindowARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMaintenanceWindowARM, MaintenanceWindowARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMaintenanceWindowARM runs a test to see if a specific instance of MaintenanceWindowARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMaintenanceWindowARM(subject MaintenanceWindowARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MaintenanceWindowARM
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

// Generator of MaintenanceWindowARM instances for property testing - lazily instantiated by
// MaintenanceWindowARMGenerator()
var maintenanceWindowARMGenerator gopter.Gen

// MaintenanceWindowARMGenerator returns a generator of MaintenanceWindowARM instances for property testing.
func MaintenanceWindowARMGenerator() gopter.Gen {
	if maintenanceWindowARMGenerator != nil {
		return maintenanceWindowARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMaintenanceWindowARM(generators)
	maintenanceWindowARMGenerator = gen.Struct(reflect.TypeOf(MaintenanceWindowARM{}), generators)

	return maintenanceWindowARMGenerator
}

// AddIndependentPropertyGeneratorsForMaintenanceWindowARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMaintenanceWindowARM(gens map[string]gopter.Gen) {
	gens["CustomWindow"] = gen.PtrOf(gen.AlphaString())
	gens["DayOfWeek"] = gen.PtrOf(gen.Int())
	gens["StartHour"] = gen.PtrOf(gen.Int())
	gens["StartMinute"] = gen.PtrOf(gen.Int())
}

func Test_NetworkARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkARM, NetworkARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkARM runs a test to see if a specific instance of NetworkARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkARM(subject NetworkARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkARM
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

// Generator of NetworkARM instances for property testing - lazily instantiated by NetworkARMGenerator()
var networkARMGenerator gopter.Gen

// NetworkARMGenerator returns a generator of NetworkARM instances for property testing.
func NetworkARMGenerator() gopter.Gen {
	if networkARMGenerator != nil {
		return networkARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkARM(generators)
	networkARMGenerator = gen.Struct(reflect.TypeOf(NetworkARM{}), generators)

	return networkARMGenerator
}

// AddIndependentPropertyGeneratorsForNetworkARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkARM(gens map[string]gopter.Gen) {
	gens["DelegatedSubnetResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["PrivateDnsZoneArmResourceId"] = gen.PtrOf(gen.AlphaString())
}

func Test_StorageARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageARM, StorageARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageARM runs a test to see if a specific instance of StorageARM round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageARM(subject StorageARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageARM
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

// Generator of StorageARM instances for property testing - lazily instantiated by StorageARMGenerator()
var storageARMGenerator gopter.Gen

// StorageARMGenerator returns a generator of StorageARM instances for property testing.
func StorageARMGenerator() gopter.Gen {
	if storageARMGenerator != nil {
		return storageARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageARM(generators)
	storageARMGenerator = gen.Struct(reflect.TypeOf(StorageARM{}), generators)

	return storageARMGenerator
}

// AddIndependentPropertyGeneratorsForStorageARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageARM(gens map[string]gopter.Gen) {
	gens["StorageSizeGB"] = gen.PtrOf(gen.Int())
}