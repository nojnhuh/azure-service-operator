// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210501

// Deprecated version of ManagedClusters_AgentPool_STATUS. Use v1beta20210501.ManagedClusters_AgentPool_STATUS instead
type ManagedClusters_AgentPool_STATUS_ARM struct {
	Id         *string                                              `json:"id,omitempty"`
	Name       *string                                              `json:"name,omitempty"`
	Properties *ManagedClusterAgentPoolProfileProperties_STATUS_ARM `json:"properties,omitempty"`
	Type       *string                                              `json:"type,omitempty"`
}

// Deprecated version of ManagedClusterAgentPoolProfileProperties_STATUS. Use v1beta20210501.ManagedClusterAgentPoolProfileProperties_STATUS instead
type ManagedClusterAgentPoolProfileProperties_STATUS_ARM struct {
	AvailabilityZones         []string                             `json:"availabilityZones,omitempty"`
	Count                     *int                                 `json:"count,omitempty"`
	EnableAutoScaling         *bool                                `json:"enableAutoScaling,omitempty"`
	EnableEncryptionAtHost    *bool                                `json:"enableEncryptionAtHost,omitempty"`
	EnableFIPS                *bool                                `json:"enableFIPS,omitempty"`
	EnableNodePublicIP        *bool                                `json:"enableNodePublicIP,omitempty"`
	EnableUltraSSD            *bool                                `json:"enableUltraSSD,omitempty"`
	GpuInstanceProfile        *GPUInstanceProfile_STATUS           `json:"gpuInstanceProfile,omitempty"`
	KubeletConfig             *KubeletConfig_STATUS_ARM            `json:"kubeletConfig,omitempty"`
	KubeletDiskType           *KubeletDiskType_STATUS              `json:"kubeletDiskType,omitempty"`
	LinuxOSConfig             *LinuxOSConfig_STATUS_ARM            `json:"linuxOSConfig,omitempty"`
	MaxCount                  *int                                 `json:"maxCount,omitempty"`
	MaxPods                   *int                                 `json:"maxPods,omitempty"`
	MinCount                  *int                                 `json:"minCount,omitempty"`
	Mode                      *AgentPoolMode_STATUS                `json:"mode,omitempty"`
	NodeImageVersion          *string                              `json:"nodeImageVersion,omitempty"`
	NodeLabels                map[string]string                    `json:"nodeLabels,omitempty"`
	NodePublicIPPrefixID      *string                              `json:"nodePublicIPPrefixID,omitempty"`
	NodeTaints                []string                             `json:"nodeTaints,omitempty"`
	OrchestratorVersion       *string                              `json:"orchestratorVersion,omitempty"`
	OsDiskSizeGB              *int                                 `json:"osDiskSizeGB,omitempty"`
	OsDiskType                *OSDiskType_STATUS                   `json:"osDiskType,omitempty"`
	OsSKU                     *OSSKU_STATUS                        `json:"osSKU,omitempty"`
	OsType                    *OSType_STATUS                       `json:"osType,omitempty"`
	PodSubnetID               *string                              `json:"podSubnetID,omitempty"`
	PowerState                *PowerState_STATUS_ARM               `json:"powerState,omitempty"`
	ProvisioningState         *string                              `json:"provisioningState,omitempty"`
	ProximityPlacementGroupID *string                              `json:"proximityPlacementGroupID,omitempty"`
	ScaleSetEvictionPolicy    *ScaleSetEvictionPolicy_STATUS       `json:"scaleSetEvictionPolicy,omitempty"`
	ScaleSetPriority          *ScaleSetPriority_STATUS             `json:"scaleSetPriority,omitempty"`
	SpotMaxPrice              *float64                             `json:"spotMaxPrice,omitempty"`
	Tags                      map[string]string                    `json:"tags,omitempty"`
	Type                      *AgentPoolType_STATUS                `json:"type,omitempty"`
	UpgradeSettings           *AgentPoolUpgradeSettings_STATUS_ARM `json:"upgradeSettings,omitempty"`
	VmSize                    *string                              `json:"vmSize,omitempty"`
	VnetSubnetID              *string                              `json:"vnetSubnetID,omitempty"`
}

// Deprecated version of AgentPoolUpgradeSettings_STATUS. Use v1beta20210501.AgentPoolUpgradeSettings_STATUS instead
type AgentPoolUpgradeSettings_STATUS_ARM struct {
	MaxSurge *string `json:"maxSurge,omitempty"`
}

// Deprecated version of KubeletConfig_STATUS. Use v1beta20210501.KubeletConfig_STATUS instead
type KubeletConfig_STATUS_ARM struct {
	AllowedUnsafeSysctls  []string `json:"allowedUnsafeSysctls,omitempty"`
	ContainerLogMaxFiles  *int     `json:"containerLogMaxFiles,omitempty"`
	ContainerLogMaxSizeMB *int     `json:"containerLogMaxSizeMB,omitempty"`
	CpuCfsQuota           *bool    `json:"cpuCfsQuota,omitempty"`
	CpuCfsQuotaPeriod     *string  `json:"cpuCfsQuotaPeriod,omitempty"`
	CpuManagerPolicy      *string  `json:"cpuManagerPolicy,omitempty"`
	FailSwapOn            *bool    `json:"failSwapOn,omitempty"`
	ImageGcHighThreshold  *int     `json:"imageGcHighThreshold,omitempty"`
	ImageGcLowThreshold   *int     `json:"imageGcLowThreshold,omitempty"`
	PodMaxPids            *int     `json:"podMaxPids,omitempty"`
	TopologyManagerPolicy *string  `json:"topologyManagerPolicy,omitempty"`
}

// Deprecated version of LinuxOSConfig_STATUS. Use v1beta20210501.LinuxOSConfig_STATUS instead
type LinuxOSConfig_STATUS_ARM struct {
	SwapFileSizeMB             *int                     `json:"swapFileSizeMB,omitempty"`
	Sysctls                    *SysctlConfig_STATUS_ARM `json:"sysctls,omitempty"`
	TransparentHugePageDefrag  *string                  `json:"transparentHugePageDefrag,omitempty"`
	TransparentHugePageEnabled *string                  `json:"transparentHugePageEnabled,omitempty"`
}

// Deprecated version of SysctlConfig_STATUS. Use v1beta20210501.SysctlConfig_STATUS instead
type SysctlConfig_STATUS_ARM struct {
	FsAioMaxNr                     *int    `json:"fsAioMaxNr,omitempty"`
	FsFileMax                      *int    `json:"fsFileMax,omitempty"`
	FsInotifyMaxUserWatches        *int    `json:"fsInotifyMaxUserWatches,omitempty"`
	FsNrOpen                       *int    `json:"fsNrOpen,omitempty"`
	KernelThreadsMax               *int    `json:"kernelThreadsMax,omitempty"`
	NetCoreNetdevMaxBacklog        *int    `json:"netCoreNetdevMaxBacklog,omitempty"`
	NetCoreOptmemMax               *int    `json:"netCoreOptmemMax,omitempty"`
	NetCoreRmemDefault             *int    `json:"netCoreRmemDefault,omitempty"`
	NetCoreRmemMax                 *int    `json:"netCoreRmemMax,omitempty"`
	NetCoreSomaxconn               *int    `json:"netCoreSomaxconn,omitempty"`
	NetCoreWmemDefault             *int    `json:"netCoreWmemDefault,omitempty"`
	NetCoreWmemMax                 *int    `json:"netCoreWmemMax,omitempty"`
	NetIpv4IpLocalPortRange        *string `json:"netIpv4IpLocalPortRange,omitempty"`
	NetIpv4NeighDefaultGcThresh1   *int    `json:"netIpv4NeighDefaultGcThresh1,omitempty"`
	NetIpv4NeighDefaultGcThresh2   *int    `json:"netIpv4NeighDefaultGcThresh2,omitempty"`
	NetIpv4NeighDefaultGcThresh3   *int    `json:"netIpv4NeighDefaultGcThresh3,omitempty"`
	NetIpv4TcpFinTimeout           *int    `json:"netIpv4TcpFinTimeout,omitempty"`
	NetIpv4TcpKeepaliveProbes      *int    `json:"netIpv4TcpKeepaliveProbes,omitempty"`
	NetIpv4TcpKeepaliveTime        *int    `json:"netIpv4TcpKeepaliveTime,omitempty"`
	NetIpv4TcpMaxSynBacklog        *int    `json:"netIpv4TcpMaxSynBacklog,omitempty"`
	NetIpv4TcpMaxTwBuckets         *int    `json:"netIpv4TcpMaxTwBuckets,omitempty"`
	NetIpv4TcpTwReuse              *bool   `json:"netIpv4TcpTwReuse,omitempty"`
	NetIpv4TcpkeepaliveIntvl       *int    `json:"netIpv4TcpkeepaliveIntvl,omitempty"`
	NetNetfilterNfConntrackBuckets *int    `json:"netNetfilterNfConntrackBuckets,omitempty"`
	NetNetfilterNfConntrackMax     *int    `json:"netNetfilterNfConntrackMax,omitempty"`
	VmMaxMapCount                  *int    `json:"vmMaxMapCount,omitempty"`
	VmSwappiness                   *int    `json:"vmSwappiness,omitempty"`
	VmVfsCachePressure             *int    `json:"vmVfsCachePressure,omitempty"`
}
