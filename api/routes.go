package api

import (
	"github.com/drborges/wetagit/api/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router() http.Handler {
	router := gin.New()

	// TODO implements authentication middleware
	//router.Use(middlewares.Auth.Authenticate)
	//router.Use(middlewares.Auth.Authorize)

	tags := new(controllers.Tags)

	router.Use(tags.Register)

	router.GET("/tags", tags.List)
	router.POST("/tags", tags.Create)
	router.GET("/tags/:id", tags.Retrieve)
	router.DELETE("/tags/:id", tags.Remove)
	return router
}
