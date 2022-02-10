// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type VirtualNetworks_SPECARM struct {
	AzureName string `json:"azureName"`

	//ExtendedLocation: The extended location of the virtual network.
	ExtendedLocation *ExtendedLocation_SpecARM `json:"extendedLocation,omitempty"`
	Id               *string                   `json:"id,omitempty"`

	//Location: Resource location.
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name"`

	//Properties: Properties of the virtual network.
	Properties *VirtualNetworkPropertiesFormat_SpecARM `json:"properties,omitempty"`

	//Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &VirtualNetworks_SPECARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-11-01"
func (specarm VirtualNetworks_SPECARM) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetName returns the Name of the resource
func (specarm VirtualNetworks_SPECARM) GetName() string {
	return specarm.Name
}

// GetType returns the ARM Type of the resource. This is always ""
func (specarm VirtualNetworks_SPECARM) GetType() string {
	return ""
}

type VirtualNetworkPropertiesFormat_SpecARM struct {
	//AddressSpace: The AddressSpace that contains an array of IP address ranges that
	//can be used by subnets.
	AddressSpace *AddressSpace_SpecARM `json:"addressSpace,omitempty"`

	//BgpCommunities: Bgp Communities sent over ExpressRoute with each route
	//corresponding to a prefix in this VNET.
	BgpCommunities *VirtualNetworkBgpCommunities_SpecARM `json:"bgpCommunities,omitempty"`

	//DdosProtectionPlan: The DDoS protection plan associated with the virtual network.
	DdosProtectionPlan *SubResource_SpecARM `json:"ddosProtectionPlan,omitempty"`

	//DhcpOptions: The dhcpOptions that contains an array of DNS servers available to
	//VMs deployed in the virtual network.
	DhcpOptions *DhcpOptions_SpecARM `json:"dhcpOptions,omitempty"`

	//EnableDdosProtection: Indicates if DDoS protection is enabled for all the
	//protected resources in the virtual network. It requires a DDoS protection plan
	//associated with the resource.
	EnableDdosProtection *bool `json:"enableDdosProtection,omitempty"`

	//EnableVmProtection: Indicates if VM protection is enabled for all the subnets in
	//the virtual network.
	EnableVmProtection *bool `json:"enableVmProtection,omitempty"`

	//IpAllocations: Array of IpAllocation which reference this VNET.
	IpAllocations []SubResource_SpecARM `json:"ipAllocations,omitempty"`

	//Subnets: A list of subnets in a Virtual Network.
	Subnets []Subnet_Spec_VirtualNetwork_SubResourceEmbeddedARM `json:"subnets,omitempty"`

	//VirtualNetworkPeerings: A list of peerings in a Virtual Network.
	VirtualNetworkPeerings []VirtualNetworkPeering_SpecARM `json:"virtualNetworkPeerings,omitempty"`
}

type DhcpOptions_SpecARM struct {
	//DnsServers: The list of DNS servers IP addresses.
	DnsServers []string `json:"dnsServers,omitempty"`
}

type Subnet_Spec_VirtualNetwork_SubResourceEmbeddedARM struct {
	Id *string `json:"id,omitempty"`
}

type VirtualNetworkPeering_SpecARM struct {
	Id *string `json:"id,omitempty"`
}
