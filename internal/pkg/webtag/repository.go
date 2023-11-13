package webtag

import "gorm.io/gorm"

type repository struct {
	DbClient *gorm.DB
}

type WebtagRepository interface {
	GetWebtag() error
}

func NewRepository(dbClient *gorm.DB) *repository {
	return &repository{dbClient}
}

func (r *repository) GetWebtag() error {
	r.DbClient = nil
	return nil
}
