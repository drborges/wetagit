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

func (this Controller) RenderOk(data ...interface{}) {
	if len(data) == 1 {
		this.Renderer.JSON(http.StatusOK, data[0])
		return
	}

	this.Renderer.Status(http.StatusOK)
}

func (this Controller) RenderCreated(data ...interface{}) {
	if len(data) == 1 {
		this.Renderer.JSON(http.StatusCreated, data[0])
		return
	}

	this.Renderer.Status(http.StatusCreated)
}

func (this Controller) RenderStatusNotFound(message string) {
	this.Renderer.JSON(http.StatusNotFound, map[string]string{
		"message": fmt.Sprintf("%v", message),
	})
}

func (this Controller) RenderInternalServerError(message string) {
	this.Renderer.JSON(http.StatusNotFound, map[string]string{
		"message": fmt.Sprintf("%v", message),
	})
}
