package web

import (
	"net/http"

	appinterface "github.com/ghostrepo00/go-dashboard/internal/pkg/app_interface"
	"github.com/gin-gonic/gin"
)

func ConfigureHomeRouter(router *gin.Engine, webtagApp appinterface.WebtagApp) {
	router.GET("", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
}
