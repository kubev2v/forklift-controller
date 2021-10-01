package graphql

import (
	"github.com/konveyor/controller/pkg/inventory/container"
	libweb "github.com/konveyor/controller/pkg/inventory/web"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/base"
)

//
// Build all handlers.
func Handlers(container *container.Container) []libweb.RequestHandler {
	return []libweb.RequestHandler{
		&GraphHandler{
			Handler: base.Handler{
				Container: container,
			},
		},
	}
}
