// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type VirtualNetworksSubnets_SPECARM struct {
	AzureName string  `json:"azureName"`
	Id        *string `json:"id,omitempty"`

	//Name: The name of the resource that is unique within a resource group. This name
	//can be used to access the resource.
	Name string `json:"name,omitempty"`

	//Properties: Properties of the subnet.
	Properties *SubnetPropertiesFormat_SpecARM `json:"properties,omitempty"`

	//Type: Resource type.
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ARMResourceSpec = &VirtualNetworksSubnets_SPECARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-11-01"
func (specarm VirtualNetworksSubnets_SPECARM) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetName returns the Name of the resource
func (specarm VirtualNetworksSubnets_SPECARM) GetName() string {
	return specarm.Name
}

// GetType returns the ARM Type of the resource. This is always ""
func (specarm VirtualNetworksSubnets_SPECARM) GetType() string {
	return ""
}

type SubnetPropertiesFormat_SpecARM struct {
	//AddressPrefix: The address prefix for the subnet.
	AddressPrefix *string `json:"addressPrefix,omitempty"`

	//AddressPrefixes: List of address prefixes for the subnet.
	AddressPrefixes []string `json:"addressPrefixes,omitempty"`

	//ApplicationGatewayIpConfigurations: Application gateway IP configurations of
	//virtual network resource.
	ApplicationGatewayIpConfigurations []ApplicationGatewayIPConfiguration_SpecARM `json:"applicationGatewayIpConfigurations,omitempty"`

	//Delegations: An array of references to the delegations on the subnet.
	Delegations []Delegation_SpecARM `json:"delegations,omitempty"`

	//IpAllocations: Array of IpAllocation which reference this subnet.
	IpAllocations []SubResource_SpecARM `json:"ipAllocations,omitempty"`

	//NatGateway: Nat gateway associated with this subnet.
	NatGateway *SubResource_SpecARM `json:"natGateway,omitempty"`

	//NetworkSecurityGroup: The reference to the NetworkSecurityGroup resource.
	NetworkSecurityGroup *NetworkSecurityGroup_Spec_VirtualNetworksSubnet_SubResourceEmbeddedARM `json:"networkSecurityGroup,omitempty"`

	//PrivateEndpointNetworkPolicies: Enable or Disable apply network policies on
	//private end point in the subnet.
	PrivateEndpointNetworkPolicies *SubnetPropertiesFormatSpecPrivateEndpointNetworkPolicies `json:"privateEndpointNetworkPolicies,omitempty"`

	//PrivateLinkServiceNetworkPolicies: Enable or Disable apply network policies on
	//private link service in the subnet.
	PrivateLinkServiceNetworkPolicies *SubnetPropertiesFormatSpecPrivateLinkServiceNetworkPolicies `json:"privateLinkServiceNetworkPolicies,omitempty"`

	//RouteTable: The reference to the RouteTable resource.
	RouteTable *RouteTable_SpecARM `json:"routeTable,omitempty"`

	//ServiceEndpointPolicies: An array of service endpoint policies.
	ServiceEndpointPolicies []ServiceEndpointPolicy_SpecARM `json:"serviceEndpointPolicies,omitempty"`

	//ServiceEndpoints: An array of service endpoints.
	ServiceEndpoints []ServiceEndpointPropertiesFormat_SpecARM `json:"serviceEndpoints,omitempty"`
}

type ApplicationGatewayIPConfiguration_SpecARM struct {
	Id *string `json:"id,omitempty"`

	//Name: Name of the IP configuration that is unique within an Application Gateway.
	Name *string `json:"name,omitempty"`

	//Properties: Properties of the application gateway IP configuration.
	Properties *ApplicationGatewayIPConfigurationPropertiesFormat_SpecARM `json:"properties,omitempty"`
}

type Delegation_SpecARM struct {
	Id *string `json:"id,omitempty"`

	//Name: The name of the resource that is unique within a subnet. This name can be
	//used to access the resource.
	Name *string `json:"name,omitempty"`

	//Properties: Properties of the subnet.
	Properties *ServiceDelegationPropertiesFormat_SpecARM `json:"properties,omitempty"`

	//Type: Resource type.
	Type *string `json:"type,omitempty"`
}

type NetworkSecurityGroup_Spec_VirtualNetworksSubnet_SubResourceEmbeddedARM struct {
	Id *string `json:"id,omitempty"`

	//Location: Resource location.
	Location *string `json:"location,omitempty"`

	//Properties: Properties of the network security group.
	Properties *NetworkSecurityGroupPropertiesFormat_Spec_VirtualNetworksSubnet_SubResourceEmbeddedARM `json:"properties,omitempty"`

	//Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

type RouteTable_SpecARM struct {
	Id *string `json:"id,omitempty"`

	//Location: Resource location.
	Location *string `json:"location,omitempty"`

	//Properties: Properties of the route table.
	Properties *RouteTablePropertiesFormat_Spec_VirtualNetworksSubnet_SubResourceEmbeddedARM `json:"properties,omitempty"`

	//Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

type ServiceEndpointPolicy_SpecARM struct {
	Id *string `json:"id,omitempty"`

	//Location: Resource location.
	Location *string `json:"location,omitempty"`

	//Properties: Properties of the service end point policy.
	Properties *ServiceEndpointPolicyPropertiesFormat_Spec_VirtualNetworksSubnet_SubResourceEmbeddedARM `json:"properties,omitempty"`

	//Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

type ServiceEndpointPropertiesFormat_SpecARM struct {
	//Locations: A list of locations.
	Locations []string `json:"locations,omitempty"`

	//Service: The type of the endpoint service.
	Service *string `json:"service,omitempty"`
}

type ApplicationGatewayIPConfigurationPropertiesFormat_SpecARM struct {
	//Subnet: Reference to the subnet resource. A subnet from where application
	//gateway gets its private address.
	Subnet *SubResource_SpecARM `json:"subnet,omitempty"`
}

type NetworkSecurityGroupPropertiesFormat_Spec_VirtualNetworksSubnet_SubResourceEmbeddedARM struct {
	//SecurityRules: A collection of security rules of the network security group.
	SecurityRules []SecurityRule_Spec_VirtualNetworksSubnet_SubResourceEmbeddedARM `json:"securityRules,omitempty"`
}

type RouteTablePropertiesFormat_Spec_VirtualNetworksSubnet_SubResourceEmbeddedARM struct {
	//DisableBgpRoutePropagation: Whether to disable the routes learned by BGP on that
	//route table. True means disable.
	DisableBgpRoutePropagation *bool `json:"disableBgpRoutePropagation,omitempty"`

	//Routes: Collection of routes contained within a route table.
	Routes []Route_Spec_VirtualNetworksSubnet_SubResourceEmbeddedARM `json:"routes,omitempty"`
}

type ServiceDelegationPropertiesFormat_SpecARM struct {
	//ServiceName: The name of the service to whom the subnet should be delegated
	//(e.g. Microsoft.Sql/servers).
	ServiceName *string `json:"serviceName,omitempty"`
}

type ServiceEndpointPolicyPropertiesFormat_Spec_VirtualNetworksSubnet_SubResourceEmbeddedARM struct {
	//ServiceEndpointPolicyDefinitions: A collection of service endpoint policy
	//definitions of the service endpoint policy.
	ServiceEndpointPolicyDefinitions []ServiceEndpointPolicyDefinition_Spec_VirtualNetworksSubnet_SubResourceEmbeddedARM `json:"serviceEndpointPolicyDefinitions,omitempty"`
}

type Route_Spec_VirtualNetworksSubnet_SubResourceEmbeddedARM struct {
	Id *string `json:"id,omitempty"`
}

type SecurityRule_Spec_VirtualNetworksSubnet_SubResourceEmbeddedARM struct {
	Id *string `json:"id,omitempty"`
}

type ServiceEndpointPolicyDefinition_Spec_VirtualNetworksSubnet_SubResourceEmbeddedARM struct {
	Id *string `json:"id,omitempty"`
}
