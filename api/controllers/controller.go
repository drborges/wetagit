package controllers

import (
	"github.com/drborges/datastore-model"
	"github.com/drborges/wetagit/api/services"
	"github.com/martini-contrib/render"
	"net/http"
	"net/url"
	"fmt"
)

type datasource interface {
	Load(db.Entity) error
	LoadAll(...db.Entity) error
	Create(db.Entity) error
	CreateAll(...db.Entity) error
	Update(db.Entity) error
	UpdateAll(...db.Entity) error
	Delete(db.Entity) error
	DeleteAll(...db.Entity) error
	Query(q *db.Query) *db.Querier
}

type Controller struct {
	Datastore datasource
	Query     url.Values
	Request   *http.Request
	Headers   http.Header
	Renderer  render.Render
}

func (this *Controller) Register(render render.Render, req *http.Request) {
	this.Datastore = datasource(db.NewDatastore(services.Gae{req}.NewContext()))
	this.Query = req.URL.Query()
	this.Headers = render.Header()
	this.Renderer = render
	this.Request = req
}

func (this Controller) RenderOkMessage(message string, args ...interface{}) {
	if len(args) > 0 {
		this.Renderer.JSON(http.StatusOK, map[string]string {
			"message": fmt.Sprintf(message, args...),
		})
		return
	}

	this.Renderer.Status(http.StatusOK)
}

func (this Controller) RenderCreatedData(data db.Entity) {
	this.Renderer.Header().Add("Location", Resource{data}.Path())
	this.Renderer.JSON(http.StatusCreated, map[string]interface{} {
		"data": data,
	})
	return
}

func (this Controller) RenderData(data interface{}) {
	this.Renderer.JSON(http.StatusOK, map[string]interface{} {
		"data": data,
	})
}

func (this Controller) RenderStatusNotFoundMessage(message string, args ...interface{}) {
	this.Renderer.JSON(http.StatusNotFound, map[string]string{
		"message": fmt.Sprintf(message, args...),
	})
}

func (this Controller) RenderInternalServerErrorMessage(message string, args ...interface{}) {
	this.Renderer.JSON(http.StatusNotFound, map[string]string{
		"message": fmt.Sprintf(message, args...),
	})
}
