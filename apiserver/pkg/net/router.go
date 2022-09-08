package net

import (
	"xurl/apiserver/api"
	"xurl/apiserver/views"

	"github.com/gin-gonic/gin"
)

func RouteInit(root string) *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob(root + "views/templates/*")

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
