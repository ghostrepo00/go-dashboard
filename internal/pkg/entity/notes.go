package entity

import "github.com/google/uuid"

type Notes struct {
	Id      *uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Title   string
	Content string
}

func (Notes) TableName() string {
	return "notes"
}
