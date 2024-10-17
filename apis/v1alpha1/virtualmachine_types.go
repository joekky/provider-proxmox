package v1alpha1

import (
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// VirtualMachineParameters defines the desired state of VirtualMachine
type VirtualMachineParameters struct {
    // Name is the name of the virtual machine in Proxmox
    Name string `json:"name"`
    
    // Node is the name of the Proxmox node to create the VM on
    Node string `json:"node"`
    
    // CPU is the number of CPU cores
    CPU int `json:"cpu"`
    
    // Memory is the amount of RAM in MB
    Memory int `json:"memory"`
    
    // DiskSize is the size of the main disk in GB
    DiskSize int `json:"diskSize"`
}

// VirtualMachineObservation defines the observed state of VirtualMachine
type VirtualMachineObservation struct {
    // State is the current state of the virtual machine
    State string `json:"state,omitempty"`
    
    // IP is the IP address of the virtual machine
    IP string `json:"ip,omitempty"`
}

// A VirtualMachineSpec defines the desired state of a VirtualMachine.
type VirtualMachineSpec struct {
    xpv1.ResourceSpec `json:",inline"`
    ForProvider       VirtualMachineParameters `json:"forProvider"`
}

// A VirtualMachineStatus represents the observed state of a VirtualMachine.
type VirtualMachineStatus struct {
    xpv1.ResourceStatus `json:",inline"`
    AtProvider          VirtualMachineObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A VirtualMachine is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,proxmox}
type VirtualMachine struct {
    metav1.TypeMeta   `json:",inline"`
    metav1.ObjectMeta `json:"metadata,omitempty"`

    Spec   VirtualMachineSpec   `json:"spec"`
    Status VirtualMachineStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualMachineList contains a list of VirtualMachine
type VirtualMachineList struct {
    metav1.TypeMeta `json:",inline"`
    metav1.ListMeta `json:"metadata,omitempty"`
    Items           []VirtualMachine `json:"items"`
}