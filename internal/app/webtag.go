package app

import (
	"github.com/ghostrepo00/go-dashboard/config"
	appinterface "github.com/ghostrepo00/go-dashboard/internal/pkg/app_interface"
)

type webtagApp struct {
	repo appinterface.WebtagRepository
}

func NewWebtagApp(appConfig *config.AppConfig, repo appinterface.WebtagRepository) *webtagApp {
	return &webtagApp{repo}
}

func (r *webtagApp) GetWebtag() error {
	return r.repo.GetWebtag()
}

func (r *webtagApp) GetWebtagNonInterface() error {
	return r.repo.GetWebtag()
}
