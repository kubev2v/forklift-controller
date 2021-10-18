package graphql

import (
	"fmt"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	libcontainer "github.com/konveyor/controller/pkg/inventory/container"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/generated"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graphql/resolver/vsphere/host"
)

func TestQueryVsphereHost(t *testing.T) {
	c := libcontainer.New()
	fmt.Println("+++++++++++++++++++++++++++++++++========================>")
	fmt.Printf("%+v \n", c)
	fmt.Println("+++++++++++++++++++++++++++++++++========================>")
	config := generated.Config{
		Resolvers: &graph.Resolver{
			Host: host.Resolver{
				Resolver: newBaseResolver(c, "host"),
			},
		},
	}

	client := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(config)))

	t.Run("fails to list hosts with unknown provider", func(t *testing.T) {
		// TODO remove recover once we have a container mock
		defer func() { recover() }()

		var resp struct {
			vsphereHost struct {
				Name     string
				Typename string `json:"__typename"`
			}
		}

		client.MustPost(`{ vsphereHosts(provider:"provider1") { name, __typename } }`, &resp)
	})
}
