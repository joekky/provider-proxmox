// internal/controller/config/providerconfig.go

package config

import (
	"context"

	"github.com/crossplane/crossplane-runtime/pkg/controller"
	"github.com/crossplane/crossplane-runtime/pkg/logging"
	"github.com/crossplane/crossplane-runtime/pkg/ratelimiter"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/joekky/provider-proxmox/apis/proxmox/v1alpha1"
)

// Setup adds a controller that reconciles ProviderConfigs.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	name := "proxmox-providerconfig"

	logger := o.Logger.WithValues("controller", name)

	r := &Reconciler{
		logger: logger,
	}

	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		For(&v1alpha1.ProviderConfig{}).
		Complete(ratelimiter.NewReconciler(name, r, o.GlobalRateLimiter))
}

// Reconciler reconciles a ProviderConfig object
type Reconciler struct {
	logger logging.Logger
}

// Reconcile is the main logic of the reconciler
func (r *Reconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	// Implement your reconciliation logic here.
	// For now, we'll just log that we're reconciling and return.
	r.logger.Debug("Reconciling ProviderConfig", "namespace", req.Namespace, "name", req.Name)
	return ctrl.Result{}, nil
}
