package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// VirtualMachineParameters defines the desired state of VirtualMachine
type VirtualMachineParameters struct {
	Name   string `json:"name"`
	Node   string `json:"node"`
	CPU    int    `json:"cpu"`
	Memory int    `json:"memory"`
}

// VirtualMachineObservation defines the observed state of VirtualMachine
type VirtualMachineObservation struct {
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

// Ensure VirtualMachine implements resource.Managed interface
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

// GetPublishConnectionDetailsTo returns the publishing configuration of this VirtualMachine.
func (vm *VirtualMachine) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return vm.Spec.PublishConnectionDetailsTo
}

// SetPublishConnectionDetailsTo sets the publishing configuration of this VirtualMachine.
func (vm *VirtualMachine) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	vm.Spec.PublishConnectionDetailsTo = r
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

// GetManagementPolicies of this VirtualMachine
func (vm *VirtualMachine) GetManagementPolicies() xpv1.ManagementPolicies {
	return vm.Spec.ManagementPolicies
}

// SetManagementPolicies of this VirtualMachine
func (vm *VirtualMachine) SetManagementPolicies(p xpv1.ManagementPolicies) {
	vm.Spec.ManagementPolicies = p
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
