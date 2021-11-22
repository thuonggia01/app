package models

import (
	"time"

	"github.com/jinzhu/gorm/dialects/postgres"
)

type Position struct {
	ID        uint `gorm:"primary_key"`
	Name      postgres.Jsonb
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
