package models

import (
	"time"

	"github.com/jinzhu/gorm/dialects/postgres"
)

type Branch struct {
	ID        uint `gorm:"primary_key"`
	Name      postgres.Jsonb
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
