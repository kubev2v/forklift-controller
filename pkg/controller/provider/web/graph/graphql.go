package graph

import (
	"github.com/gin-gonic/gin"
	"github.com/konveyor/controller/pkg/inventory/container"
	"github.com/konveyor/controller/pkg/logging"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/base"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/generated"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/resolver"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/resolver/vsphere/cluster"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/resolver/vsphere/datacenter"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/resolver/vsphere/datastore"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/resolver/vsphere/host"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/resolver/vsphere/network"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/resolver/vsphere/provider"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/resolver/vsphere/vm"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

//
// Routes.
const (
	GraphqlRoot = "graphql"
)

//
// GraphQL handler.
type GraphHandler struct {
	base.Handler
}

//
// Add routes to the `gin` router.
func (h *GraphHandler) AddRoutes(e *gin.Engine) {
	e.POST(GraphqlRoot, h.Post)
	e.GET(GraphqlRoot+"/playground", h.Get)
}

func newBaseResolver(c *container.Container, name string) resolver.Resolver {
	return resolver.Resolver{
		Container: c,
		Log:       logging.WithName("graphql|" + name),
	}
}

//
// GraphQL Queries handler.
func (h GraphHandler) Post(ctx *gin.Context) {
	config := generated.Config{
		Resolvers: &Resolver{
			Provider: provider.Resolver{
				Resolver: newBaseResolver(h.Container, "provider"),
			},
			Datacenter: datacenter.Resolver{
				Resolver: newBaseResolver(h.Container, "datacenter"),
			},
			Cluster: cluster.Resolver{
				Resolver: newBaseResolver(h.Container, "cluster"),
			},
			Host: host.Resolver{
				Resolver: newBaseResolver(h.Container, "host"),
			},
			Datastore: datastore.Resolver{
				Resolver: newBaseResolver(h.Container, "datastore"),
			},
			Network: network.Resolver{
				Resolver: newBaseResolver(h.Container, "network"),
			},
			VM: vm.Resolver{
				Resolver: newBaseResolver(h.Container, "vm"),
			},
		},
	}

	handler := handler.NewDefaultServer(generated.NewExecutableSchema(config))

	handler.ServeHTTP(ctx.Writer, ctx.Request)
}

//
// GraphQL Playground plugin handler.
func (h GraphHandler) Get(ctx *gin.Context) {
	handler := playground.Handler("GraphQL Playground", "/query")
	handler.ServeHTTP(ctx.Writer, ctx.Request)
}
