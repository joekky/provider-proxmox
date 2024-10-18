package virtualmachine

import (
	"context"

	"github.com/pkg/errors"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/crossplane-runtime/pkg/controller"
	"github.com/crossplane/crossplane-runtime/pkg/event"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/pkg/resource"

	"github.com/joekky/provider-proxmox/apis/proxmox/v1alpha1"
)

// Setup adds a controller that reconciles VirtualMachine managed resources.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	name := managed.ControllerName(v1alpha1.VirtualMachineGroupKind)

	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(o.ForControllerRuntime()).
		For(&v1alpha1.VirtualMachine{}).
		Complete(managed.NewReconciler(mgr,
			resource.ManagedKind(v1alpha1.VirtualMachineGroupVersionKind),
			managed.WithExternalConnecter(&connector{
				kube: mgr.GetClient(),
			}),
			managed.WithLogger(o.Logger.WithValues("controller", name)),
			managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name)))))
}

type connector struct {
	kube client.Client
}

func (c *connector) Connect(ctx context.Context, mg resource.Managed) (managed.ExternalClient, error) {
	// Implement connection logic here
	return nil, errors.New("not implemented")
}

type external struct{}

func (e *external) Observe(ctx context.Context, mg resource.Managed) (managed.ExternalObservation, error) {
	// Implement observation logic here
	return managed.ExternalObservation{}, errors.New("not implemented")
}

func (e *external) Create(ctx context.Context, mg resource.Managed) (managed.ExternalCreation, error) {
	// Implement creation logic here
	return managed.ExternalCreation{}, errors.New("not implemented")
}

func (e *external) Update(ctx context.Context, mg resource.Managed) (managed.ExternalUpdate, error) {
	// Implement update logic here
	return managed.ExternalUpdate{}, errors.New("not implemented")
}

func (e *external) Delete(ctx context.Context, mg resource.Managed) error {
	// Implement deletion logic here
	return errors.New("not implemented")
}
