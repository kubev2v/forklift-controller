package builder

import (
	liberr "github.com/konveyor/controller/pkg/error"
	api "github.com/konveyor/virt-controller/pkg/apis/virt/v1alpha1"
	"github.com/konveyor/virt-controller/pkg/apis/virt/v1alpha1/plan"
	"github.com/konveyor/virt-controller/pkg/controller/plan/builder/vsphere"
	"github.com/konveyor/virt-controller/pkg/controller/provider/web"
	vmio "github.com/kubevirt/vm-import-operator/pkg/apis/v2v/v1beta1"
	core "k8s.io/api/core/v1"
)

//
// Builder API.
// Builds/updates objects as needed with provider
// specific constructs.
type Builder interface {
	// Build secret.
	Secret(in, object *core.Secret) (err error)
	// Build VMIO resource mapping.
	Mapping(mp *plan.Map, object *vmio.ResourceMapping) error
	// Build VMIO import source.
	Source(vmID string, object *vmio.VirtualMachineImportSourceSpec) error
	// Build tasks.
	Tasks(vmID string) ([]*plan.Task, error)
}

//
// Builder factory.
func New(provider *api.Provider, client web.Client) (builder Builder, err error) {
	switch provider.Type() {
	case api.VSphere:
		builder = &vsphere.Builder{
			Provider: provider,
			Client:   client,
		}
	default:
		liberr.New("provider not supported.")
	}

	return
}