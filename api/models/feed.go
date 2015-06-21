package models

import (
	"github.com/drborges/appx"
	"strings"
)

// Feed is not a persistable model thus
// does not use appx.Model
type Feed struct {
	Tags     []string `json:"tags"`
	Posts    Posts    `json:"posts"`
	NextPage string   `json:"next_page"`
}

func NewFeed(id string) *Feed {
	return &Feed{
		Tags: strings.Split(id, " "),
	}
}

func (this *Feed) Load(ds *appx.Datastore) error {
	iter := ds.Query(this.Posts.ByTagNames(this.Tags)).PagesIterator()
	if err := iter.LoadNext(&this.Posts); err != nil {
		return err
	}

	this.NextPage = iter.Cursor()
	return nil
}
