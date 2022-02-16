// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515

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

func Test_DatabaseAccountsMongodbDatabasesCollection_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccountsMongodbDatabasesCollection_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccountsMongodbDatabasesCollection_SpecARM, DatabaseAccountsMongodbDatabasesCollection_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccountsMongodbDatabasesCollection_SpecARM runs a test to see if a specific instance of DatabaseAccountsMongodbDatabasesCollection_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccountsMongodbDatabasesCollection_SpecARM(subject DatabaseAccountsMongodbDatabasesCollection_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccountsMongodbDatabasesCollection_SpecARM
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

// Generator of DatabaseAccountsMongodbDatabasesCollection_SpecARM instances for property testing - lazily instantiated
//by DatabaseAccountsMongodbDatabasesCollection_SpecARMGenerator()
var databaseAccountsMongodbDatabasesCollection_specARMGenerator gopter.Gen

// DatabaseAccountsMongodbDatabasesCollection_SpecARMGenerator returns a generator of DatabaseAccountsMongodbDatabasesCollection_SpecARM instances for property testing.
// We first initialize databaseAccountsMongodbDatabasesCollection_specARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccountsMongodbDatabasesCollection_SpecARMGenerator() gopter.Gen {
	if databaseAccountsMongodbDatabasesCollection_specARMGenerator != nil {
		return databaseAccountsMongodbDatabasesCollection_specARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollection_SpecARM(generators)
	databaseAccountsMongodbDatabasesCollection_specARMGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsMongodbDatabasesCollection_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollection_SpecARM(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollection_SpecARM(generators)
	databaseAccountsMongodbDatabasesCollection_specARMGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsMongodbDatabasesCollection_SpecARM{}), generators)

	return databaseAccountsMongodbDatabasesCollection_specARMGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollection_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollection_SpecARM(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollection_SpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollection_SpecARM(gens map[string]gopter.Gen) {
	gens["Properties"] = MongoDBCollectionPropertiesARMGenerator()
}

func Test_MongoDBCollectionPropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoDBCollectionPropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoDBCollectionPropertiesARM, MongoDBCollectionPropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoDBCollectionPropertiesARM runs a test to see if a specific instance of MongoDBCollectionPropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoDBCollectionPropertiesARM(subject MongoDBCollectionPropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoDBCollectionPropertiesARM
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

// Generator of MongoDBCollectionPropertiesARM instances for property testing - lazily instantiated by
//MongoDBCollectionPropertiesARMGenerator()
var mongoDBCollectionPropertiesARMGenerator gopter.Gen

// MongoDBCollectionPropertiesARMGenerator returns a generator of MongoDBCollectionPropertiesARM instances for property testing.
func MongoDBCollectionPropertiesARMGenerator() gopter.Gen {
	if mongoDBCollectionPropertiesARMGenerator != nil {
		return mongoDBCollectionPropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForMongoDBCollectionPropertiesARM(generators)
	mongoDBCollectionPropertiesARMGenerator = gen.Struct(reflect.TypeOf(MongoDBCollectionPropertiesARM{}), generators)

	return mongoDBCollectionPropertiesARMGenerator
}

// AddRelatedPropertyGeneratorsForMongoDBCollectionPropertiesARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongoDBCollectionPropertiesARM(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(CreateUpdateOptionsARMGenerator())
	gens["Resource"] = MongoDBCollectionResourceARMGenerator()
}

func Test_MongoDBCollectionResourceARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoDBCollectionResourceARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoDBCollectionResourceARM, MongoDBCollectionResourceARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoDBCollectionResourceARM runs a test to see if a specific instance of MongoDBCollectionResourceARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoDBCollectionResourceARM(subject MongoDBCollectionResourceARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoDBCollectionResourceARM
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

// Generator of MongoDBCollectionResourceARM instances for property testing - lazily instantiated by
//MongoDBCollectionResourceARMGenerator()
var mongoDBCollectionResourceARMGenerator gopter.Gen

// MongoDBCollectionResourceARMGenerator returns a generator of MongoDBCollectionResourceARM instances for property testing.
// We first initialize mongoDBCollectionResourceARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func MongoDBCollectionResourceARMGenerator() gopter.Gen {
	if mongoDBCollectionResourceARMGenerator != nil {
		return mongoDBCollectionResourceARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoDBCollectionResourceARM(generators)
	mongoDBCollectionResourceARMGenerator = gen.Struct(reflect.TypeOf(MongoDBCollectionResourceARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoDBCollectionResourceARM(generators)
	AddRelatedPropertyGeneratorsForMongoDBCollectionResourceARM(generators)
	mongoDBCollectionResourceARMGenerator = gen.Struct(reflect.TypeOf(MongoDBCollectionResourceARM{}), generators)

	return mongoDBCollectionResourceARMGenerator
}

// AddIndependentPropertyGeneratorsForMongoDBCollectionResourceARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongoDBCollectionResourceARM(gens map[string]gopter.Gen) {
	gens["AnalyticalStorageTtl"] = gen.PtrOf(gen.Int())
	gens["Id"] = gen.AlphaString()
	gens["ShardKey"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForMongoDBCollectionResourceARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongoDBCollectionResourceARM(gens map[string]gopter.Gen) {
	gens["Indexes"] = gen.SliceOf(MongoIndexARMGenerator())
}

func Test_MongoIndexARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoIndexARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoIndexARM, MongoIndexARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoIndexARM runs a test to see if a specific instance of MongoIndexARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoIndexARM(subject MongoIndexARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoIndexARM
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

// Generator of MongoIndexARM instances for property testing - lazily instantiated by MongoIndexARMGenerator()
var mongoIndexARMGenerator gopter.Gen

// MongoIndexARMGenerator returns a generator of MongoIndexARM instances for property testing.
func MongoIndexARMGenerator() gopter.Gen {
	if mongoIndexARMGenerator != nil {
		return mongoIndexARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForMongoIndexARM(generators)
	mongoIndexARMGenerator = gen.Struct(reflect.TypeOf(MongoIndexARM{}), generators)

	return mongoIndexARMGenerator
}

// AddRelatedPropertyGeneratorsForMongoIndexARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongoIndexARM(gens map[string]gopter.Gen) {
	gens["Key"] = gen.PtrOf(MongoIndexKeysARMGenerator())
	gens["Options"] = gen.PtrOf(MongoIndexOptionsARMGenerator())
}

func Test_MongoIndexKeysARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoIndexKeysARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoIndexKeysARM, MongoIndexKeysARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoIndexKeysARM runs a test to see if a specific instance of MongoIndexKeysARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoIndexKeysARM(subject MongoIndexKeysARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoIndexKeysARM
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

// Generator of MongoIndexKeysARM instances for property testing - lazily instantiated by MongoIndexKeysARMGenerator()
var mongoIndexKeysARMGenerator gopter.Gen

// MongoIndexKeysARMGenerator returns a generator of MongoIndexKeysARM instances for property testing.
func MongoIndexKeysARMGenerator() gopter.Gen {
	if mongoIndexKeysARMGenerator != nil {
		return mongoIndexKeysARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoIndexKeysARM(generators)
	mongoIndexKeysARMGenerator = gen.Struct(reflect.TypeOf(MongoIndexKeysARM{}), generators)

	return mongoIndexKeysARMGenerator
}

// AddIndependentPropertyGeneratorsForMongoIndexKeysARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongoIndexKeysARM(gens map[string]gopter.Gen) {
	gens["Keys"] = gen.SliceOf(gen.AlphaString())
}

func Test_MongoIndexOptionsARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoIndexOptionsARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoIndexOptionsARM, MongoIndexOptionsARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoIndexOptionsARM runs a test to see if a specific instance of MongoIndexOptionsARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoIndexOptionsARM(subject MongoIndexOptionsARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoIndexOptionsARM
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

// Generator of MongoIndexOptionsARM instances for property testing - lazily instantiated by
//MongoIndexOptionsARMGenerator()
var mongoIndexOptionsARMGenerator gopter.Gen

// MongoIndexOptionsARMGenerator returns a generator of MongoIndexOptionsARM instances for property testing.
func MongoIndexOptionsARMGenerator() gopter.Gen {
	if mongoIndexOptionsARMGenerator != nil {
		return mongoIndexOptionsARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoIndexOptionsARM(generators)
	mongoIndexOptionsARMGenerator = gen.Struct(reflect.TypeOf(MongoIndexOptionsARM{}), generators)

	return mongoIndexOptionsARMGenerator
}

// AddIndependentPropertyGeneratorsForMongoIndexOptionsARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongoIndexOptionsARM(gens map[string]gopter.Gen) {
	gens["ExpireAfterSeconds"] = gen.PtrOf(gen.Int())
	gens["Unique"] = gen.PtrOf(gen.Bool())
}