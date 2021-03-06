package controllers

import (
	"github.com/drborges/wetagit/api/models"
	"github.com/go-martini/martini"
)

var Tags = tags{}

type tags struct {
	Controller
}

func (this *tags) List() {
	owner := this.Query.Get("owner")
	tags := models.Tags{}

	if err := this.Datastore.Query(tags.ByOwner(owner)).Results(&tags); err != nil {
		this.RenderOkMessage("Count not find tags for owner %v", owner)
		return
	}

	this.RenderData(tags)
}

func (this *tags) Create(tag models.Tag) {
	if err := this.CachedDatastore.Create(&tag); err != nil {
		this.RenderInternalServerErrorMessage(err.Error())
		return
	}

	this.RenderCreatedResource(&tag)
}

func (this *tags) Retrieve(params martini.Params) {
	tag := &models.Tag{Name: params["id"]}

	if err := this.CachedDatastore.Load(tag); err != nil {
		this.RenderStatusNotFoundMessage("Could not find tag for id %v", tag.ResourceID())
		return
	}

	this.RenderData(tag)
}

func (this *tags) Remove(params martini.Params) {
	tag := &models.Tag{Name: params["id"]}

	if err := this.CachedDatastore.Delete(tag); err != nil {
		this.RenderInternalServerErrorMessage(err.Error())
		return
	}

	this.RenderOkMessage("Tag %v successfully removed", tag.ResourceID())
}
