package virtualmachine

import (
	"context"
	"fmt"

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
	cr, ok := mg.(*v1alpha1.VirtualMachine)
	if !ok {
		return nil, errors.New("not a VirtualMachine resource")
	}

	return &external{kube: c.kube}, nil
}

type external struct {
	kube client.Client
}

func (e *external) Observe(ctx context.Context, mg resource.Managed) (managed.ExternalObservation, error) {
	vm, ok := mg.(*v1alpha1.VirtualMachine)
	if !ok {
		return managed.ExternalObservation{}, errors.New("not a VirtualMachine resource")
	}

	fmt.Printf("Observing VM: %s\n", vm.Name)
	return managed.ExternalObservation{
		ResourceExists: false,
	}, nil
}

func (e *external) Create(ctx context.Context, mg resource.Managed) (managed.ExternalCreation, error) {
	vm, ok := mg.(*v1alpha1.VirtualMachine)
	if !ok {
		return managed.ExternalCreation{}, errors.New("not a VirtualMachine resource")
	}

	fmt.Printf("Creating VM: %s\n", vm.Name)
	return managed.ExternalCreation{}, nil
}

func (e *external) Update(ctx context.Context, mg resource.Managed) (managed.ExternalUpdate, error) {
	vm, ok := mg.(*v1alpha1.VirtualMachine)
	if !ok {
		return managed.ExternalUpdate{}, errors.New("not a VirtualMachine resource")
	}

	fmt.Printf("Updating VM: %s\n", vm.Name)
	return managed.ExternalUpdate{}, nil
}

func (e *external) Delete(ctx context.Context, mg resource.Managed) error {
	vm, ok := mg.(*v1alpha1.VirtualMachine)
	if !ok {
		return errors.New("not a VirtualMachine resource")
	}

	fmt.Printf("Deleting VM: %s\n", vm.Name)
	return nil
}
