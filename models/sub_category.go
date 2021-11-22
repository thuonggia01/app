package models

import (
	"time"

	"github.com/jinzhu/gorm/dialects/postgres"
)

type Sub_Category struct {
	ID          uint `gorm:"primary_key"`
	Menu_Id     uint
	Parent_id   uint
	Name        postgres.Jsonb
	Description postgres.Jsonb
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}
