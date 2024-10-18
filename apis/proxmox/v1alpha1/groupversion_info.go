// +kubebuilder:object:generate=true
// +groupName=proxmox.crossplane.io

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// GroupVersion is the group and version for the Proxmox API.
var GroupVersion = schema.GroupVersion{Group: "proxmox.crossplane.io", Version: "v1alpha1"}

var (
	// SchemeBuilder initializes a scheme builder
	SchemeBuilder = &metav1.SchemeBuilder{Register: AddToScheme}

	// AddToScheme is a method to add this group-version to a runtime.Scheme
	AddToScheme = SchemeBuilder.AddToScheme
)
