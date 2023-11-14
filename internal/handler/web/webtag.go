package web

import (
	"net/http"

	appinterface "github.com/ghostrepo00/go-dashboard/internal/pkg/app_interface"
	"github.com/gin-gonic/gin"
)

func ConfigureWebtagRouter(router *gin.Engine, webtagApp appinterface.WebtagApp) {
	webtagGroup := router.Group("/webtag")
	{
		webtagGroup.GET("", func(c *gin.Context) {
			c.HTML(http.StatusOK, "webtag.body", gin.H{})
		})

		webtagGroup.GET("test", func(c *gin.Context) {
			items, _ := webtagApp.GetWebtag()
			c.HTML(http.StatusOK, "webtag.result", gin.H{"items": items})
		})
	}
}
