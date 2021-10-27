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
	webvsphere "github.com/konveyor/forklift-controller/pkg/controller/provider/web/vsphere"
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
// Get all providers DBs.
func (t *Resolver) GetDBs(provider *string) (map[string]map[string]libmodel.DB, error) {
	vsphere := map[string]libmodel.DB{}
	ovirt := map[string]libmodel.DB{}
	var allProviders = map[string]map[string]libmodel.DB{
		api.VSphere: vsphere,
		api.OVirt:   ovirt,
	}
	list := t.Container.List()

	for _, collector := range list {
		if p, cast := collector.Owner().(*api.Provider); cast {
			m := &model.Provider{}
			m.With(p)
			r := webvsphere.Provider{}
			r.With(m)

			db, err := t.GetDB(m.UID)
			if err != nil {
				return nil, err
			}

			switch p.Spec.Type {
			case api.VSphere:
				vsphere = make(map[string]libmodel.DB)
				vsphere[m.UID] = db
				allProviders[api.VSphere] = vsphere

			case api.OVirt:
				ovirt = make(map[string]libmodel.DB)
				ovirt[m.UID] = db
				allProviders[api.OVirt] = ovirt
			}
		}
	}

	if provider == nil {
		return allProviders, nil
	}

	providers := t.FindProvider(*provider, allProviders)
	return providers, nil
}

func (t *Resolver) FindProvider(provider string, providers map[string]map[string]libmodel.DB) map[string]map[string]libmodel.DB {
	for providerType, v := range providers {
		for uid := range v {
			if uid == provider {
				found := map[string]map[string]libmodel.DB{
					providerType: v,
				}
				return found
			}
		}
	}
	return nil
}
