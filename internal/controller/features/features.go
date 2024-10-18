package features

import (
	"github.com/crossplane/crossplane-runtime/pkg/feature"
)

// Feature flags.
const (
	// EnableAlphaExternalSecretStores enables alpha support for External Secret Stores.
	EnableAlphaExternalSecretStores feature.Flag = "EnableAlphaExternalSecretStores"

	// EnableAlphaManagementPolicies enables alpha support for Management Policies.
	EnableAlphaManagementPolicies feature.Flag = "EnableAlphaManagementPolicies"
	//
	// EnableProxmoxClusterSupport enables support for Proxmox cluster operations.
	EnableProxmoxClusterSupport feature.Flag = "EnableProxmoxClusterSupport"

	// EnableProxmoxBackupFeatures enables advanced backup and restore features.
	EnableProxmoxBackupFeatures feature.Flag = "EnableProxmoxBackupFeatures"
)

// TODO: Add any Proxmox-specific feature flags here
