package webtag

import (
	"github.com/ghostrepo00/go-dashboard/config"
)

type app struct {
	repo WebtagRepository
}

type WebtagApp interface {
	GetWebtag() error
}

func NewApp(appConfig *config.AppConfig, repo WebtagRepository) *app {
	return &app{repo}
}

func (r *app) GetWebtag() error {
	return r.repo.GetWebtag()
}
