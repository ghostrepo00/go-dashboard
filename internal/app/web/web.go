package web

import (
	"github.com/ghostrepo00/go-dashboard/config"
	"github.com/ghostrepo00/go-dashboard/internal/app/web/handler"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ConfigureWebRouter(router *gin.Engine, appConfig *config.AppConfig, dbClient *gorm.DB) {
	router.LoadHTMLGlob("internal/app/web/template/*")
	router.Static("/asset", "internal/app/web/asset")

	handler.ConfigureBookmarkRouter(router)
}
