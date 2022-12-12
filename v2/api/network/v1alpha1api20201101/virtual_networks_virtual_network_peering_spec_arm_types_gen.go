// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of VirtualNetworks_VirtualNetworkPeering_Spec. Use v1beta20201101.VirtualNetworks_VirtualNetworkPeering_Spec instead
type VirtualNetworks_VirtualNetworkPeering_Spec_ARM struct {
	Name       string                                     `json:"name,omitempty"`
	Properties *VirtualNetworkPeeringPropertiesFormat_ARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &VirtualNetworks_VirtualNetworkPeering_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-11-01"
func (peering VirtualNetworks_VirtualNetworkPeering_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (peering *VirtualNetworks_VirtualNetworkPeering_Spec_ARM) GetName() string {
	return peering.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/virtualNetworks/virtualNetworkPeerings"
func (peering *VirtualNetworks_VirtualNetworkPeering_Spec_ARM) GetType() string {
	return "Microsoft.Network/virtualNetworks/virtualNetworkPeerings"
}

// Deprecated version of VirtualNetworkPeeringPropertiesFormat. Use v1beta20201101.VirtualNetworkPeeringPropertiesFormat instead
type VirtualNetworkPeeringPropertiesFormat_ARM struct {
	AllowForwardedTraffic     *bool                                               `json:"allowForwardedTraffic,omitempty"`
	AllowGatewayTransit       *bool                                               `json:"allowGatewayTransit,omitempty"`
	AllowVirtualNetworkAccess *bool                                               `json:"allowVirtualNetworkAccess,omitempty"`
	DoNotVerifyRemoteGateways *bool                                               `json:"doNotVerifyRemoteGateways,omitempty"`
	PeeringState              *VirtualNetworkPeeringPropertiesFormat_PeeringState `json:"peeringState,omitempty"`
	RemoteAddressSpace        *AddressSpace_ARM                                   `json:"remoteAddressSpace,omitempty"`
	RemoteBgpCommunities      *VirtualNetworkBgpCommunities_ARM                   `json:"remoteBgpCommunities,omitempty"`
	RemoteVirtualNetwork      *SubResource_ARM                                    `json:"remoteVirtualNetwork,omitempty"`
	UseRemoteGateways         *bool                                               `json:"useRemoteGateways,omitempty"`
}