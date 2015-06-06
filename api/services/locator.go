package services

import (
	"github.com/drborges/datastore-model"
	"github.com/gin-gonic/gin"
)

type DatastoreProvider func() db.Datastore

type Locator struct {
	context *gin.Context
}

func (this *Locator) Register(c *gin.Context) {
	this.context = c
}

func (this *Locator) Datastore() db.Datastore {
	context := Gae{this.context.Request}.NewContext()
	return db.NewDatastore(context)
}
