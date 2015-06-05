package api

import (
	"github.com/drborges/wetagit/api/controllers"
	"github.com/drborges/wetagit/api/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router() http.Handler {
	router := gin.New()
	locator := &services.Locator{}
	// Providers are used to initialize request bound services
	// to be used within controllers through lookup mechanism.
	//
	// This solution is not as ideal as DI, though it removes the
	// service initialization logic out of controllers.
	router.Use(locator.Register)

	// Middlewares are used to perform bits of sequential logic
	// before/after an incoming request, such as authentication,
	// logging...

	// TODO implements authentication middleware
	//router.Use(middlewares.Auth.Authenticate)
	//router.Use(middlewares.Auth.Authorize)

	tags := controllers.Tags{locator.Datastore}

	router.GET("/tags", tags.List)
	router.POST("/tags", tags.Create)
	router.GET("/tags/:id", tags.Retrieve)
	router.DELETE("/tags/:id", tags.Remove)
	return router
}
