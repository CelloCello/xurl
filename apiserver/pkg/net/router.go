package net

import (
	"short_url/apiserver/api"
	"short_url/apiserver/views"

	"github.com/gin-gonic/gin"
)

func RouteInit() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("views/templates/*")

	// views
	r.GET("/hello", views.HelloView)
	r.GET("/:code", views.RedirectView)

	// APIs
	r.GET("api/ping", api.TestPing)
	r.GET("api/links", api.GetLinks)
	r.POST("api/links", api.CreateLink)
	r.DELETE("api/links/:id", api.DeleteLink)

	return r
}
