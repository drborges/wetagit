package controllers

import (
	"github.com/drborges/datastore-model"
	"github.com/gin-gonic/gin"
	"github.com/drborges/wetagit/api/services"
	"appengine"
)

type Controller struct {
	GaeContext appengine.Context
	Datasource db.Datasource
}

func (this *Controller) Register(c *gin.Context) {
	this.GaeContext = services.Gae{c.Request}.NewContext()
	this.Datasource = db.NewDatastore(this.GaeContext)
}
