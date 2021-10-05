package graphql

import (
	"github.com/gin-gonic/gin"
	"github.com/konveyor/controller/pkg/logging"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/base"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/generated"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graphql/resolver/vsphere/cluster"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graphql/resolver/vsphere/datacenter"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graphql/resolver/vsphere/host"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graphql/resolver/vsphere/provider"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

var log = logging.WithName("web|graphql")

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
	provider := provider.Repository{
		Container: h.Container,
		Log:       log,
	}

	datacenter := datacenter.Repository{
		Container: h.Container,
		Log:       log,
	}

	cluster := cluster.Repository{
		Container: h.Container,
		Log:       log,
	}

	host := host.Repository{
		Container: h.Container,
		Log:       log,
	}

	config := generated.Config{Resolvers: &graph.Resolver{Provider: provider, Datacenter: datacenter, Cluster: cluster, Host: host}}
	handler := handler.NewDefaultServer(generated.NewExecutableSchema(config))

	handler.ServeHTTP(ctx.Writer, ctx.Request)
}

//
// GraphQL Playground plugin handler.
func (h GraphHandler) Get(ctx *gin.Context) {
	handler := playground.Handler("GraphQL Playground", "/query")
	handler.ServeHTTP(ctx.Writer, ctx.Request)
}
