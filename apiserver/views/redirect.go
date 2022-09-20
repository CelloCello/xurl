package views

import (
	"log"
	"net/http"
	"xurl/apiserver/pkg/database"

	"github.com/gin-gonic/gin"
)

func RedirectView(c *gin.Context) {
	code := c.Param("code")
	link := database.Link{}

	db := database.GetDB()
	result := db.First(&link, "Code = ?", code)
	if result.Error != nil {
		c.HTML(http.StatusOK, "notfound.tmpl", gin.H{
			"code": code,
		})
		return
	}
	log.Println(link)
	log.Printf("redirect to: %s\n", link.Url)
	c.Redirect(http.StatusMovedPermanently, link.Url)
}
