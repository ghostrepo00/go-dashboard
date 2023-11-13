package entity

import "github.com/lib/pq"

type Webtag struct {
	Id    string `gorm:"primaryKey"`
	Url   string
	Title string
	Note  string
	Tags  pq.Int64Array `gorm:"type:integer[]"`
}

func (Webtag) TableName() string {
	return "web_tag"
}
