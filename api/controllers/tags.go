package controllers

import (
	"github.com/drborges/wetagit/api/models"
	"github.com/drborges/wetagit/api/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/drborges/wetagit/api/services/db"
	"fmt"
)

type Tags struct {
	Locate *services.Locator
}

func (this Tags) Create(c *gin.Context) {
	tag := new(models.Tag)
	c.Bind(tag)

	err := this.Locate.Datastore().Create(tag)
	if err == db.ErrEntityExists {
		c.Header("Location", "/tags/" + tag.UUID())
		c.JSON(http.StatusNotModified, tag)
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("%v", err),
		})
		return
	}

	c.Header("Location", "/tags/" + tag.UUID())
	c.JSON(http.StatusCreated, gin.H{
		"message": "Tag successfully created",
	})
}

func (this Tags) Retrieve(c *gin.Context) {
	tag := new(models.Tag)
	tag.SetUUID(c.Params.ByName("id"))

	if err := this.Locate.Datastore().Load(tag); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("Could not find tag for id %v", tag.UUID()),
		})
		return
	}

	c.JSON(http.StatusOK, tag)
}

func (this Tags) Remove(c *gin.Context) {
	tag := new(models.Tag)
	tag.SetUUID(c.Params.ByName("id"))

	err := this.Locate.Datastore().Delete(tag)
	if err == db.ErrNoSuchEntity {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("%v", err),
		})
		return
	}

	c.JSON(http.StatusOK, tag)
}
