package base

import (
	"errors"
	"fmt"

	libcontainer "github.com/konveyor/controller/pkg/inventory/container"
	libmodel "github.com/konveyor/controller/pkg/inventory/model"
	"github.com/konveyor/controller/pkg/logging"
	api "github.com/konveyor/forklift-controller/pkg/apis/forklift/v1beta1"
	model "github.com/konveyor/forklift-controller/pkg/controller/provider/model/ocp"
	vspheremodel "github.com/konveyor/forklift-controller/pkg/controller/provider/model/vsphere"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

type Resolver struct {
	Container *libcontainer.Container
	Log       *logging.Logger
}

func (t *Resolver) GetDB(provider string) (libmodel.DB, error) {
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
		return nil, errors.New(msg)
	}

	db := collector.DB()
	return db, nil
}

func (t *Resolver) GetChildrenIDs(db libmodel.DB, folderId, kind string) (list []string) {
	folder := &vspheremodel.Folder{
		Base: vspheremodel.Base{
			ID: folderId,
		},
	}

	err := db.Get(folder)
	if err != nil {
		return nil
	}

	for _, c := range folder.Children {
		if c.Kind == kind {
			list = append(list, c.ID)
		}
		if c.Kind == "Folder" {
			list = append(list, t.GetChildrenIDs(db, c.ID, kind)...)
		}
	}

	return list
}

//
// Get DB for all providers.
func (t *Resolver) GetDBs() (map[string]libmodel.DB, error) {
	var providers = map[string]libmodel.DB{}

	list := t.Container.List()

	for _, collector := range list {
		if p, cast := collector.Owner().(*api.Provider); cast {
			if p.Type() == api.VSphere {
				m := &model.Provider{}
				m.With(p)

				db, err := t.GetDB(m.UID)
				if err != nil {
					return nil, err
				}
				providers[m.UID] = db
			}
		}
	}

	return providers, nil
}
