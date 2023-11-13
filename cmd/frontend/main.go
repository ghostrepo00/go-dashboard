package main

import (
	"path/filepath"

	"github.com/caarlos0/env"
	"github.com/ghostrepo00/go-dashboard/config"
	"github.com/ghostrepo00/go-dashboard/internal/pkg/model"
	"github.com/ghostrepo00/go-dashboard/internal/web"
)

func main() {
	var envSettings model.EnvironmentSetting
	if err := env.Parse(&envSettings); err == nil {

		var configFileExtension string
		if envSettings.EnvironmentName == "dev" {
			configFileExtension = ".dev"
		}

		if appConfig, err := config.NewAppConfig(filepath.Join("./config/config.json", configFileExtension)); err == nil {
			appConfig.DbConnection = envSettings.ConnectionString
			web.Run(appConfig)
		} else {
			panic(err)
		}
	} else {
		panic(err)
	}
}
