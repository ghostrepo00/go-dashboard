package app

import (
	"github.com/ghostrepo00/go-dashboard/config"
	appinterface "github.com/ghostrepo00/go-dashboard/internal/pkg/app_interface"
	"github.com/ghostrepo00/go-dashboard/internal/pkg/model"
)

type webtagApp struct {
	repo appinterface.WebtagRepository
}

func NewWebtagApp(appConfig *config.AppConfig, repo appinterface.WebtagRepository) *webtagApp {
	return &webtagApp{repo}
}

func (r *webtagApp) GetWebtag() ([]model.Webtag, error) {
	return r.repo.GetWebtag()
}
