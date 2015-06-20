package models

import (
	"appengine/datastore"
	"github.com/drborges/appx"
	"fmt"
)

type Post struct {
	appx.Model
	Content string   `json:"content"`
	Author  string   `json:"author"`
	Tags    []string `json:"tags"`
}

func (this Post) KeyMetadata() *appx.KeyMetadata {
	return &appx.KeyMetadata{
		Kind: "Posts",
	}
}

// Warning: It MUST be a slice of pointers for now
// Otherwise, Tag is initialized without a default
// instance of db.Model thus it won't be a db.entity
type Posts []*Post

func (this Posts) ByAuthor(author string) *datastore.Query {
	return appx.From(&Post{}).Filter("Author=", author)
}

func (this Posts) ByTagNames(tagNames []string) *datastore.Query {
	q := appx.From(&Post{})
	for _, name := range tagNames {
		println(fmt.Sprintf("######### %v", name))
		q = q.Filter("Tags=", name)
	}
	return q
}