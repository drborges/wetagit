package controllers

import (
	"fmt"
	"github.com/drborges/appx"
	"github.com/martini-contrib/render"
	"net/http"
	"net/url"
)

type Controller struct {
	Datastore       *appx.Datastore
	CachedDatastore *appx.CachedDatastore
	Query           url.Values
	Request         *http.Request
	Headers         http.Header
	Renderer        render.Render
}

func (this *Controller) Register(render render.Render, req *http.Request, ds *appx.Datastore, cds *appx.CachedDatastore) {
	this.Datastore = ds
	this.CachedDatastore = cds
	this.Query = req.URL.Query()
	this.Headers = render.Header()
	this.Renderer = render
	this.Request = req
}

func (this Controller) RenderOkMessage(message string, args ...interface{}) {
	if len(args) > 0 {
		this.Renderer.JSON(http.StatusOK, map[string]string{
			"message": fmt.Sprintf(message, args...),
		})
		return
	}

	this.Renderer.Status(http.StatusOK)
}

func (this Controller) RenderCreatedResource(resource appx.Resource) {
	this.Renderer.Header().Add("Location", Resource{resource}.Path())
	this.Renderer.JSON(http.StatusCreated, map[string]interface{}{
		"data": resource,
	})
	return
}

func (this Controller) RenderData(data interface{}) {
	this.Renderer.JSON(http.StatusOK, map[string]interface{}{
		"data": data,
	})
}

func (this Controller) RenderStatusNotFoundMessage(message string, args ...interface{}) {
	this.Renderer.JSON(http.StatusNotFound, map[string]string{
		"message": fmt.Sprintf(message, args...),
	})
}

func (this Controller) RenderInternalServerErrorMessage(message string, args ...interface{}) {
	this.Renderer.JSON(http.StatusInternalServerError, map[string]string{
		"message": fmt.Sprintf(message, args...),
	})
}
