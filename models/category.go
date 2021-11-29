package models

import "github.com/jinzhu/gorm"

type (
	Category struct {
		gorm.Model
		MenuID      int
		Name        string
		Description string
		ParentID    int
	}

	SubCategory struct {
		gorm.Model
		MenuID      int
		Name        string
		Description string
		ParentID    int
	}
)
