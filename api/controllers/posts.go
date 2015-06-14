package controllers

import (
	"github.com/drborges/wetagit/api/models"
	"github.com/go-martini/martini"
)

var Posts = posts{}

type posts struct {
	Controller
}

func (this *posts) List() {
	author := this.Query.Get("author")
	posts := models.Posts{}

	if err := this.Datastore.Query(posts.ByAuthor(author)).Results(&posts); err != nil {
		this.RenderOkMessage("Count not find posts for owner %v", author)
		return
	}

	this.RenderData(posts)
}

func (this *posts) Create(post models.Post) {
	if err := this.Datastore.Create(&post); err != nil {
		this.RenderInternalServerErrorMessage(err.Error())
		return
	}

	this.RenderCreatedData(&post)
}

func (this *posts) Retrieve(params martini.Params) {
	post := &models.Post{}
	post.SetUUID(params["id"])

	if err := this.Datastore.Load(post); err != nil {
		this.RenderStatusNotFoundMessage("Could not find post for id %v", post.UUID())
		return
	}

	this.RenderData(post)
}

func (this *posts) Remove(params martini.Params) {
	post := &models.Post{}
	post.SetUUID(params["id"])

	if err := this.Datastore.Delete(post); err != nil {
		this.RenderInternalServerErrorMessage(err.Error())
		return
	}

	this.RenderOkMessage("Post %v successfully removed", post.UUID())
}
