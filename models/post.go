package models

import (
	"time"

	"github.com/jinzhu/gorm/dialects/postgres"
)

type Post struct {
	ID              uint `gorm:"primary_key"`
	Title           postgres.Jsonb
	Description     postgres.Jsonb
	User_Id         uint
	Category_Id     uint
	Image_Id        uint
	Sub_Category_Id uint
	Content         string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time
}
