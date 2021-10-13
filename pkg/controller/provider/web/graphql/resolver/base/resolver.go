package base

import (
	"fmt"

	libcontainer "github.com/konveyor/controller/pkg/inventory/container"
	libmodel "github.com/konveyor/controller/pkg/inventory/model"
	"github.com/konveyor/controller/pkg/logging"
	api "github.com/konveyor/forklift-controller/pkg/apis/forklift/v1beta1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

type DB interface {
	GetDB(provider string) *libmodel.DB
}

type Resolver struct {
	Container *libcontainer.Container
	Log       *logging.Logger
	DB
}

func (t *Resolver) GetDB(provider string) *libmodel.DB {
	p := &api.Provider{
		ObjectMeta: meta.ObjectMeta{
			UID: types.UID(provider),
		},
	}

	var found bool
	var collector libcontainer.Collector
	if collector, found = t.Container.Get(p); !found {
		msg := fmt.Sprintf("provider '%s' not found", provider)
		t.Log.Info(msg)
		return nil
	}

	db := collector.DB()
	return &db
}