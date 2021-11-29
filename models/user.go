package models

import (
	"github.com/jinzhu/gorm"
)

type (
	User struct {
		gorm.Model
		Email       string
		Password    string
		Status      bool
		BranchID    int
		PartID      int
		PositionID  int
		ImageID     int
		FullName    string
		Phone       string
		Birthday    string
		Description string
		Facebook    string
	}

	UserLogin struct {
		FullName string
		Email    string
		Password string
		Url      string
		Base64   string
	}
)
