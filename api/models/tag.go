package models

import (
	"github.com/drborges/datastore-model"
)

type Tag struct {
	db.Model `db:"Tags"`
	Name  string `json:"name" db:"id"`
	Owner string `json:"owner"`
}

// Warning: It MUST be a slice of pointers for now
// Otherwise, Tag is initialized without a default
// instance of db.Model thus it won't be a db.entity
type Tags []*Tag

func (this Tags) ByOwner(owner string) *db.Query {
	return db.From(new(Tag)).Filter("Owner=", owner)
}
