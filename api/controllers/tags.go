package controllers

import (
	"fmt"
	"github.com/drborges/wetagit/api/models"
	"github.com/drborges/wetagit/api/services"
	"github.com/gin-gonic/gin"
	"net/http"
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

	if err := this.Datastore().Create(tag); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("%v", err),
		})
		return
	}

	c.Header("Location", "/tags/"+tag.StringId())
	c.JSON(http.StatusCreated, tag)
}

func (this Tags) Retrieve(c *gin.Context) {
	tag := new(models.Tag)
	tag.SetStringId(c.Params.ByName("id"))

	if err := this.Datastore().Load(tag); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("Could not find tag for id %v", tag.StringId()),
		})
		return
	}

	c.JSON(http.StatusOK, tag)
}

func (this Tags) Remove(c *gin.Context) {
	tag := new(models.Tag)
	tag.SetStringId(c.Params.ByName("id"))

	if err := this.Datastore().Delete(tag); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("%v", err),
		})
		return
	}

	c.JSON(http.StatusOK, "")
}
