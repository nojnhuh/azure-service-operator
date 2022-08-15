// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210601

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Profiles_SpecARM struct {
	// Location: Location to deploy resource to
	Location *string `json:"location,omitempty"`

	// Name: Name of the Azure Front Door Standard or Azure Front Door Premium or CDN profile which is unique within the
	// resource group.
	Name string `json:"name,omitempty"`

	// Properties: The JSON object that contains the properties required to create a profile.
	Properties *ProfilePropertiesARM `json:"properties,omitempty"`

	// Sku: Standard_Verizon = The SKU name for a Standard Verizon CDN profile.
	// Premium_Verizon = The SKU name for a Premium Verizon CDN profile.
	// Custom_Verizon = The SKU name for a Custom Verizon CDN profile.
	// Standard_Akamai = The SKU name for an Akamai CDN profile.
	// Standard_ChinaCdn = The SKU name for a China CDN profile for VOD, Web and download scenarios using GB based billing
	// model.
	// Standard_Microsoft = The SKU name for a Standard Microsoft CDN profile.
	// Standard_AzureFrontDoor =  The SKU name for an Azure Front Door Standard profile.
	// Premium_AzureFrontDoor = The SKU name for an Azure Front Door Premium profile.
	// Standard_955BandWidth_ChinaCdn = The SKU name for a China CDN profile for VOD, Web and download scenarios using 95-5
	// peak bandwidth billing model.
	// Standard_AvgBandWidth_ChinaCdn = The SKU name for a China CDN profile for VOD, Web and download scenarios using monthly
	// average peak bandwidth billing model.
	// StandardPlus_ChinaCdn = The SKU name for a China CDN profile for live-streaming using GB based billing model.
	// StandardPlus_955BandWidth_ChinaCdn = The SKU name for a China CDN live-streaming profile using 95-5 peak bandwidth
	// billing model.
	// StandardPlus_AvgBandWidth_ChinaCdn = The SKU name for a China CDN live-streaming profile using monthly average peak
	// bandwidth billing model.
	Sku *SkuARM `json:"sku,omitempty"`

	// Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Profiles_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-06-01"
func (profiles Profiles_SpecARM) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetName returns the Name of the resource
func (profiles *Profiles_SpecARM) GetName() string {
	return profiles.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cdn/profiles"
func (profiles *Profiles_SpecARM) GetType() string {
	return "Microsoft.Cdn/profiles"
}

// Generated from: https://schema.management.azure.com/schemas/2021-06-01/Microsoft.Cdn.json#/definitions/ProfileProperties
type ProfilePropertiesARM struct {
	// OriginResponseTimeoutSeconds: Send and receive timeout on forwarding request to the origin. When timeout is reached, the
	// request fails and returns.
	OriginResponseTimeoutSeconds *int `json:"originResponseTimeoutSeconds,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2021-06-01/Microsoft.Cdn.json#/definitions/Sku
type SkuARM struct {
	// Name: Name of the pricing tier.
	Name *SkuName `json:"name,omitempty"`
}

// +kubebuilder:validation:Enum={"Custom_Verizon","Premium_AzureFrontDoor","Premium_Verizon","Standard_955BandWidth_ChinaCdn","Standard_Akamai","Standard_AvgBandWidth_ChinaCdn","Standard_AzureFrontDoor","Standard_ChinaCdn","Standard_Microsoft","StandardPlus_955BandWidth_ChinaCdn","StandardPlus_AvgBandWidth_ChinaCdn","StandardPlus_ChinaCdn","Standard_Verizon"}
type SkuName string

const (
	SkuNameCustomVerizon                    = SkuName("Custom_Verizon")
	SkuNamePremiumAzureFrontDoor            = SkuName("Premium_AzureFrontDoor")
	SkuNamePremiumVerizon                   = SkuName("Premium_Verizon")
	SkuNameStandard955BandWidthChinaCdn     = SkuName("Standard_955BandWidth_ChinaCdn")
	SkuNameStandardAkamai                   = SkuName("Standard_Akamai")
	SkuNameStandardAvgBandWidthChinaCdn     = SkuName("Standard_AvgBandWidth_ChinaCdn")
	SkuNameStandardAzureFrontDoor           = SkuName("Standard_AzureFrontDoor")
	SkuNameStandardChinaCdn                 = SkuName("Standard_ChinaCdn")
	SkuNameStandardMicrosoft                = SkuName("Standard_Microsoft")
	SkuNameStandardPlus955BandWidthChinaCdn = SkuName("StandardPlus_955BandWidth_ChinaCdn")
	SkuNameStandardPlusAvgBandWidthChinaCdn = SkuName("StandardPlus_AvgBandWidth_ChinaCdn")
	SkuNameStandardPlusChinaCdn             = SkuName("StandardPlus_ChinaCdn")
	SkuNameStandardVerizon                  = SkuName("Standard_Verizon")
)