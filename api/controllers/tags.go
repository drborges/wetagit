package controllers

import (
	"github.com/drborges/wetagit/api/models"
	"github.com/go-martini/martini"
	"net/http"
	"fmt"
)

var Tags = tags{}

type tags struct {
	Controller
}

func (this *tags) List() {
	owner := this.Query.Get("owner")
	tags := models.Tags{}

	if err := this.Datastore.Query(tags.ByOwner(owner)).All(&tags); err != nil {
		this.RenderStatusNotFound("Count not find tags for owner " + owner)
		return
	}

	this.render.JSON(http.StatusOK, tags)
}

func (this *tags) Create(tag models.Tag) {
	if err := this.Datastore.Create(&tag); err != nil {
		this.RenderError(err.Error())
		return
	}

	this.Headers.Add("Location", "/tags/"+tag.StringId())
	this.RenderCreated(tag)
}

func (this *tags) Retrieve(params martini.Params) {
	tag := &models.Tag{}
	tag.SetStringId(params["id"])

	if err := this.Datastore.Load(tag); err != nil {
		this.RenderStatusNotFound("Could not find tag for id " + tag.StringId())
		return
	}

	this.RenderOk(tag)
}

func (this *tags) Remove(params martini.Params) {
	tag := &models.Tag{}
	tag.SetStringId(params["id"])

	if err := this.Datastore.Delete(tag); err != nil {
		this.RenderError(err.Error())
		return
	}

	this.RenderOk(fmt.Sprintf("Tag %v successfully removed", tag.StringId()))
}
