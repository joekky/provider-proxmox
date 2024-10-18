// package container

// // Import necessary packages

// func Setup(mgr ctrl.Manager, o controller.Options) error {
//     // Setup logic similar to VirtualMachine
// }

// type external struct {
//     client ProxmoxClient
// }

// func (e *external) Observe(ctx context.Context, mg resource.Managed) (managed.ExternalObservation, error) {
//     // Implement observation logic
// }

// func (e *external) Create(ctx context.Context, mg resource.Managed) (managed.ExternalCreation, error) {
//     // Implement creation logic
// }

// func (e *external) Update(ctx context.Context, mg resource.Managed) (managed.ExternalUpdate, error) {
//     // Implement update logic
// }

// func (e *external) Delete(ctx context.Context, mg resource.Managed) error {
//     // Implement deletion logic
// }
// -----

package container

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/crossplane-runtime/pkg/connection"
	"github.com/crossplane/crossplane-runtime/pkg/controller"
	"github.com/crossplane/crossplane-runtime/pkg/event"
	"github.com/crossplane/crossplane-runtime/pkg/ratelimiter"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/pkg/resource"

	"github.com/joekky/provider-proxmox/apis/proxmox/v1alpha1"
	apisv1alpha1 "github.com/joekky/provider-proxmox/apis/proxmox/v1alpha1"
	"github.com/joekky/provider-proxmox/internal/controller/features"
)

const (
	errNotContainer = "managed resource is not a Container custom resource"
	errTrackPCUsage = "cannot track ProviderConfig usage"
	errGetPC        = "cannot get ProviderConfig"
	errGetCreds     = "cannot get credentials"
	errNewClient    = "cannot create new Service"
)

func Setup(mgr ctrl.Manager, o controller.Options) error {
	name := managed.ControllerName(v1alpha1.ContainerGroupKind)

	cps := []managed.ConnectionPublisher{managed.NewAPISecretPublisher(mgr.GetClient(), mgr.GetScheme())}
	if o.Features.Enabled(features.EnableAlphaExternalSecretStores) {
		cps = append(cps, connection.NewDetailsManager(mgr.GetClient(), apisv1alpha1.StoreConfigGroupVersionKind))
	}

	r := managed.NewReconciler(mgr,
		resource.ManagedKind(v1alpha1.ContainerGroupVersionKind),
		managed.WithExternalConnecter(&connector{
			kube:  mgr.GetClient(),
			usage: resource.NewProviderConfigUsageTracker(mgr.GetClient(), &apisv1alpha1.ProviderConfigUsage{}),
		}),
		managed.WithLogger(o.Logger.WithValues("controller", name)),
		managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name))),
		managed.WithConnectionPublishers(cps...))

	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(o.ForControllerRuntime()).
		For(&v1alpha1.Container{}).
		Complete(ratelimiter.NewReconciler(name, r, o.GlobalRateLimiter))
}

type connector struct {
	kube  client.Client
	usage resource.Tracker
}

func (c *connector) Connect(ctx context.Context, mg resource.Managed) (managed.ExternalClient, error) {
	cr, ok := mg.(*v1alpha1.Container)
	if !ok {
		return nil, errors.New(errNotContainer)
	}

	if err := c.usage.Track(ctx, mg); err != nil {
		return nil, errors.Wrap(err, errTrackPCUsage)
	}

	pc := &apisv1alpha1.ProviderConfig{}
	if err := c.kube.Get(ctx, types.NamespacedName{Name: cr.GetProviderConfigReference().Name}, pc); err != nil {
		return nil, errors.Wrap(err, errGetPC)
	}

	cd := pc.Spec.Credentials
	data, err := resource.CommonCredentialExtractor(ctx, cd.Source, c.kube, cd.CommonCredentialSelectors)
	if err != nil {
		return nil, errors.Wrap(err, errGetCreds)
	}

	svc, err := newProxmoxService(data)
	if err != nil {
		return nil, errors.Wrap(err, errNewClient)
	}

	return &external{service: svc}, nil
}

type external struct {
	service interface{}
}

func (c *external) Observe(ctx context.Context, mg resource.Managed) (managed.ExternalObservation, error) {
	cr, ok := mg.(*v1alpha1.Container)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errNotContainer)
	}

	fmt.Printf("Observing: %+v", cr)

	// TODO: Implement actual observation logic

	return managed.ExternalObservation{
		ResourceExists:   true,
		ResourceUpToDate: true,
	}, nil
}

func (c *external) Create(ctx context.Context, mg resource.Managed) (managed.ExternalCreation, error) {
	cr, ok := mg.(*v1alpha1.Container)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errNotContainer)
	}

	fmt.Printf("Creating: %+v", cr)

	// TODO: Implement actual creation logic

	return managed.ExternalCreation{}, nil
}

func (c *external) Update(ctx context.Context, mg resource.Managed) (managed.ExternalUpdate, error) {
	cr, ok := mg.(*v1alpha1.Container)
	if !ok {
		return managed.ExternalUpdate{}, errors.New(errNotContainer)
	}

	fmt.Printf("Updating: %+v", cr)

	// TODO: Implement actual update logic

	return managed.ExternalUpdate{}, nil
}

func (c *external) Delete(ctx context.Context, mg resource.Managed) error {
	cr, ok := mg.(*v1alpha1.Container)
	if !ok {
		return errors.New(errNotContainer)
	}

	fmt.Printf("Deleting: %+v", cr)

	// TODO: Implement actual deletion logic

	return nil
}

func newProxmoxService(data []byte) (interface{}, error) {
	// TODO: Implement the actual Proxmox client creation here
	return nil, nil
}

