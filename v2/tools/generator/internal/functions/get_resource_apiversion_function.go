/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"fmt"
	"strings"

	"github.com/dave/dst"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// NewGetAPIVersionFunction returns a function that returns a static API version string
func NewGetAPIVersionFunction(
	apiVersionTypeName astmodel.TypeName,
	apiVersionEnumValue astmodel.EnumValue,
	idFactory astmodel.IdentifierFactory) astmodel.Function {

	value := strings.Trim(apiVersionEnumValue.Value, "\"")
	comment := fmt.Sprintf("returns the ARM API version of the resource. This is always %q", value)
	result := NewObjectFunction("Get"+astmodel.APIVersionProperty, idFactory,
		createBodyReturningValue(
			astbuilder.CallFunc("string", dst.NewIdent(apiVersionTypeName.Name()+apiVersionEnumValue.Identifier)),
			astmodel.StringType,
			comment,
			ReceiverTypeStruct))

	result.AddPackageReference(astmodel.GenRuntimeReference)
	result.AddReferencedTypes(apiVersionTypeName)

	return result
}
