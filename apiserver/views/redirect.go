package views

import (
	"net/http"
	"xurl/apiserver/pkg/database"

	"github.com/gin-gonic/gin"
)

func RedirectView(c *gin.Context) {
	code := c.Param("code")
	link := database.Link{Code: code}

	db := database.GetDB()
	result := db.First(&link)
	if result.Error != nil {
		c.HTML(http.StatusOK, "notfound.tmpl", gin.H{
			"code": code,
		})
		return
	}
	c.Redirect(http.StatusMovedPermanently, link.Url)
}
