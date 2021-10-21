package folder

import (
	"errors"
	"fmt"

	libmodel "github.com/konveyor/controller/pkg/inventory/model"
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

	providers := t.ListDBs(provider)
	for provider, db := range providers {
		list := []vspheremodel.Folder{}
		listOptions := libmodel.ListOptions{Detail: libmodel.MaxDetail}
		err := db.List(&list, listOptions)
		if err != nil {
			return nil, nil
		}

		for _, m := range list {
			f := with(&m)
			f.Provider = provider
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

	h := with(m)
	h.Provider = provider

	return h, nil
}

//
// Get folders for specific IDs.
func (t *Resolver) GetByIDs(ids []string, provider string) ([]graphmodel.VsphereFolderGroup, error) {
	var folders []graphmodel.VsphereFolderGroup

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
		c := with(&m)
		c.Provider = provider
		folders = append(folders, c)
	}

	return folders, nil
}

func with(m *vspheremodel.Folder) (h *graphmodel.VsphereFolder) {
	var children []string
	for _, c := range m.Children {
		children = append(children, c.ID)
	}

	return &graphmodel.VsphereFolder{
		ID:          m.ID,
		Name:        m.Name,
		Parent:      m.Parent.ID,
		ChildrenIDs: children,
	}
}
