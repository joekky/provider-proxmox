// internal/controller/proxmox.go

package controller

import (
	"github.com/crossplane/crossplane-runtime/pkg/controller"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/joekky/provider-proxmox/internal/controller/config"
	// Import other controllers as needed
)

// Setup creates all Proxmox controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	if err := config.Setup(mgr, o); err != nil {
		return err
	}
	// Set up other controllers here, e.g., virtual machine, container, etc.
	return nil
}
