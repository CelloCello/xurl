package net

import (
	"xurl/apiserver/api"
	"xurl/apiserver/views"

	"github.com/gin-gonic/gin"
)

func RouteInit(g *gin.Engine, root string) *gin.Engine {

	g.LoadHTMLGlob(root + "views/templates/*")

	// views
	g.GET("/hello", views.HelloView)
	g.GET("/:code", views.RedirectView)

	// APIs
	g.GET("api/ping", api.TestPing)
	g.GET("api/links", api.GetLinks)
	g.POST("api/links", api.CreateLink)
	g.DELETE("api/links/:id", api.DeleteLink)

	return g
}
