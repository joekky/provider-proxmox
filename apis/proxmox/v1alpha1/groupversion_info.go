// apis/proxmox/v1alpha1/groupversion_info.go

/*
Licensed under the Apache License, Version 2.0...
*/

// +kubebuilder:object:generate=true
// +groupName=proxmox.crossplane.io

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

var (
	// GroupVersion is group version used to register these objects
	GroupVersion = schema.GroupVersion{Group: "proxmox.crossplane.io", Version: "v1alpha1"}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: GroupVersion}

	// AddToScheme adds the types in this group-version to the given scheme.
	AddToScheme = SchemeBuilder.AddToScheme
)

// Define kind and group kind variables for each type

var (
	// VirtualMachine
	VirtualMachineKind             = "VirtualMachine"
	VirtualMachineGroupKind        = schema.GroupKind{Group: GroupVersion.Group, Kind: VirtualMachineKind}.String()
	VirtualMachineKindAPIVersion   = VirtualMachineKind + "." + GroupVersion.String()
	VirtualMachineGroupVersionKind = GroupVersion.WithKind(VirtualMachineKind)

	// Container
	ContainerKind             = "Container"
	ContainerGroupKind        = schema.GroupKind{Group: GroupVersion.Group, Kind: ContainerKind}.String()
	ContainerKindAPIVersion   = ContainerKind + "." + GroupVersion.String()
	ContainerGroupVersionKind = GroupVersion.WithKind(ContainerKind)

	// ProviderConfig
	ProviderConfigKind             = "ProviderConfig"
	ProviderConfigGroupKind        = schema.GroupKind{Group: GroupVersion.Group, Kind: ProviderConfigKind}.String()
	ProviderConfigKindAPIVersion   = ProviderConfigKind + "." + GroupVersion.String()
	ProviderConfigGroupVersionKind = GroupVersion.WithKind(ProviderConfigKind)

	// StoreConfig (if applicable)
	StoreConfigKind             = "StoreConfig"
	StoreConfigGroupKind        = schema.GroupKind{Group: GroupVersion.Group, Kind: StoreConfigKind}.String()
	StoreConfigKindAPIVersion   = StoreConfigKind + "." + GroupVersion.String()
	StoreConfigGroupVersionKind = GroupVersion.WithKind(StoreConfigKind)
)
