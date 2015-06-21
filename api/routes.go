package api

import (
	"github.com/drborges/wetagit/api/controllers"
	"github.com/drborges/wetagit/api/injectables"
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
	router.Use(injectables.GaeContextProvider)
	router.Use(injectables.DatastoreProvider)
	router.Use(injectables.CachedDatastoreProvider)
	router.Use(injectables.CurrentUserProvider)

	router.Use(controllers.Tags.Register)
	{
		router.Get("/tags", controllers.Tags.List)
		router.Get("/tags/:id", controllers.Tags.Retrieve)
		router.Delete("/tags/:id", controllers.Tags.Remove)
		router.Post("/tags", Body(models.Tag{}), controllers.Tags.Create)
	}

	router.Use(controllers.Posts.Register)
	{
		router.Get("/posts", controllers.Posts.List)
		router.Get("/posts/:id", controllers.Posts.Retrieve)
		router.Delete("/posts/:id", controllers.Posts.Remove)
		router.Post("/posts", Body(models.Post{}), controllers.Posts.Create)
	}

	router.Use(controllers.Feeds.Register)
	{
		router.Get("/feeds/:id", controllers.Feeds.Retrieve)
		router.Post("/feeds/:id/subscribe", controllers.Feeds.Subscribe)
		router.Post("/feeds/:id/unsubscribe", controllers.Feeds.Unsubscribe)
	}

	router.Use(controllers.Users.Register)
	{
		router.Get("/users/:id", controllers.Users.Retrieve)
		router.Post("/users", Body(models.User{}), controllers.Users.Create)
	}

	return router
}
