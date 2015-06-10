package controllers

import (
	"github.com/drborges/datastore-model"
	"github.com/drborges/wetagit/api/services"
	"github.com/martini-contrib/render"
	"net/http"
	"net/url"
	"fmt"
)

type Controller struct {
	Datastore db.Datastore
	Query     url.Values
	Request   *http.Request
	Headers   http.Header
	render    render.Render
}

func (this *Controller) Register(render render.Render, req *http.Request) {
	this.Datastore = db.NewDatastore(services.Gae{req}.NewContext())
	this.Query = req.URL.Query()
	this.Headers = render.Header()
	this.render = render
	this.Request = req
}

func (this Controller) RenderOk(data ...interface{}) {
	if len(data) == 1 {
		this.render.JSON(http.StatusOK, data)
		return
	}

	this.render.Status(http.StatusOK)
}

func (this Controller) RenderCreated(data ...interface{}) {
	if len(data) == 1 {
		this.render.JSON(http.StatusCreated, data)
		return
	}

	this.render.Status(http.StatusCreated)
}

func (this Controller) RenderStatusNotFound(message string) {
	this.render.JSON(http.StatusNotFound, map[string]string{
		"message": fmt.Sprintf("%v", message),
	})
}

func (this Controller) RenderError(message string) {
	this.render.JSON(http.StatusNotFound, map[string]string{
		"message": fmt.Sprintf("%v", message),
	})
}
