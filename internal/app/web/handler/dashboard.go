package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ConfigureBookmarkRouter(router *gin.Engine) {
	router.GET("", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "Dashboard", "value2": "value 2"})
	})
}
