package host

import (
	"context"
	cnd "github.com/konveyor/controller/pkg/condition"
	liberr "github.com/konveyor/controller/pkg/error"
	libref "github.com/konveyor/controller/pkg/ref"
	api "github.com/konveyor/virt-controller/pkg/apis/virt/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

//
// Types
const (
	ProviderNotValid = "ProviderNotValid"
)

//
// Categories
const (
	Advisory = cnd.Advisory
	Critical = cnd.Critical
	Error    = cnd.Error
	Warn     = cnd.Warn
)

// Reasons
const (
	NotSet   = "NotSet"
	NotFound = "NotFound"
)

// Statuses
const (
	True  = cnd.True
	False = cnd.False
)

// Messages
const (
	ReadyMessage            = "The host is ready."
	ProviderNotValidMessage = "`provider` not valid."
)

//
// Validate the mp resource.
func (r *Reconciler) validate(host *api.Host) error {
	err := r.validateProvider(host)
	if err != nil {
		return liberr.Wrap(err)
	}

	return nil
}

//
// Validate provider field.
func (r *Reconciler) validateProvider(host *api.Host) error {
	ref := host.Spec.Provider
	if !libref.RefSet(&ref) {
		host.Status.SetCondition(
			cnd.Condition{
				Type:     ProviderNotValid,
				Status:   True,
				Reason:   NotSet,
				Category: Critical,
				Message:  ProviderNotValidMessage,
			})
	} else {
		provider := &api.Provider{}
		key := client.ObjectKey{
			Namespace: ref.Namespace,
			Name:      ref.Name,
		}
		err := r.Get(context.TODO(), key, provider)
		if errors.IsNotFound(err) {
			err = nil
			host.Status.SetCondition(
				cnd.Condition{
					Type:     ProviderNotValid,
					Status:   True,
					Reason:   NotFound,
					Category: Critical,
					Message:  ProviderNotValidMessage,
				})
		}
		if err != nil {
			return liberr.Wrap(err)
		}
	}

	return nil
}