package graph

import (
	"github.com/gin-gonic/gin"
	"github.com/konveyor/controller/pkg/logging"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/base"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

var log = logging.WithName("web|graphql")

//
// Routes.
const (
	GraphqlRoot = Root
)

//
// GraphQL handler.
type GraphHandler struct {
	base.Handler
}

//
// Add routes to the `gin` router.
func (h *GraphHandler) AddRoutes(e *gin.Engine) {
	e.POST(GraphqlRoot, h.GraphqlHandler())
	e.GET(GraphqlRoot+"/playground", playgroundHandler())
}

//
// GraphQL Handler.
func (h GraphHandler) GraphqlHandler() gin.HandlerFunc {
	gqlh := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{}}))

	return func(c *gin.Context) {
		gqlh.ServeHTTP(c.Writer, c.Request)
	}
}

//
// GraphQL Playground plugin Handler.
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL Playground", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
