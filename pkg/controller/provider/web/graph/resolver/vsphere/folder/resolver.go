package folder

import (
	"errors"
	"fmt"

	libmodel "github.com/konveyor/controller/pkg/inventory/model"
	api "github.com/konveyor/forklift-controller/pkg/apis/forklift/v1beta1"
	vspheremodel "github.com/konveyor/forklift-controller/pkg/controller/provider/model/vsphere"
	graphmodel "github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/model"
	base "github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/resolver"
)

type Resolver struct {
	base.Resolver
}

//
// List all folders.
func (t *Resolver) List(provider *string) ([]*graphmodel.VsphereFolder, error) {
	var folders []*graphmodel.VsphereFolder

	providers, _ := t.GetDBs(provider)
	for provider, db := range providers[api.VSphere] {
		list := []vspheremodel.Folder{}
		listOptions := libmodel.ListOptions{Detail: libmodel.MaxDetail}
		err := db.List(&list, listOptions)
		if err != nil {
			return nil, nil
		}

		for _, m := range list {
			f := t.with(&m, provider)
			folders = append(folders, f)
		}
	}

	return folders, nil
}

//
// Get a specific folder.
func (t *Resolver) Get(id string, provider string) (*graphmodel.VsphereFolder, error) {
	db, err := t.GetDB(provider)
	if err != nil {
		return nil, err
	}
	m := &vspheremodel.Folder{
		Base: vspheremodel.Base{
			ID: id,
		},
	}

	err = db.Get(m)
	if errors.Is(err, vspheremodel.NotFound) {
		msg := fmt.Sprintf("folder '%s' not found", id)
		t.Log.Info(msg)
		return nil, errors.New(msg)
	}

	h := t.with(m, provider)

	return h, nil
}

//
// Get folders for specific IDs.
func (t *Resolver) GetByIDs(ids []string, provider string) ([]graphmodel.VsphereVMGroup, error) {
	var folders []graphmodel.VsphereVMGroup

	db, err := t.GetDB(provider)
	if err != nil {
		return nil, err
	}

	list := []vspheremodel.Folder{}
	listOptions := libmodel.ListOptions{Detail: libmodel.MaxDetail, Predicate: libmodel.Eq("id", ids)}
	err = db.List(&list, listOptions)
	if err != nil {
		return nil, nil
	}

	for _, m := range list {
		c := t.with(&m, provider)
		folders = append(folders, c)
	}

	return folders, nil
}

func (t *Resolver) GetChildren(folderId, providerId string) (children []graphmodel.VsphereFolderGroup) {
	db, err := t.GetDB(providerId)
	if err != nil {
		return nil
	}

	folder := &vspheremodel.Folder{
		Base: vspheremodel.Base{
			ID: folderId,
		},
	}

	err = db.Get(folder)
	if err != nil {
		return nil
	}

	for _, c := range folder.Children {
		if c.Kind == "Folder" {
			var f *graphmodel.VsphereFolder
			f, err = t.Get(c.ID, providerId)
			if err != nil {
				return []graphmodel.VsphereFolderGroup{}
			}
			children = append(children, f)

		}
		if c.Kind == "Cluster" {
			cluster := vspheremodel.Cluster{
				Base: vspheremodel.Base{
					ID: c.ID,
				},
			}
			err = db.Get(&cluster)
			if err != nil {
				return []graphmodel.VsphereFolderGroup{}
			}

			c := t.WithVsphereCluster(&cluster, providerId)
			children = append(children, c)
		}
		if c.Kind == "Datastore" {
			storage := vspheremodel.Datastore{
				Base: vspheremodel.Base{
					ID: c.ID,
				},
			}
			err = db.Get(&storage)
			if err != nil {
				return []graphmodel.VsphereFolderGroup{}
			}

			c := t.WithVsphereStorage(&storage, providerId)
			children = append(children, c)
		}
		if c.Kind == "Network" {
			network := vspheremodel.Network{
				Base: vspheremodel.Base{
					ID: c.ID,
				},
			}
			err = db.Get(&network)
			if err != nil {
				return []graphmodel.VsphereFolderGroup{}
			}

			switch network.Variant {
			case vspheremodel.NetStandard:
				c := t.WithVsphereNetwork(&network, providerId)
				children = append(children, c)
			case vspheremodel.NetDvPortGroup:
				c := t.WithDvPortGroup(&network, providerId)
				children = append(children, c)
			case vspheremodel.NetDvSwitch:
				c := t.WithDvSwitch(&network, providerId)
				children = append(children, c)
			}

		}
		if c.Kind == "VM" {
			vm := vspheremodel.VM{
				Base: vspheremodel.Base{
					ID: c.ID,
				},
			}

			err = db.Get(&vm)
			if err != nil {
				return []graphmodel.VsphereFolderGroup{}
			}

			v := t.WithVsphereVM(&vm, providerId)
			children = append(children, v)
		}
	}

	return children
}

func (t *Resolver) with(m *vspheremodel.Folder, providerId string) (h *graphmodel.VsphereFolder) {
	var childrenIDs []string
	for _, c := range m.Children {
		childrenIDs = append(childrenIDs, c.ID)
	}

	return &graphmodel.VsphereFolder{
		ID:          m.ID,
		Name:        m.Name,
		Provider:    providerId,
		Kind:        "Folder",
		Parent:      m.Parent.ID,
		ChildrenIDs: childrenIDs,
	}
}
