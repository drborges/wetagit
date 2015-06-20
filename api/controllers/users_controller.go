package controllers

import (
	"github.com/drborges/wetagit/api/models"
	"github.com/go-martini/martini"
)

var Users = users{}

type users struct {
	Controller
}

func (this *users) Create(user models.User) {
	if err := this.CachedDatastore.Create(&user); err != nil {
		this.RenderInternalServerErrorMessage(err.Error())
		return
	}

	this.RenderCreatedResource(&user)
}

func (this *users) Retrieve(params martini.Params) {
	user := &models.User{Name: params["id"]}

	if err := this.CachedDatastore.Load(user); err != nil {
		this.RenderStatusNotFoundMessage("Could not find user for id %v", user.ResourceID())
		return
	}

	this.RenderData(user)
}

