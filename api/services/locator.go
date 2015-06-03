package services

import (
	"github.com/gin-gonic/gin"
	"github.com/drborges/datastore-model"
)

type Locator struct {
	context *gin.Context
}

func (this *Locator) Register(c *gin.Context) {
	this.context = c
}

func (this *Locator) Datastore() db.Datastore {
	context := Gae{this.context.Request}.NewContext()
	return db.Datastore{context}
}