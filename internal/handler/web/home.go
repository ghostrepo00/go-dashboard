package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ConfigureHomeRouter(router *gin.Engine) {
	router.GET("", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
}
