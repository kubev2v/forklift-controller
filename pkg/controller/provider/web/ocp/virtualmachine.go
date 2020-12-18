package ocp

import (
	"errors"
	"github.com/gin-gonic/gin"
	libmodel "github.com/konveyor/controller/pkg/inventory/model"
	api "github.com/konveyor/forklift-controller/pkg/apis/forklift/v1alpha1"
	model "github.com/konveyor/forklift-controller/pkg/controller/provider/model/ocp"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/base"
	kubevirt "kubevirt.io/client-go/api/v1"
	"net/http"
)

//
// Routes.
const (
	VirtualMachineParam    = "virtualmachine"
	VirtualMachinesRoot    = NamespaceRoot + "/virtualmachines"
	AllVirtualMachinesRoot = ProviderRoot + "/virtualmachines"
	VirtualMachineRoot     = VirtualMachinesRoot + "/:" +
		VirtualMachineParam
)

//
// VirtualMachine handler.
type VirtualMachineHandler struct {
	Handler
}

//
// Add routes to the `gin` router.
func (h *VirtualMachineHandler) AddRoutes(e *gin.Engine) {
	e.GET(AllVirtualMachinesRoot, h.ListAll)
	e.GET(VirtualMachinesRoot, h.List)
	e.GET(VirtualMachinesRoot+"/", h.List)
	e.GET(VirtualMachineRoot, h.Get)
}

//
// List resources in a REST collection (all namespaces).
func (h VirtualMachineHandler) ListAll(ctx *gin.Context) {
	status := h.Prepare(ctx)
	if status != http.StatusOK {
		ctx.Status(status)
		return
	}
	db := h.Reconciler.DB()
	list := []model.VirtualMachine{}
	err := db.List(
		&list,
		libmodel.ListOptions{
			Page: &h.Page,
		})
	if err != nil {
		Log.Trace(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}
	content := []interface{}{}
	for _, m := range list {
		r := &VirtualMachine{}
		r.With(&m)
		r.SelfLink = h.Link(h.Provider, &m)
		content = append(content, r.Content(h.Detail))
	}

	ctx.JSON(http.StatusOK, content)
}

//
// List resources in a REST collection.
func (h VirtualMachineHandler) List(ctx *gin.Context) {
	status := h.Prepare(ctx)
	if status != http.StatusOK {
		ctx.Status(status)
		return
	}
	db := h.Reconciler.DB()
	list := []model.VirtualMachine{}
	err := db.List(
		&list,
		libmodel.ListOptions{
			Predicate: libmodel.Eq("Namespace", ctx.Param(Ns2Param)),
			Page:      &h.Page,
		})
	if err != nil {
		Log.Trace(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}
	content := []interface{}{}
	for _, m := range list {
		r := &VirtualMachine{}
		r.With(&m)
		r.SelfLink = h.Link(h.Provider, &m)
		content = append(content, r.Content(h.Detail))
	}

	ctx.JSON(http.StatusOK, content)
}

//
// Get a specific REST resource.
func (h VirtualMachineHandler) Get(ctx *gin.Context) {
	status := h.Prepare(ctx)
	if status != http.StatusOK {
		ctx.Status(status)
		return
	}
	m := &model.VirtualMachine{
		Base: model.Base{
			Namespace: ctx.Param(Ns2Param),
			Name:      ctx.Param(VirtualMachineParam),
		},
	}
	db := h.Reconciler.DB()
	err := db.Get(m)
	if errors.Is(err, model.NotFound) {
		ctx.Status(http.StatusNotFound)
		return
	}
	if err != nil {
		Log.Trace(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}
	r := &VirtualMachine{}
	r.With(m)
	r.SelfLink = h.Link(h.Provider, m)
	content := r.Content(true)

	ctx.JSON(http.StatusOK, content)
}

//
// Build self link (URI).
func (h VirtualMachineHandler) Link(p *api.Provider, m *model.VirtualMachine) string {
	return h.Handler.Link(
		VirtualMachineRoot,
		base.Params{
			base.NsParam:        p.Namespace,
			base.ProviderParam:  p.Name,
			Ns2Param:            m.Namespace,
			VirtualMachineParam: m.Name,
		})
}

//
// REST Resource.
type VirtualMachine struct {
	Resource
	Object kubevirt.VirtualMachine `json:"object"`
}

//
// Set fields with the specified object.
func (r *VirtualMachine) With(m *model.VirtualMachine) {
	r.Resource.With(&m.Base)
	m.DecodeObject(&r.Object)
}

//
// As content.
func (r *VirtualMachine) Content(detail bool) interface{} {
	if !detail {
		return r.Resource
	}

	return r
}
