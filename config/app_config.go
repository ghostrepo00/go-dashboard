package config

import (
	"encoding/json"
	"os"
)

type LogSettings struct {
	Path     string
	FileName string
}

type ApiSettings struct {
	Host          string
	BasePath      string
	FaviconFolder string
}

type AppConfig struct {
	Log          LogSettings
	Api          ApiSettings
	DbConnection string
}

func NewAppConfig(filePath string) (result *AppConfig, err error) {
	var (
		appConfig     *AppConfig
		appConfigFile *os.File
	)

	if appConfigFile, err = os.Open(filePath); err == nil {
		jsonParser := json.NewDecoder(appConfigFile)
		jsonParser.Decode(&appConfig)
	}

	defer appConfigFile.Close()

	return appConfig, err
}
