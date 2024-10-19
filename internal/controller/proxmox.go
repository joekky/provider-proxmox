// internal/controller/proxmox.go

package controller

import (
	"github.com/crossplane/crossplane-runtime/pkg/controller"
	"github.com/joekky/provider-proxmox/internal/controller/config"
	ctrl "sigs.k8s.io/controller-runtime"
	// Import other controllers as needed
)

func Setup(mgr ctrl.Manager, o controller.Options) error {
	if err := config.Setup(mgr, o); err != nil {
		return err
	}
	// Set up other controllers
	return nil
}
