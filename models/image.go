package models

import (
	"time"

	"github.com/jinzhu/gorm/dialects/postgres"
)

type Image struct {
	ID        uint `gorm:"primary_key"`
	Url       postgres.Jsonb
	Base64    postgres.Jsonb
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
