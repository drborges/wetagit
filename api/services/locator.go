package services

import (
	"github.com/gin-gonic/gin"
)

type Locator struct {
	context *gin.Context
}

func (this *Locator) Register(c *gin.Context) {
	this.context = c
}

func (this *Locator) Datastore() Datastore {
	context := Gae{this.context.Request}.NewContext()
	return Datastore{context}
}