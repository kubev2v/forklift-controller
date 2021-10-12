package graphql

import (
	"github.com/gin-gonic/gin"
	"github.com/konveyor/controller/pkg/logging"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/base"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/generated"
	baseresolver "github.com/konveyor/forklift-controller/pkg/controller/provider/web/graphql/resolver/base"
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

//
// GraphQL Queries handler.
func (h GraphHandler) Post(ctx *gin.Context) {
	provider := provider.Resolver{
		Resolver: baseresolver.Resolver{
			Container: h.Container,
			Log:       logging.WithName("graphql|provider"),
		},
	}

	datacenter := datacenter.Resolver{
		Resolver: baseresolver.Resolver{
			Container: h.Container,
			Log:       logging.WithName("graphql|datacenter"),
		},
	}

	cluster := cluster.Resolver{
		Resolver: baseresolver.Resolver{
			Container: h.Container,
			Log:       logging.WithName("graphql|cluster"),
		},
	}

	host := host.Resolver{
		Resolver: baseresolver.Resolver{
			Container: h.Container,
			Log:       logging.WithName("graphql|host"),
		},
	}

	network := network.Resolver{
		Resolver: baseresolver.Resolver{
			Container: h.Container,
			Log:       logging.WithName("graphql|network"),
		},
	}

	datastore := datastore.Resolver{
		Resolver: baseresolver.Resolver{
			Container: h.Container,
			Log:       logging.WithName("graphql|datastore"),
		},
	}

	vm := vm.Resolver{
		Resolver: baseresolver.Resolver{
			Container: h.Container,
			Log:       logging.WithName("graphql|vm"),
		},
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
