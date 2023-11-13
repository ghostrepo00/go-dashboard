package app

import (
	"io"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/ghostrepo00/go-dashboard/config"
	"github.com/ghostrepo00/go-dashboard/internal/app/web"
	appconstant "github.com/ghostrepo00/go-dashboard/internal/pkg/app_constant"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getLogFileName(appConfig *config.AppConfig) (result string) {
	currentDate := time.Now()
	return strings.Join([]string{appConfig.Log.Path, "/", currentDate.Format(appconstant.TimestampFormat), "_", appConfig.Log.FileName}, "")
}

func useSlog(appConfig *config.AppConfig) (logFile *os.File, err error) {
	logFile, err = os.OpenFile(getLogFileName(appConfig), os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)

	slogHandler := slog.NewJSONHandler(
		io.MultiWriter(os.Stdout, logFile),
		&slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelDebug,
		})
	logger := slog.New(slogHandler)
	slog.SetDefault(logger)
	return
}

func Run(appConfig *config.AppConfig) {
	if fileLog, err := useSlog(appConfig); err == nil {

		defer fileLog.Close()

		slog.Info("App started")
		// Configure Database
		if dbClient, err := gorm.Open(postgres.Open(appConfig.DbConnection), &gorm.Config{}); err == nil {
			defer func() {
				dbCon, _ := dbClient.DB()
				dbCon.Close()
			}()

			slog.Info("Database connected")

			router := gin.Default()
			router.Use(cors.Default())
			//api.ConfigureApiRouter(router, appConfig, dbClient)
			web.ConfigureWebRouter(router, appConfig, dbClient)
			router.Run(appConfig.Api.Host)

		} else {
			panic(err)
		}

	} else {
		panic(err)
	}
}