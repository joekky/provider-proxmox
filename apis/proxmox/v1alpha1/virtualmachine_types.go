// apis/proxmox/v1alpha1/virtualmachine_types.go

/*
Licensed under the Apache License, Version 2.0...
*/

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// VirtualMachineParameters defines the desired state of VirtualMachine
type VirtualMachineParameters struct {
	// Name is the name of the virtual machine in Proxmox
	Name string `json:"name"`

	// Node is the name of the Proxmox node to create the VM on
	Node string `json:"node"`

	// CPU is the number of CPU cores
	CPU int `json:"cpu"`

	// Memory is the size of RAM in MB
	Memory int `json:"memory"`
}

// VirtualMachineObservation defines the observed state of VirtualMachine
type VirtualMachineObservation struct {
	// ID is the ID of the virtual machine in Proxmox
	ID string `json:"id,omitempty"`
}

// VirtualMachineSpec defines the desired state of a VirtualMachine.
type VirtualMachineSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       VirtualMachineParameters `json:"forProvider"`
}

// VirtualMachineStatus represents the observed state of a VirtualMachine.
type VirtualMachineStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          VirtualMachineObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,proxmox}

// VirtualMachine is the Schema for the virtualmachines API
type VirtualMachine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualMachineSpec   `json:"spec"`
	Status VirtualMachineStatus `json:"status,omitempty"`
}

// Implement resource.Managed interface
var _ resource.Managed = &VirtualMachine{}

// GetCondition of this VirtualMachine
func (vm *VirtualMachine) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return vm.Status.GetCondition(ct)
}

// SetConditions of this VirtualMachine
func (vm *VirtualMachine) SetConditions(c ...xpv1.Condition) {
	vm.Status.SetConditions(c...)
}

// GetProviderConfigReference of this VirtualMachine
func (vm *VirtualMachine) GetProviderConfigReference() *xpv1.Reference {
	return vm.Spec.ProviderConfigReference
}

// SetProviderConfigReference of this VirtualMachine
func (vm *VirtualMachine) SetProviderConfigReference(r *xpv1.Reference) {
	vm.Spec.ProviderConfigReference = r
}

// Deprecated: Use GetProviderConfigReference.
func (vm *VirtualMachine) GetProviderReference() *xpv1.Reference {
	return vm.Spec.ProviderReference
}

// Deprecated: Use SetProviderConfigReference.
func (vm *VirtualMachine) SetProviderReference(r *xpv1.Reference) {
	vm.Spec.ProviderReference = r
}

// GetWriteConnectionSecretToReference of this VirtualMachine
func (vm *VirtualMachine) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return vm.Spec.WriteConnectionSecretToReference
}

// SetWriteConnectionSecretToReference of this VirtualMachine
func (vm *VirtualMachine) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	vm.Spec.WriteConnectionSecretToReference = r
}

// GetDeletionPolicy of this VirtualMachine
func (vm *VirtualMachine) GetDeletionPolicy() xpv1.DeletionPolicy {
	return vm.Spec.DeletionPolicy
}

// SetDeletionPolicy of this VirtualMachine
func (vm *VirtualMachine) SetDeletionPolicy(dp xpv1.DeletionPolicy) {
	vm.Spec.DeletionPolicy = dp
}

// +kubebuilder:object:root=true

// VirtualMachineList contains a list of VirtualMachine
type VirtualMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualMachine `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VirtualMachine{}, &VirtualMachineList{})
}
