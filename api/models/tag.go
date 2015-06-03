package models

import (
	"appengine"
	"appengine/datastore"
	"github.com/drborges/wetagit/api/services/db"
)

type Tag struct {
	db.Model
	Value string `json:"value"`
}

func (this *Tag) Kind() string {
	return "Tags"
}

func (this *Tag) NewKey(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, this.Kind(), this.Value, 0, nil)
}
