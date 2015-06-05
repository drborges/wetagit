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
	Datastore services.DatastoreProvider
}

func (this Tags) List(c *gin.Context) {
	owner := c.Query("owner")
	tags := models.Tags{}
	if err := this.Datastore().Query(tags.ByOwner(owner)).All(&tags); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("%v", err),
		})
		return
	}

	c.JSON(http.StatusOK, tags)
}

func (this Tags) Create(c *gin.Context) {
	tag := new(models.Tag)
	c.Bind(tag)

	err := this.Datastore().Create(tag)
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

	if err := this.Datastore().Load(tag); err != nil {
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

	err := this.Datastore().Delete(tag)
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
