package models

import (
	"time"

	"github.com/jinzhu/gorm/dialects/postgres"
)

type Category struct {
	ID          uint `gorm:"primary_key"`
	Menu_Id     uint
	Name        postgres.Jsonb
	Description postgres.Jsonb
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}
