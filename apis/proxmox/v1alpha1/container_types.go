// File: apis/proxmox/v1alpha1/container_types.go

/*
Licensed under the Apache License, Version 2.0...
*/

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ContainerParameters define the desired state of a Proxmox Container
type ContainerParameters struct {
	Node       string `json:"node"`
	Hostname   string `json:"hostname"`
	OSTemplate string `json:"osTemplate"`
	Memory     int    `json:"memory"`
	Swap       int    `json:"swap,omitempty"`
	Cores      int    `json:"cores"`
	// Add other fields from the Terraform schema
}

// ContainerObservation defines the observed state of a Proxmox Container
type ContainerObservation struct {
	ID     string `json:"id,omitempty"`
	Status string `json:"status,omitempty"`
	// Add other observed state fields
}

// ContainerSpec defines the desired state of Container
type ContainerSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ContainerParameters `json:"forProvider"`
}

// ContainerStatus defines the observed state of Container
type ContainerStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ContainerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,proxmox}

// Container is the Schema for the containers API
type Container struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ContainerSpec   `json:"spec"`
	Status ContainerStatus `json:"status,omitempty"`
}

// Implementing resource.Managed interface
var _ resource.Managed = &Container{}

// GetCondition of this Container
func (c *Container) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return c.Status.GetCondition(ct)
}

// SetConditions of this Container
func (c *Container) SetConditions(conditions ...xpv1.Condition) {
	c.Status.SetConditions(conditions...)
}

// GetProviderConfigReference of this Container
func (c *Container) GetProviderConfigReference() *xpv1.Reference {
	return c.Spec.ProviderConfigReference
}

// SetProviderConfigReference of this Container
func (c *Container) SetProviderConfigReference(r *xpv1.Reference) {
	c.Spec.ProviderConfigReference = r
}

// Deprecated: Use GetProviderConfigReference.
func (c *Container) GetProviderReference() *xpv1.Reference {
	return c.Spec.ProviderReference
}

// Deprecated: Use SetProviderConfigReference.
func (c *Container) SetProviderReference(r *xpv1.Reference) {
	c.Spec.ProviderReference = r
}

// GetWriteConnectionSecretToReference of this Container
func (c *Container) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return c.Spec.WriteConnectionSecretToReference
}

// SetWriteConnectionSecretToReference of this Container
func (c *Container) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	c.Spec.WriteConnectionSecretToReference = r
}

// GetDeletionPolicy of this Container
func (c *Container) GetDeletionPolicy() xpv1.DeletionPolicy {
	return c.Spec.DeletionPolicy
}

// SetDeletionPolicy of this Container
func (c *Container) SetDeletionPolicy(dp xpv1.DeletionPolicy) {
	c.Spec.DeletionPolicy = dp
}

// +kubebuilder:object:root=true

// ContainerList contains a list of Container
type ContainerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Container `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Container{}, &ContainerList{})
}
