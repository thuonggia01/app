package repository

import (
	"github.com/jinzhu/gorm"

	model "Hybrid-blog/models"
)

type User interface {
	ForgotPassword(email string, db *gorm.DB) (error, bool)
	ResetPassword(email string, pass string, db *gorm.DB) (error, bool)
	IsLogin(email string, db *gorm.DB) (*model.UserLogin, error)
}
