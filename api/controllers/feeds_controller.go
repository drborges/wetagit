package controllers

import (
	"github.com/drborges/wetagit/api/models"
	"github.com/go-martini/martini"
)

var Feeds = feeds_controller{}

type feeds_controller struct {
	Controller
}

func (this *feeds_controller) Fetch(params martini.Params) {
	feedID := params["id"]
	feed := models.NewFeed(feedID)

	if err := feed.Load(this.Datastore); err != nil {
		this.RenderStatusNotFoundMessage("Was not able to fetch feed %v. Error: %v", feedID, err)
		return
	}

	this.RenderData(feed)
}
