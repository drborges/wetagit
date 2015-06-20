package models

import (
	"appengine/datastore"
	"github.com/drborges/appx"
)

type Tag struct {
	appx.Model
	Name  string `json:"name"`
	Owner string `json:"owner"`
}

func (this Tag) KeyMetadata() *appx.KeyMetadata {
	return &appx.KeyMetadata{
		Kind:     "Tags",
		StringID: this.Name,
	}
}

func (this Tag) CacheID() string {
	return this.Name
}

func (this Tag) ResourceID() string {
	return this.Name
}

func (this *Tag) SetResourceID(id string) error {
	this.Name = id
	return nil
}

// Warning: It MUST be a slice of pointers for now
// Otherwise, Tag is initialized without a default
// instance of db.Model thus it won't be a db.entity
type Tags []*Tag

func (this Tags) ByOwner(owner string) *datastore.Query {
	return datastore.NewQuery(Tag{}.KeyMetadata().Kind).Filter("Owner=", owner)
}
