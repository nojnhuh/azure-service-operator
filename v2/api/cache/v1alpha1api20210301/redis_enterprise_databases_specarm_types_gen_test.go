// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210301

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

func Test_RedisEnterpriseDatabases_SPECARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisEnterpriseDatabases_SPECARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisEnterpriseDatabases_SPECARM, RedisEnterpriseDatabases_SPECARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisEnterpriseDatabases_SPECARM runs a test to see if a specific instance of RedisEnterpriseDatabases_SPECARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisEnterpriseDatabases_SPECARM(subject RedisEnterpriseDatabases_SPECARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisEnterpriseDatabases_SPECARM
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

// Generator of RedisEnterpriseDatabases_SPECARM instances for property testing - lazily instantiated by
//RedisEnterpriseDatabases_SPECARMGenerator()
var redisEnterpriseDatabases_specarmGenerator gopter.Gen

// RedisEnterpriseDatabases_SPECARMGenerator returns a generator of RedisEnterpriseDatabases_SPECARM instances for property testing.
// We first initialize redisEnterpriseDatabases_specarmGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RedisEnterpriseDatabases_SPECARMGenerator() gopter.Gen {
	if redisEnterpriseDatabases_specarmGenerator != nil {
		return redisEnterpriseDatabases_specarmGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisEnterpriseDatabases_SPECARM(generators)
	redisEnterpriseDatabases_specarmGenerator = gen.Struct(reflect.TypeOf(RedisEnterpriseDatabases_SPECARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisEnterpriseDatabases_SPECARM(generators)
	AddRelatedPropertyGeneratorsForRedisEnterpriseDatabases_SPECARM(generators)
	redisEnterpriseDatabases_specarmGenerator = gen.Struct(reflect.TypeOf(RedisEnterpriseDatabases_SPECARM{}), generators)

	return redisEnterpriseDatabases_specarmGenerator
}

// AddIndependentPropertyGeneratorsForRedisEnterpriseDatabases_SPECARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisEnterpriseDatabases_SPECARM(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForRedisEnterpriseDatabases_SPECARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisEnterpriseDatabases_SPECARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(DatabaseProperties_SpecARMGenerator())
}

func Test_DatabaseProperties_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseProperties_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseProperties_SpecARM, DatabaseProperties_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseProperties_SpecARM runs a test to see if a specific instance of DatabaseProperties_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseProperties_SpecARM(subject DatabaseProperties_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseProperties_SpecARM
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

// Generator of DatabaseProperties_SpecARM instances for property testing - lazily instantiated by
//DatabaseProperties_SpecARMGenerator()
var databaseProperties_specARMGenerator gopter.Gen

// DatabaseProperties_SpecARMGenerator returns a generator of DatabaseProperties_SpecARM instances for property testing.
// We first initialize databaseProperties_specARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseProperties_SpecARMGenerator() gopter.Gen {
	if databaseProperties_specARMGenerator != nil {
		return databaseProperties_specARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseProperties_SpecARM(generators)
	databaseProperties_specARMGenerator = gen.Struct(reflect.TypeOf(DatabaseProperties_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseProperties_SpecARM(generators)
	AddRelatedPropertyGeneratorsForDatabaseProperties_SpecARM(generators)
	databaseProperties_specARMGenerator = gen.Struct(reflect.TypeOf(DatabaseProperties_SpecARM{}), generators)

	return databaseProperties_specARMGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseProperties_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseProperties_SpecARM(gens map[string]gopter.Gen) {
	gens["ClientProtocol"] = gen.PtrOf(gen.OneConstOf(DatabaseProperties_ClientProtocol_SpecEncrypted, DatabaseProperties_ClientProtocol_SpecPlaintext))
	gens["ClusteringPolicy"] = gen.PtrOf(gen.OneConstOf(DatabaseProperties_ClusteringPolicy_SpecEnterpriseCluster, DatabaseProperties_ClusteringPolicy_SpecOSSCluster))
	gens["EvictionPolicy"] = gen.PtrOf(gen.OneConstOf(
		DatabaseProperties_EvictionPolicy_SpecAllKeysLFU,
		DatabaseProperties_EvictionPolicy_SpecAllKeysLRU,
		DatabaseProperties_EvictionPolicy_SpecAllKeysRandom,
		DatabaseProperties_EvictionPolicy_SpecNoEviction,
		DatabaseProperties_EvictionPolicy_SpecVolatileLFU,
		DatabaseProperties_EvictionPolicy_SpecVolatileLRU,
		DatabaseProperties_EvictionPolicy_SpecVolatileRandom,
		DatabaseProperties_EvictionPolicy_SpecVolatileTTL))
	gens["Port"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForDatabaseProperties_SpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseProperties_SpecARM(gens map[string]gopter.Gen) {
	gens["Modules"] = gen.SliceOf(Module_SpecARMGenerator())
	gens["Persistence"] = gen.PtrOf(Persistence_SpecARMGenerator())
}

func Test_Module_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Module_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForModule_SpecARM, Module_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForModule_SpecARM runs a test to see if a specific instance of Module_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForModule_SpecARM(subject Module_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Module_SpecARM
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

// Generator of Module_SpecARM instances for property testing - lazily instantiated by Module_SpecARMGenerator()
var module_specARMGenerator gopter.Gen

// Module_SpecARMGenerator returns a generator of Module_SpecARM instances for property testing.
func Module_SpecARMGenerator() gopter.Gen {
	if module_specARMGenerator != nil {
		return module_specARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForModule_SpecARM(generators)
	module_specARMGenerator = gen.Struct(reflect.TypeOf(Module_SpecARM{}), generators)

	return module_specARMGenerator
}

// AddIndependentPropertyGeneratorsForModule_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForModule_SpecARM(gens map[string]gopter.Gen) {
	gens["Args"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
}

func Test_Persistence_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Persistence_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPersistence_SpecARM, Persistence_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPersistence_SpecARM runs a test to see if a specific instance of Persistence_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPersistence_SpecARM(subject Persistence_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Persistence_SpecARM
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

// Generator of Persistence_SpecARM instances for property testing - lazily instantiated by
//Persistence_SpecARMGenerator()
var persistence_specARMGenerator gopter.Gen

// Persistence_SpecARMGenerator returns a generator of Persistence_SpecARM instances for property testing.
func Persistence_SpecARMGenerator() gopter.Gen {
	if persistence_specARMGenerator != nil {
		return persistence_specARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPersistence_SpecARM(generators)
	persistence_specARMGenerator = gen.Struct(reflect.TypeOf(Persistence_SpecARM{}), generators)

	return persistence_specARMGenerator
}

// AddIndependentPropertyGeneratorsForPersistence_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPersistence_SpecARM(gens map[string]gopter.Gen) {
	gens["AofEnabled"] = gen.PtrOf(gen.Bool())
	gens["AofFrequency"] = gen.PtrOf(gen.OneConstOf(Persistence_AofFrequency_Spec1S, Persistence_AofFrequency_SpecAlways))
	gens["RdbEnabled"] = gen.PtrOf(gen.Bool())
	gens["RdbFrequency"] = gen.PtrOf(gen.OneConstOf(Persistence_RdbFrequency_Spec12H, Persistence_RdbFrequency_Spec1H, Persistence_RdbFrequency_Spec6H))
}