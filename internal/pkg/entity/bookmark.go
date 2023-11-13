package entity

import (
	"time"

	"github.com/google/uuid"
)

type Bookmark struct {
	Id        *uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name      string
	URL       string
	GroupID   int64     `column:"group_id"`
	CreatedAt time.Time `column:"created_at"`
	Icon      []byte
}

func (Bookmark) TableName() string {
	return "bookmark"
}
