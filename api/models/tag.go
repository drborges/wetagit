package models

import (
	"appengine"
	"appengine/datastore"
)

type Tag struct {
	Value   string `json:"value"`
}

func (this *Tag) Key(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, "Tags", this.Value, 0, nil)
}
