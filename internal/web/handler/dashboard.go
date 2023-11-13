package handler

import (
	"net/http"

	"github.com/ghostrepo00/go-dashboard/config"
	"github.com/ghostrepo00/go-dashboard/internal/pkg/webtag"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ConfigureBookmarkRouter(router *gin.Engine, appConfig *config.AppConfig, dbClient *gorm.DB) {
	webtagApp := webtag.NewApp(appConfig, webtag.NewRepository(dbClient))

	router.GET("", func(c *gin.Context) {
		webtagApp.GetWebtag()
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "Dashboardc", "value2": "value 2"})
	})
}
