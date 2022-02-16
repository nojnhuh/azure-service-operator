// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201201

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type VirtualMachine_SpecARM struct {
	AzureName string `json:"azureName"`

	//ExtendedLocation: The extended location of the Virtual Machine.
	ExtendedLocation *ExtendedLocationARM `json:"extendedLocation,omitempty"`

	//Identity: The identity of the virtual machine, if configured.
	Identity *VirtualMachineIdentityARM `json:"identity,omitempty"`

	//Location: Resource location
	Location string `json:"location"`
	Name     string `json:"name"`

	//Plan: Specifies information about the marketplace image used to create the
	//virtual machine. This element is only used for marketplace images. Before you
	//can use a marketplace image from an API, you must enable the image for
	//programmatic use.  In the Azure portal, find the marketplace image that you want
	//to use and then click Want to deploy programmatically, Get Started ->. Enter any
	//required information and then click Save.
	Plan       *PlanARM                     `json:"plan,omitempty"`
	Properties *VirtualMachinePropertiesARM `json:"properties,omitempty"`

	//Tags: Resource tags
	Tags map[string]string `json:"tags,omitempty"`

	//Zones: The virtual machine zones.
	Zones []string `json:"zones,omitempty"`
}

var _ genruntime.ARMResourceSpec = &VirtualMachine_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-12-01"
func (machine VirtualMachine_SpecARM) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetName returns the Name of the resource
func (machine VirtualMachine_SpecARM) GetName() string {
	return machine.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Compute/virtualMachines"
func (machine VirtualMachine_SpecARM) GetType() string {
	return "Microsoft.Compute/virtualMachines"
}

type VirtualMachineIdentityARM struct {
	//Type: The type of identity used for the virtual machine. The type
	//'SystemAssigned, UserAssigned' includes both an implicitly created identity and
	//a set of user assigned identities. The type 'None' will remove any identities
	//from the virtual machine.
	Type *VirtualMachineIdentityType `json:"type,omitempty"`
}

type VirtualMachinePropertiesARM struct {
	//AdditionalCapabilities: Specifies additional capabilities enabled or disabled on
	//the virtual machine.
	AdditionalCapabilities *AdditionalCapabilitiesARM `json:"additionalCapabilities,omitempty"`

	//AvailabilitySet: Specifies information about the availability set that the
	//virtual machine should be assigned to. Virtual machines specified in the same
	//availability set are allocated to different nodes to maximize availability. For
	//more information about availability sets, see [Manage the availability of
	//virtual
	//machines](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-manage-availability?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json).
	//For more information on Azure planned maintenance, see [Planned maintenance for
	//virtual machines in
	//Azure](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-planned-maintenance?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json)
	//Currently, a VM can only be added to availability set at creation time. The
	//availability set to which the VM is being added should be under the same
	//resource group as the availability set resource. An existing VM cannot be added
	//to an availability set.
	//This property cannot exist along with a non-null
	//properties.virtualMachineScaleSet reference.
	AvailabilitySet *SubResourceARM `json:"availabilitySet,omitempty"`

	//BillingProfile: Specifies the billing related details of a Azure Spot virtual
	//machine.
	//Minimum api-version: 2019-03-01.
	BillingProfile *BillingProfileARM `json:"billingProfile,omitempty"`

	//DiagnosticsProfile: Specifies the boot diagnostic settings state.
	//Minimum api-version: 2015-06-15.
	DiagnosticsProfile *DiagnosticsProfileARM `json:"diagnosticsProfile,omitempty"`

	//EvictionPolicy: Specifies the eviction policy for the Azure Spot virtual machine
	//and Azure Spot scale set.
	//For Azure Spot virtual machines, both 'Deallocate' and 'Delete' are supported
	//and the minimum api-version is 2019-03-01.
	//For Azure Spot scale sets, both 'Deallocate' and 'Delete' are supported and the
	//minimum api-version is 2017-10-30-preview.
	EvictionPolicy *EvictionPolicy `json:"evictionPolicy,omitempty"`

	//ExtensionsTimeBudget: Specifies the time alloted for all extensions to start.
	//The time duration should be between 15 minutes and 120 minutes (inclusive) and
	//should be specified in ISO 8601 format. The default value is 90 minutes
	//(PT1H30M).
	//Minimum api-version: 2020-06-01
	ExtensionsTimeBudget *string `json:"extensionsTimeBudget,omitempty"`

	//HardwareProfile: Specifies the hardware settings for the virtual machine.
	HardwareProfile *HardwareProfileARM `json:"hardwareProfile,omitempty"`

	//Host: Specifies information about the dedicated host that the virtual machine
	//resides in.
	//Minimum api-version: 2018-10-01.
	Host *SubResourceARM `json:"host,omitempty"`

	//HostGroup: Specifies information about the dedicated host group that the virtual
	//machine resides in.
	//Minimum api-version: 2020-06-01.
	//NOTE: User cannot specify both host and hostGroup properties.
	HostGroup *SubResourceARM `json:"hostGroup,omitempty"`

	//LicenseType: Specifies that the image or disk that is being used was licensed
	//on-premises.
	//Possible values for Windows Server operating system are:
	//Windows_Client
	//Windows_Server
	//Possible values for Linux Server operating system are:
	//RHEL_BYOS (for RHEL)
	//SLES_BYOS (for SUSE)
	//For more information, see [Azure Hybrid Use Benefit for Windows
	//Server](https://docs.microsoft.com/azure/virtual-machines/windows/hybrid-use-benefit-licensing)
	//[Azure Hybrid Use Benefit for Linux
	//Server](https://docs.microsoft.com/azure/virtual-machines/linux/azure-hybrid-benefit-linux)
	//Minimum api-version: 2015-06-15
	LicenseType *string `json:"licenseType,omitempty"`

	//NetworkProfile: Specifies the network interfaces of the virtual machine.
	NetworkProfile *NetworkProfileARM `json:"networkProfile,omitempty"`

	//OsProfile: Specifies the operating system settings used while creating the
	//virtual machine. Some of the settings cannot be changed once VM is provisioned.
	OsProfile *OSProfileARM `json:"osProfile,omitempty"`

	//PlatformFaultDomain: Specifies the scale set logical fault domain into which the
	//Virtual Machine will be created. By default, the Virtual Machine will by
	//automatically assigned to a fault domain that best maintains balance across
	//available fault domains.
	//<li>This is applicable only if the 'virtualMachineScaleSet' property of this
	//Virtual Machine is set.<li>The Virtual Machine Scale Set that is referenced,
	//must have 'platformFaultDomainCount' &gt; 1.<li>This property cannot be updated
	//once the Virtual Machine is created.<li>Fault domain assignment can be viewed in
	//the Virtual Machine Instance View.
	//Minimum api‐version: 2020‐12‐01
	PlatformFaultDomain *int `json:"platformFaultDomain,omitempty"`

	//Priority: Specifies the priority for the virtual machine.
	//Minimum api-version: 2019-03-01
	Priority *Priority `json:"priority,omitempty"`

	//ProximityPlacementGroup: Specifies information about the proximity placement
	//group that the virtual machine should be assigned to.
	//Minimum api-version: 2018-04-01.
	ProximityPlacementGroup *SubResourceARM `json:"proximityPlacementGroup,omitempty"`

	//SecurityProfile: Specifies the Security related profile settings for the virtual
	//machine.
	SecurityProfile *SecurityProfileARM `json:"securityProfile,omitempty"`

	//StorageProfile: Specifies the storage settings for the virtual machine disks.
	StorageProfile *StorageProfileARM `json:"storageProfile,omitempty"`

	//VirtualMachineScaleSet: Specifies information about the virtual machine scale
	//set that the virtual machine should be assigned to. Virtual machines specified
	//in the same virtual machine scale set are allocated to different nodes to
	//maximize availability. Currently, a VM can only be added to virtual machine
	//scale set at creation time. An existing VM cannot be added to a virtual machine
	//scale set.
	//This property cannot exist along with a non-null properties.availabilitySet
	//reference.
	//Minimum api‐version: 2019‐03‐01
	VirtualMachineScaleSet *SubResourceARM `json:"virtualMachineScaleSet,omitempty"`
}

type BillingProfileARM struct {
	//MaxPrice: Specifies the maximum price you are willing to pay for a Azure Spot
	//VM/VMSS. This price is in US Dollars.
	//This price will be compared with the current Azure Spot price for the VM size.
	//Also, the prices are compared at the time of create/update of Azure Spot VM/VMSS
	//and the operation will only succeed if  the maxPrice is greater than the current
	//Azure Spot price.
	//The maxPrice will also be used for evicting a Azure Spot VM/VMSS if the current
	//Azure Spot price goes beyond the maxPrice after creation of VM/VMSS.
	//Possible values are:
	//- Any decimal value greater than zero. Example: 0.01538
	//-1 – indicates default price to be up-to on-demand.
	//You can set the maxPrice to -1 to indicate that the Azure Spot VM/VMSS should
	//not be evicted for price reasons. Also, the default max price is -1 if it is not
	//provided by you.
	//Minimum api-version: 2019-03-01.
	MaxPrice *float64 `json:"maxPrice,omitempty"`
}

type DiagnosticsProfileARM struct {
	//BootDiagnostics: Boot Diagnostics is a debugging feature which allows you to
	//view Console Output and Screenshot to diagnose VM status.
	//You can easily view the output of your console log.
	//Azure also enables you to see a screenshot of the VM from the hypervisor.
	BootDiagnostics *BootDiagnosticsARM `json:"bootDiagnostics,omitempty"`
}

type HardwareProfileARM struct {
	//VmSize: Specifies the size of the virtual machine.
	//The enum data type is currently deprecated and will be removed by December 23rd
	//2023.
	//Recommended way to get the list of available sizes is using these APIs:
	//[List all available virtual machine sizes in an availability
	//set](https://docs.microsoft.com/rest/api/compute/availabilitysets/listavailablesizes)
	//[List all available virtual machine sizes in a region](
	//https://docs.microsoft.com/en-us/rest/api/compute/resourceskus/list)
	//[List all available virtual machine sizes for
	//resizing](https://docs.microsoft.com/rest/api/compute/virtualmachines/listavailablesizes).
	//For more information about virtual machine sizes, see [Sizes for virtual
	//machines](https://docs.microsoft.com/en-us/azure/virtual-machines/sizes).
	//The available VM sizes depend on region and availability set.
	VmSize *HardwareProfileVmSize `json:"vmSize,omitempty"`
}

type NetworkProfileARM struct {
	//NetworkInterfaces: Specifies the list of resource Ids for the network interfaces
	//associated with the virtual machine.
	NetworkInterfaces []NetworkInterfaceReferenceARM `json:"networkInterfaces,omitempty"`
}

type OSProfileARM struct {
	//AdminPassword: Specifies the password of the administrator account.
	//Minimum-length (Windows): 8 characters
	//Minimum-length (Linux): 6 characters
	//Max-length (Windows): 123 characters
	//Max-length (Linux): 72 characters
	//Complexity requirements: 3 out of 4 conditions below need to be fulfilled
	//Has lower characters
	//Has upper characters
	//Has a digit
	//Has a special character (Regex match [\W_])
	//Disallowed values: "abc@123", "P@$$w0rd", "P@ssw0rd", "P@ssword123", "Pa$$word",
	//"pass@word1", "Password!", "Password1", "Password22", "iloveyou!"
	//For resetting the password, see [How to reset the Remote Desktop service or its
	//login password in a Windows
	//VM](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-reset-rdp?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json)
	//For resetting root password, see [Manage users, SSH, and check or repair disks
	//on Azure Linux VMs using the VMAccess
	//Extension](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-using-vmaccess-extension?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json#reset-root-password)
	AdminPassword *string `json:"adminPassword,omitempty"`

	//AdminUsername: Specifies the name of the administrator account.
	//This property cannot be updated after the VM is created.
	//Windows-only restriction: Cannot end in "."
	//Disallowed values: "administrator", "admin", "user", "user1", "test", "user2",
	//"test1", "user3", "admin1", "1", "123", "a", "actuser", "adm", "admin2",
	//"aspnet", "backup", "console", "david", "guest", "john", "owner", "root",
	//"server", "sql", "support", "support_388945a0", "sys", "test2", "test3",
	//"user4", "user5".
	//Minimum-length (Linux): 1  character
	//Max-length (Linux): 64 characters
	//Max-length (Windows): 20 characters
	//<li> For root access to the Linux VM, see [Using root privileges on Linux
	//virtual machines in
	//Azure](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-use-root-privileges?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json)
	//<li> For a list of built-in system users on Linux that should not be used in
	//this field, see [Selecting User Names for Linux on
	//Azure](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-usernames?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json)
	AdminUsername *string `json:"adminUsername,omitempty"`

	//AllowExtensionOperations: Specifies whether extension operations should be
	//allowed on the virtual machine.
	//This may only be set to False when no extensions are present on the virtual
	//machine.
	AllowExtensionOperations *bool `json:"allowExtensionOperations,omitempty"`

	//ComputerName: Specifies the host OS name of the virtual machine.
	//This name cannot be updated after the VM is created.
	//Max-length (Windows): 15 characters
	//Max-length (Linux): 64 characters.
	//For naming conventions and restrictions see [Azure infrastructure services
	//implementation
	//guidelines](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-infrastructure-subscription-accounts-guidelines?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json#1-naming-conventions).
	ComputerName *string `json:"computerName,omitempty"`

	//CustomData: Specifies a base-64 encoded string of custom data. The base-64
	//encoded string is decoded to a binary array that is saved as a file on the
	//Virtual Machine. The maximum length of the binary array is 65535 bytes.
	//Note: Do not pass any secrets or passwords in customData property
	//This property cannot be updated after the VM is created.
	//customData is passed to the VM to be saved as a file, for more information see
	//[Custom Data on Azure
	//VMs](https://azure.microsoft.com/en-us/blog/custom-data-and-cloud-init-on-windows-azure/)
	//For using cloud-init for your Linux VM, see [Using cloud-init to customize a
	//Linux VM during
	//creation](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-using-cloud-init?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json)
	CustomData *string `json:"customData,omitempty"`

	//LinuxConfiguration: Specifies the Linux operating system settings on the virtual
	//machine.
	//For a list of supported Linux distributions, see [Linux on Azure-Endorsed
	//Distributions](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-endorsed-distros?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json)
	//For running non-endorsed distributions, see [Information for Non-Endorsed
	//Distributions](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-create-upload-generic?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json).
	LinuxConfiguration *LinuxConfigurationARM `json:"linuxConfiguration,omitempty"`

	//RequireGuestProvisionSignal: Specifies whether the guest provision signal is
	//required to infer provision success of the virtual machine.  Note: This property
	//is for private testing only, and all customers must not set the property to
	//false.
	RequireGuestProvisionSignal *bool `json:"requireGuestProvisionSignal,omitempty"`

	//Secrets: Specifies set of certificates that should be installed onto the virtual
	//machine.
	Secrets []VaultSecretGroupARM `json:"secrets,omitempty"`

	//WindowsConfiguration: Specifies Windows operating system settings on the virtual
	//machine.
	WindowsConfiguration *WindowsConfigurationARM `json:"windowsConfiguration,omitempty"`
}

type SecurityProfileARM struct {
	//EncryptionAtHost: This property can be used by user in the request to enable or
	//disable the Host Encryption for the virtual machine or virtual machine scale
	//set. This will enable the encryption for all the disks including Resource/Temp
	//disk at host itself.
	//Default: The Encryption at host will be disabled unless this property is set to
	//true for the resource.
	EncryptionAtHost *bool `json:"encryptionAtHost,omitempty"`

	//SecurityType: Specifies the SecurityType of the virtual machine. It is set as
	//TrustedLaunch to enable UefiSettings.
	//Default: UefiSettings will not be enabled unless this property is set as
	//TrustedLaunch.
	SecurityType *SecurityProfileSecurityType `json:"securityType,omitempty"`

	//UefiSettings: Specifies the security settings like secure boot and vTPM used
	//while creating the virtual machine.
	//Minimum api-version: 2020-12-01
	UefiSettings *UefiSettingsARM `json:"uefiSettings,omitempty"`
}

type StorageProfileARM struct {
	//DataDisks: Specifies the parameters that are used to add a data disk to a
	//virtual machine.
	//For more information about disks, see [About disks and VHDs for Azure virtual
	//machines](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-about-disks-vhds?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json).
	DataDisks []DataDiskARM `json:"dataDisks,omitempty"`

	//ImageReference: Specifies information about the image to use. You can specify
	//information about platform images, marketplace images, or virtual machine
	//images. This element is required when you want to use a platform image,
	//marketplace image, or virtual machine image, but is not used in other creation
	//operations.
	ImageReference *ImageReferenceARM `json:"imageReference,omitempty"`

	//OsDisk: Specifies information about the operating system disk used by the
	//virtual machine.
	//For more information about disks, see [About disks and VHDs for Azure virtual
	//machines](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-about-disks-vhds?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json).
	OsDisk *OSDiskARM `json:"osDisk,omitempty"`
}

// +kubebuilder:validation:Enum={"None","SystemAssigned","SystemAssigned, UserAssigned","UserAssigned"}
type VirtualMachineIdentityType string

const (
	VirtualMachineIdentityTypeNone                       = VirtualMachineIdentityType("None")
	VirtualMachineIdentityTypeSystemAssigned             = VirtualMachineIdentityType("SystemAssigned")
	VirtualMachineIdentityTypeSystemAssignedUserAssigned = VirtualMachineIdentityType("SystemAssigned, UserAssigned")
	VirtualMachineIdentityTypeUserAssigned               = VirtualMachineIdentityType("UserAssigned")
)

type BootDiagnosticsARM struct {
	//Enabled: Whether boot diagnostics should be enabled on the Virtual Machine.
	Enabled *bool `json:"enabled,omitempty"`

	//StorageUri: Uri of the storage account to use for placing the console output and
	//screenshot.
	//If storageUri is not specified while enabling boot diagnostics, managed storage
	//will be used.
	StorageUri *string `json:"storageUri,omitempty"`
}

type DataDiskARM struct {
	//Caching: Specifies the caching requirements.
	//Possible values are:
	//None
	//ReadOnly
	//ReadWrite
	//Default: None for Standard storage. ReadOnly for Premium storage
	Caching *Caching `json:"caching,omitempty"`

	//CreateOption: Specifies how the virtual machine should be created.
	//Possible values are:
	//Attach \u2013 This value is used when you are using a specialized disk to create
	//the virtual machine.
	//FromImage \u2013 This value is used when you are using an image to create the
	//virtual machine. If you are using a platform image, you also use the
	//imageReference element described above. If you are using a marketplace image,
	//you  also use the plan element previously described.
	CreateOption CreateOption `json:"createOption"`

	//DetachOption: Specifies the detach behavior to be used while detaching a disk or
	//which is already in the process of detachment from the virtual machine.
	//Supported values: ForceDetach.
	//detachOption: ForceDetach is applicable only for managed data disks. If a
	//previous detachment attempt of the data disk did not complete due to an
	//unexpected failure from the virtual machine and the disk is still not released
	//then use force-detach as a last resort option to detach the disk forcibly from
	//the VM. All writes might not have been flushed when using this detach behavior.
	//This feature is still in preview mode and is not supported for
	//VirtualMachineScaleSet. To force-detach a data disk update toBeDetached to
	//'true' along with setting detachOption: 'ForceDetach'.
	DetachOption *DetachOption `json:"detachOption,omitempty"`

	//DiskSizeGB: Specifies the size of an empty data disk in gigabytes. This element
	//can be used to overwrite the size of the disk in a virtual machine image.
	//This value cannot be larger than 1023 GB
	DiskSizeGB *int `json:"diskSizeGB,omitempty"`

	//Image: The source user image virtual hard disk. The virtual hard disk will be
	//copied before being attached to the virtual machine. If SourceImage is provided,
	//the destination virtual hard drive must not exist.
	Image *VirtualHardDiskARM `json:"image,omitempty"`

	//Lun: Specifies the logical unit number of the data disk. This value is used to
	//identify data disks within the VM and therefore must be unique for each data
	//disk attached to a VM.
	Lun int `json:"lun"`

	//ManagedDisk: The managed disk parameters.
	ManagedDisk *ManagedDiskParametersARM `json:"managedDisk,omitempty"`

	//Name: The disk name.
	Name *string `json:"name,omitempty"`

	//ToBeDetached: Specifies whether the data disk is in process of detachment from
	//the VirtualMachine/VirtualMachineScaleset
	ToBeDetached *bool `json:"toBeDetached,omitempty"`

	//Vhd: The virtual hard disk.
	Vhd *VirtualHardDiskARM `json:"vhd,omitempty"`

	//WriteAcceleratorEnabled: Specifies whether writeAccelerator should be enabled or
	//disabled on the disk.
	WriteAcceleratorEnabled *bool `json:"writeAcceleratorEnabled,omitempty"`
}

type ImageReferenceARM struct {
	Id *string `json:"id,omitempty"`

	//Offer: Specifies the offer of the platform image or marketplace image used to
	//create the virtual machine.
	Offer *string `json:"offer,omitempty"`

	//Publisher: The image publisher.
	Publisher *string `json:"publisher,omitempty"`

	//Sku: The image SKU.
	Sku *string `json:"sku,omitempty"`

	//Version: Specifies the version of the platform image or marketplace image used
	//to create the virtual machine. The allowed formats are Major.Minor.Build or
	//'latest'. Major, Minor, and Build are decimal numbers. Specify 'latest' to use
	//the latest version of an image available at deploy time. Even if you use
	//'latest', the VM image will not automatically update after deploy time even if a
	//new version becomes available.
	Version *string `json:"version,omitempty"`
}

type LinuxConfigurationARM struct {
	//DisablePasswordAuthentication: Specifies whether password authentication should
	//be disabled.
	DisablePasswordAuthentication *bool `json:"disablePasswordAuthentication,omitempty"`

	//PatchSettings: [Preview Feature] Specifies settings related to VM Guest Patching
	//on Linux.
	PatchSettings *LinuxPatchSettingsARM `json:"patchSettings,omitempty"`

	//ProvisionVMAgent: Indicates whether virtual machine agent should be provisioned
	//on the virtual machine.
	//When this property is not specified in the request body, default behavior is to
	//set it to true.  This will ensure that VM Agent is installed on the VM so that
	//extensions can be added to the VM later.
	ProvisionVMAgent *bool `json:"provisionVMAgent,omitempty"`

	//Ssh: Specifies the ssh key configuration for a Linux OS.
	Ssh *SshConfigurationARM `json:"ssh,omitempty"`
}

type NetworkInterfaceReferenceARM struct {
	Id         *string                                 `json:"id,omitempty"`
	Properties *NetworkInterfaceReferencePropertiesARM `json:"properties,omitempty"`
}

type OSDiskARM struct {
	//Caching: Specifies the caching requirements.
	//Possible values are:
	//None
	//ReadOnly
	//ReadWrite
	//Default: None for Standard storage. ReadOnly for Premium storage.
	Caching *Caching `json:"caching,omitempty"`

	//CreateOption: Specifies how the virtual machine should be created.
	//Possible values are:
	//Attach \u2013 This value is used when you are using a specialized disk to create
	//the virtual machine.
	//FromImage \u2013 This value is used when you are using an image to create the
	//virtual machine. If you are using a platform image, you also use the
	//imageReference element described above. If you are using a marketplace image,
	//you  also use the plan element previously described.
	CreateOption CreateOption `json:"createOption"`

	//DiffDiskSettings: Specifies the ephemeral Disk Settings for the operating system
	//disk used by the virtual machine.
	DiffDiskSettings *DiffDiskSettingsARM `json:"diffDiskSettings,omitempty"`

	//DiskSizeGB: Specifies the size of an empty data disk in gigabytes. This element
	//can be used to overwrite the size of the disk in a virtual machine image.
	//This value cannot be larger than 1023 GB
	DiskSizeGB *int `json:"diskSizeGB,omitempty"`

	//EncryptionSettings: Specifies the encryption settings for the OS Disk.
	//Minimum api-version: 2015-06-15
	EncryptionSettings *DiskEncryptionSettingsARM `json:"encryptionSettings,omitempty"`

	//Image: The source user image virtual hard disk. The virtual hard disk will be
	//copied before being attached to the virtual machine. If SourceImage is provided,
	//the destination virtual hard drive must not exist.
	Image *VirtualHardDiskARM `json:"image,omitempty"`

	//ManagedDisk: The managed disk parameters.
	ManagedDisk *ManagedDiskParametersARM `json:"managedDisk,omitempty"`

	//Name: The disk name.
	Name *string `json:"name,omitempty"`

	//OsType: This property allows you to specify the type of the OS that is included
	//in the disk if creating a VM from user-image or a specialized VHD.
	//Possible values are:
	//Windows
	//Linux
	OsType *OSDiskOsType `json:"osType,omitempty"`

	//Vhd: The virtual hard disk.
	Vhd *VirtualHardDiskARM `json:"vhd,omitempty"`

	//WriteAcceleratorEnabled: Specifies whether writeAccelerator should be enabled or
	//disabled on the disk.
	WriteAcceleratorEnabled *bool `json:"writeAcceleratorEnabled,omitempty"`
}

type UefiSettingsARM struct {
	//SecureBootEnabled: Specifies whether secure boot should be enabled on the
	//virtual machine.
	//Minimum api-version: 2020-12-01
	SecureBootEnabled *bool `json:"secureBootEnabled,omitempty"`

	//VTpmEnabled: Specifies whether vTPM should be enabled on the virtual machine.
	//Minimum api-version: 2020-12-01
	VTpmEnabled *bool `json:"vTpmEnabled,omitempty"`
}

type VaultSecretGroupARM struct {
	//SourceVault: The relative URL of the Key Vault containing all of the
	//certificates in VaultCertificates.
	SourceVault *SubResourceARM `json:"sourceVault,omitempty"`

	//VaultCertificates: The list of key vault references in SourceVault which contain
	//certificates.
	VaultCertificates []VaultCertificateARM `json:"vaultCertificates,omitempty"`
}

type WindowsConfigurationARM struct {
	//AdditionalUnattendContent: Specifies additional base-64 encoded XML formatted
	//information that can be included in the Unattend.xml file, which is used by
	//Windows Setup.
	AdditionalUnattendContent []AdditionalUnattendContentARM `json:"additionalUnattendContent,omitempty"`

	//EnableAutomaticUpdates: Indicates whether Automatic Updates is enabled for the
	//Windows virtual machine. Default value is true.
	//For virtual machine scale sets, this property can be updated and updates will
	//take effect on OS reprovisioning.
	EnableAutomaticUpdates *bool `json:"enableAutomaticUpdates,omitempty"`

	//PatchSettings: [Preview Feature] Specifies settings related to VM Guest Patching
	//on Windows.
	PatchSettings *PatchSettingsARM `json:"patchSettings,omitempty"`

	//ProvisionVMAgent: Indicates whether virtual machine agent should be provisioned
	//on the virtual machine.
	//When this property is not specified in the request body, default behavior is to
	//set it to true.  This will ensure that VM Agent is installed on the VM so that
	//extensions can be added to the VM later.
	ProvisionVMAgent *bool `json:"provisionVMAgent,omitempty"`

	//TimeZone: Specifies the time zone of the virtual machine. e.g. "Pacific Standard
	//Time".
	//Possible values can be
	//[TimeZoneInfo.Id](https://docs.microsoft.com/en-us/dotnet/api/system.timezoneinfo.id?#System_TimeZoneInfo_Id)
	//value from time zones returned by
	//[TimeZoneInfo.GetSystemTimeZones](https://docs.microsoft.com/en-us/dotnet/api/system.timezoneinfo.getsystemtimezones).
	TimeZone *string `json:"timeZone,omitempty"`

	//WinRM: Specifies the Windows Remote Management listeners. This enables remote
	//Windows PowerShell.
	WinRM *WinRMConfigurationARM `json:"winRM,omitempty"`
}

type AdditionalUnattendContentARM struct {
	//ComponentName: The component name. Currently, the only allowable value is
	//Microsoft-Windows-Shell-Setup.
	ComponentName *AdditionalUnattendContentComponentName `json:"componentName,omitempty"`

	//Content: Specifies the XML formatted content that is added to the unattend.xml
	//file for the specified path and component. The XML must be less than 4KB and
	//must include the root element for the setting or feature that is being inserted.
	Content *string `json:"content,omitempty"`

	//PassName: The pass name. Currently, the only allowable value is OobeSystem.
	PassName *AdditionalUnattendContentPassName `json:"passName,omitempty"`

	//SettingName: Specifies the name of the setting to which the content applies.
	//Possible values are: FirstLogonCommands and AutoLogon.
	SettingName *AdditionalUnattendContentSettingName `json:"settingName,omitempty"`
}

type DiffDiskSettingsARM struct {
	//Option: Specifies the ephemeral disk settings for operating system disk.
	Option *DiffDiskOption `json:"option,omitempty"`

	//Placement: Specifies the ephemeral disk placement for operating system disk.
	//Possible values are:
	//CacheDisk
	//ResourceDisk
	//Default: CacheDisk if one is configured for the VM size otherwise ResourceDisk
	//is used.
	//Refer to VM size documentation for Windows VM at
	//https://docs.microsoft.com/en-us/azure/virtual-machines/windows/sizes and Linux
	//VM at https://docs.microsoft.com/en-us/azure/virtual-machines/linux/sizes to
	//check which VM sizes exposes a cache disk.
	Placement *DiffDiskPlacement `json:"placement,omitempty"`
}

type DiskEncryptionSettingsARM struct {
	//DiskEncryptionKey: Specifies the location of the disk encryption key, which is a
	//Key Vault Secret.
	DiskEncryptionKey *KeyVaultSecretReferenceARM `json:"diskEncryptionKey,omitempty"`

	//Enabled: Specifies whether disk encryption should be enabled on the virtual
	//machine.
	Enabled *bool `json:"enabled,omitempty"`

	//KeyEncryptionKey: Specifies the location of the key encryption key in Key Vault.
	KeyEncryptionKey *KeyVaultKeyReferenceARM `json:"keyEncryptionKey,omitempty"`
}

type LinuxPatchSettingsARM struct {
	//PatchMode: Specifies the mode of VM Guest Patching to IaaS virtual machine.
	//Possible values are:
	//ImageDefault - The virtual machine's default patching configuration is used.
	//AutomaticByPlatform - The virtual machine will be automatically updated by the
	//platform. The property provisionVMAgent must be true
	PatchMode *LinuxPatchSettingsPatchMode `json:"patchMode,omitempty"`
}

type ManagedDiskParametersARM struct {
	//DiskEncryptionSet: Specifies the customer managed disk encryption set resource
	//id for the managed disk.
	DiskEncryptionSet *SubResourceARM `json:"diskEncryptionSet,omitempty"`
	Id                *string         `json:"id,omitempty"`

	//StorageAccountType: Specifies the storage account type for the managed disk.
	//Managed OS disk storage account type can only be set when you create the scale
	//set. NOTE: UltraSSD_LRS can only be used with data disks, it cannot be used with
	//OS Disk.
	StorageAccountType *StorageAccountType `json:"storageAccountType,omitempty"`
}

type NetworkInterfaceReferencePropertiesARM struct {
	//Primary: Specifies the primary network interface in case the virtual machine has
	//more than 1 network interface.
	Primary *bool `json:"primary,omitempty"`
}

type PatchSettingsARM struct {
	//EnableHotpatching: Enables customers to patch their Azure VMs without requiring
	//a reboot. For enableHotpatching, the 'provisionVMAgent' must be set to true and
	//'patchMode' must be set to 'AutomaticByPlatform'.
	EnableHotpatching *bool `json:"enableHotpatching,omitempty"`

	//PatchMode: Specifies the mode of VM Guest Patching to IaaS virtual machine.
	//Possible values are:
	//Manual - You  control the application of patches to a virtual machine. You do
	//this by applying patches manually inside the VM. In this mode, automatic updates
	//are disabled; the property WindowsConfiguration.enableAutomaticUpdates must be
	//false
	//AutomaticByOS - The virtual machine will automatically be updated by the OS. The
	//property WindowsConfiguration.enableAutomaticUpdates must be true.
	//AutomaticByPlatform - the virtual machine will automatically updated by the
	//platform. The properties provisionVMAgent and
	//WindowsConfiguration.enableAutomaticUpdates must be true
	PatchMode *PatchSettingsPatchMode `json:"patchMode,omitempty"`
}

type SshConfigurationARM struct {
	//PublicKeys: The list of SSH public keys used to authenticate with linux based
	//VMs.
	PublicKeys []SshPublicKeySpecARM `json:"publicKeys,omitempty"`
}

type VaultCertificateARM struct {
	//CertificateStore: For Windows VMs, specifies the certificate store on the
	//Virtual Machine to which the certificate should be added. The specified
	//certificate store is implicitly in the LocalMachine account.
	//For Linux VMs, the certificate file is placed under the /var/lib/waagent
	//directory, with the file name &lt;UppercaseThumbprint&gt;.crt for the X509
	//certificate file and &lt;UppercaseThumbprint&gt;.prv for private key. Both of
	//these files are .pem formatted.
	CertificateStore *string `json:"certificateStore,omitempty"`

	//CertificateUrl: This is the URL of a certificate that has been uploaded to Key
	//Vault as a secret. For adding a secret to the Key Vault, see [Add a key or
	//secret to the key
	//vault](https://docs.microsoft.com/azure/key-vault/key-vault-get-started/#add).
	//In this case, your certificate needs to be It is the Base64 encoding of the
	//following JSON Object which is encoded in UTF-8:
	//{
	//"data":"<Base64-encoded-certificate>",
	//"dataType":"pfx",
	//"password":"<pfx-file-password>"
	//}
	CertificateUrl *string `json:"certificateUrl,omitempty"`
}

type VirtualHardDiskARM struct {
	//Uri: Specifies the virtual hard disk's uri.
	Uri *string `json:"uri,omitempty"`
}

type WinRMConfigurationARM struct {
	//Listeners: The list of Windows Remote Management listeners
	Listeners []WinRMListenerARM `json:"listeners,omitempty"`
}

type KeyVaultKeyReferenceARM struct {
	//KeyUrl: The URL referencing a key encryption key in Key Vault.
	KeyUrl string `json:"keyUrl"`

	//SourceVault: The relative URL of the Key Vault containing the key.
	SourceVault SubResourceARM `json:"sourceVault"`
}

type KeyVaultSecretReferenceARM struct {
	//SecretUrl: The URL referencing a secret in a Key Vault.
	SecretUrl string `json:"secretUrl"`

	//SourceVault: The relative URL of the Key Vault containing the secret.
	SourceVault SubResourceARM `json:"sourceVault"`
}

type SshPublicKeySpecARM struct {
	//KeyData: SSH public key certificate used to authenticate with the VM through
	//ssh. The key needs to be at least 2048-bit and in ssh-rsa format.
	//For creating ssh keys, see [Create SSH keys on Linux and Mac for Linux VMs in
	//Azure](https://docs.microsoft.com/en-us/azure/virtual-machines/linux/mac-create-ssh-keys?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json).
	KeyData *string `json:"keyData,omitempty"`

	//Path: Specifies the full path on the created VM where ssh public key is stored.
	//If the file already exists, the specified key is appended to the file. Example:
	///home/user/.ssh/authorized_keys
	Path *string `json:"path,omitempty"`
}

type WinRMListenerARM struct {
	//CertificateUrl: This is the URL of a certificate that has been uploaded to Key
	//Vault as a secret. For adding a secret to the Key Vault, see [Add a key or
	//secret to the key
	//vault](https://docs.microsoft.com/azure/key-vault/key-vault-get-started/#add).
	//In this case, your certificate needs to be It is the Base64 encoding of the
	//following JSON Object which is encoded in UTF-8:
	//{
	//"data":"<Base64-encoded-certificate>",
	//"dataType":"pfx",
	//"password":"<pfx-file-password>"
	//}
	CertificateUrl *string `json:"certificateUrl,omitempty"`

	//Protocol: Specifies the protocol of WinRM listener.
	//Possible values are:
	//http
	//https
	Protocol *WinRMListenerProtocol `json:"protocol,omitempty"`
}