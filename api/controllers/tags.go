package controllers

import (
	"github.com/drborges/wetagit/api/models"
	"github.com/drborges/wetagit/api/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Tags struct {
	Locate *services.Locator
}

func (this Tags) Create(c *gin.Context) {
	tag := new(models.Tag)
	c.Bind(tag)

	if ok := this.Locate.Datastore().Save(tag); !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create tag",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Tag successfully created",
	})
}

