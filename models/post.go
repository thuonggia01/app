package models

import "github.com/jinzhu/gorm"

type Post struct {
	gorm.Model
	UserID      int
	Title       string
	Description string
	CategoryID  int
	ImageID     int
	Content     string
}
