package models

import (
	"appengine/datastore"
	"github.com/drborges/ds"
)

type Post struct {
	ds.Model
	Content string `json:"content"`
	Author  string `json:"author"`
}

func (this Post) KeyMetadata() *ds.KeyMetadata {
	return &ds.KeyMetadata{
		Kind: "Posts",
	}
}

// Warning: It MUST be a slice of pointers for now
// Otherwise, Tag is initialized without a default
// instance of db.Model thus it won't be a db.entity
type Posts []*Post

func (this Posts) ByAuthor(author string) *datastore.Query {
	return datastore.NewQuery(Post{}.KeyMetadata().Kind).Filter("Author=", author)
}
