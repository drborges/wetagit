package controllers

import (
	"github.com/drborges/wetagit/api/models"
	"github.com/drborges/wetagit/api/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/drborges/datastore-model"
	"fmt"
)

type Tags struct {
	Locate *services.Locator
}

func (this Tags) Create(c *gin.Context) {
	tag := new(models.Tag)
	c.Bind(tag)

	err := this.Locate.Datastore().Create(tag)
	if err != nil && err != db.ErrEntityExists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("%v", err),
		})
		return
	}

	status := http.StatusCreated
	if err == db.ErrEntityExists {
		status = http.StatusNotModified
	}

	c.Header("Location", "/tags/" + tag.UUID())
	c.JSON(status, tag)
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
