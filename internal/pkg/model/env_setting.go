package model

type EnvironmentSetting struct {
	EnvironmentName  string `env:"ENVIRONMENT_NAME"`
	ConnectionString string `env:"PG_CONNECTION" envDefault:"none"`
}
