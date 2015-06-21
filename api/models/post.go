package models

import (
	"appengine/datastore"
	"github.com/drborges/appx"
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

type Posts []*Post

func (this Posts) ByAuthor(author string) *datastore.Query {
	return appx.From(&Post{}).Filter("Author=", author)
}

func (this Posts) ByTagNames(tagNames []string) *datastore.Query {
	q := appx.From(&Post{})
	for _, name := range tagNames {
		q = q.Filter("Tags=", name)
	}
	return q
}
