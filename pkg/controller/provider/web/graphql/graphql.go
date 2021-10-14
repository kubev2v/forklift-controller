package graphql

import (
	"github.com/gin-gonic/gin"
	"github.com/konveyor/controller/pkg/logging"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/base"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/generated"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graphql/resolver"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graphql/resolver/vsphere/cluster"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graphql/resolver/vsphere/datacenter"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graphql/resolver/vsphere/datastore"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graphql/resolver/vsphere/host"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graphql/resolver/vsphere/network"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graphql/resolver/vsphere/provider"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graphql/resolver/vsphere/vm"

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

func newBaseResolver(h GraphHandler, name string) resolver.Resolver {
	return resolver.Resolver{
		Container: h.Container,
		Log:       logging.WithName("graphql|" + name),
	}
}

//
// GraphQL Queries handler.
func (h GraphHandler) Post(ctx *gin.Context) {
	provider := provider.Resolver{
		Resolver: newBaseResolver(h, "provider"),
	}

	datacenter := datacenter.Resolver{
		Resolver: newBaseResolver(h, "datacenter"),
	}

	cluster := cluster.Resolver{
		Resolver: newBaseResolver(h, "cluster"),
	}

	host := host.Resolver{
		Resolver: newBaseResolver(h, "host"),
	}

	network := network.Resolver{
		Resolver: newBaseResolver(h, "network"),
	}

	datastore := datastore.Resolver{
		Resolver: newBaseResolver(h, "datastore"),
	}

	vm := vm.Resolver{
		Resolver: newBaseResolver(h, "vm"),
	}

	config := generated.Config{Resolvers: &graph.Resolver{Provider: provider, Datacenter: datacenter, Cluster: cluster, Host: host, Datastore: datastore, Network: network, VM: vm}}
	handler := handler.NewDefaultServer(generated.NewExecutableSchema(config))

	handler.ServeHTTP(ctx.Writer, ctx.Request)
}

//
// GraphQL Playground plugin handler.
func (h GraphHandler) Get(ctx *gin.Context) {
	handler := playground.Handler("GraphQL Playground", "/query")
	handler.ServeHTTP(ctx.Writer, ctx.Request)
}
