package models

import (
	"github.com/drborges/appx"
	"strings"
)

type Feed struct {
	Tags     []string `json:"tags"`
	Posts    Posts    `json:"posts"`
	NextPage string   `json:"next_page"`
}

func NewFeed(id string) *Feed {
	tagNamesFromID := func(id string) []string {
		return strings.Split(id, " ")
	}

	return &Feed{
		Tags: tagNamesFromID(id),
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
