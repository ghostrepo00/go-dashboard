package appinterface

import "github.com/ghostrepo00/go-dashboard/internal/pkg/model"

type WebtagApp interface {
	GetWebtag() ([]model.Webtag, error)
	// GetWebtagById(id string) (model.Webtag, error)
	// CreateWebtag(model *model.Webtag) error
	// UpdateWebtag(model *model.Webtag) error
	// DeleteWebtag(id string) error
}

type WebtagRepository interface {
	GetWebtag() (result []model.Webtag, err error)
	GetWebtagById(id string) (result model.Webtag, err error)
	CreateWebtag(model *model.Webtag) (err error)
	UpdateWebtag(model *model.Webtag) (err error)
	DeleteWebtag(id string) (err error)
}
