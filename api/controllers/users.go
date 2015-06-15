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
	if err := this.Datastore.Create(&user); err != nil {
		this.RenderInternalServerErrorMessage(err.Error())
		return
	}

	this.RenderCreatedData(&user)
}

func (this *users) Retrieve(params martini.Params) {
	user := &models.User{}
	user.SetID(params["id"])

	if err := this.Datastore.Load(user); err != nil {
		this.RenderStatusNotFoundMessage("Could not find user for id %v", user.ID())
		return
	}

	this.RenderData(user)
}

