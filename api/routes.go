package api

import (
	"github.com/drborges/wetagit/api/controllers"
	"github.com/drborges/wetagit/api/models"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"net/http"
)

func Router() http.Handler {
	router := martini.Classic()
	Body := binding.Bind

	// TODO implements authentication middleware
	//router.Use(middlewares.Auth.Authenticate)
	//router.Use(middlewares.Auth.Authorize)
	router.Use(render.Renderer())
	router.Use(controllers.Tags.Register)

	router.Get("/tags", controllers.Tags.List)
	router.Post("/tags", Body(models.Tag{}), controllers.Tags.Create)
	router.Get("/tags/:id", controllers.Tags.Retrieve)
	router.Delete("/tags/:id", controllers.Tags.Remove)

	return router
}
