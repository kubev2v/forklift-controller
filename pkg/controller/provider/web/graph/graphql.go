package graph

import (
	"context"

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
	e.POST(GraphqlRoot, h.Post())
	e.GET(GraphqlRoot+"/playground", h.Get())
}

//
// GraphQL Handler.
func (h GraphHandler) Post() gin.HandlerFunc {
	handler := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{}}))

	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "HandlerContainer", h.Container)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
		handler.ServeHTTP(c.Writer, c.Request)
	}
}

//
// GraphQL Playground plugin Handler.
func (h GraphHandler) Get() gin.HandlerFunc {
	handler := playground.Handler("GraphQL Playground", "/query")

	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}
