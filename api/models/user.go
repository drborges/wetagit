package models

import (
	"appengine/datastore"
	"github.com/drborges/ds"
)

type User struct {
	ds.Model
	Name          string   `json:"username"`
	Device        string   `json:"device"`
	Subscriptions []string `json:"subscriptions"`
}

func (this User) KeyMetadata() *ds.KeyMetadata {
	return &ds.KeyMetadata{
		Kind:     "Users",
		StringID: this.Name,
	}
}

// Warning: It MUST be a slice of pointers for now
// Otherwise, Tag is initialized without a default
// instance of db.Model thus it won't be a db.entity
type Users []*User

func (this Users) ByName(username string) *datastore.Query {
	return datastore.NewQuery(User{}.KeyMetadata().Kind).Filter("Username=", username)
}