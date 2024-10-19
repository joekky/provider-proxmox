// File: apis/proxmox/v1alpha1/storeconfig_types.go

/*
Licensed under the Apache License, Version 2.0...
*/

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// StoreConfigSpec defines the desired state of a StoreConfig.
type StoreConfigSpec struct {
	xpv1.SecretStoreConfig `json:",inline"`
}

// StoreConfigStatus represents the status of a StoreConfig.
type StoreConfigStatus struct {
	xpv1.ConditionedStatus `json:",inline"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster,categories={crossplane,store,proxmox}
// +kubebuilder:subresource:status

// StoreConfig configures how the controller should store connection details.
type StoreConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StoreConfigSpec   `json:"spec"`
	Status StoreConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StoreConfigList contains a list of StoreConfig
type StoreConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StoreConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&StoreConfig{}, &StoreConfigList{})
}
