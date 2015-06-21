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

type Tags []*Tag

func (this Tags) ByOwner(owner string) *datastore.Query {
	return appx.From(&Tag{}).Filter("Owner=", owner)
}
