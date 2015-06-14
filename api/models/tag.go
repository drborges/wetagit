package models

import (
	"appengine/datastore"
	"github.com/drborges/ds"
)

type Tag struct {
	ds.Model
	Name  string `json:"name"`
	Owner string `json:"owner"`
}

func (this Tag) KeyMetadata() *ds.KeyMetadata {
	return &ds.KeyMetadata{
		Kind:     "Tags",
		StringID: this.Name,
	}
}

// Warning: It MUST be a slice of pointers for now
// Otherwise, Tag is initialized without a default
// instance of db.Model thus it won't be a db.entity
type Tags []*Tag

func (this Tags) ByOwner(owner string) *datastore.Query {
	return datastore.NewQuery(Tag{}.KeyMetadata().Kind).Filter("Owner=", owner)
}
