package controllers

import (
	"github.com/drborges/wetagit/api/models"
)

var Subscriptions = subscriptions{}

type subscriptions struct {
	Controller
}

func (this *subscriptions) Create(user models.User, subscriptions models.Subscriptions) {
	subscriptions.BelongTo(user)

	if err := this.Datastore.CreateAll(subscriptions); err != nil {
		this.RenderInternalServerErrorMessage(err.Error())
		return
	}

	this.RenderData(&subscriptions)
}
