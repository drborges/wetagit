package controllers

import (
	"github.com/drborges/wetagit/api/models"
	"github.com/go-martini/martini"
)

var Feeds = feeds_controller{}

type feeds_controller struct {
	Controller
}

func (this *feeds_controller) Retrieve(params martini.Params) {
	feedID := params["id"]
	feed := models.NewFeed(feedID)

	if err := feed.Load(this.Datastore); err != nil {
		this.RenderStatusNotFoundMessage("Was not able to fetch feed %v. Error: %v", feedID, err)
		return
	}

	this.RenderData(feed)
}

func (this *feeds_controller) Subscribe(params martini.Params, user *models.User) {
	subscription := &models.Subscription{Feed: params["id"]}
	subscription.BelongsTo(user)

	if err := this.CachedDatastore.Create(subscription); err != nil {
		this.RenderStatusNotFoundMessage("Was not able to subscribe to feed %v.", subscription.Feed)
		return
	}

	this.RenderData(subscription)
}

func (this *feeds_controller) Unsubscribe(params martini.Params, user *models.User) {
	subscription := &models.Subscription{Feed: params["id"]}
	subscription.BelongsTo(user)

	if err := this.CachedDatastore.Delete(subscription); err != nil {
		this.RenderStatusNotFoundMessage("Was not able to unsubscribe from feed %v.", subscription.Feed)
		return
	}

	this.RenderData(subscription)
}
