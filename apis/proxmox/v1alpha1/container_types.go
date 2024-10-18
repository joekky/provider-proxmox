package v1alpha1

import (
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ContainerParameters struct {
    Node string `json:"node"`
    Hostname string `json:"hostname"`
    OSTemplate string `json:"osTemplate"`
    Memory int `json:"memory"`
    Swap int `json:"swap,omitempty"`
    Cores int `json:"cores"`
    // Add other fields from the Terraform schema
}

type ContainerObservation struct {
    ID string `json:"id,omitempty"`
    Status string `json:"status,omitempty"`
    // Add other observed state fields
}

type ContainerSpec struct {
    xpv1.ResourceSpec `json:",inline"`
    ForProvider ContainerParameters `json:"forProvider"`
}

type ContainerStatus struct {
    xpv1.ResourceStatus `json:",inline"`
    AtProvider ContainerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

type Container struct {
    metav1.TypeMeta   `json:",inline"`
    metav1.ObjectMeta `json:"metadata,omitempty"`

    Spec   ContainerSpec   `json:"spec"`
    Status ContainerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

type ContainerList struct {
    metav1.TypeMeta `json:",inline"`
    metav1.ListMeta `json:"metadata,omitempty"`
    Items           []Container `json:"items"`
}