package usecase

import (
	"log"

	"github.com/jinzhu/gorm"

	"Hybrid-blog/helpers/database"
	"Hybrid-blog/helpers/utility"
	model "Hybrid-blog/models"
	repository "Hybrid-blog/repositories"
	repo "Hybrid-blog/repositories/repoImlp"
)

type UserCaseUser interface {
	IsLogin(email string, password string) (*model.UserLogin, bool)
	ForgotPassword(email string) bool
	ResetPassword(email string, password string) bool
}

type (
	userCaseUser struct {
		db             *gorm.DB
		userrepository repository.User
	}
)

func NewuserCaseTodo() UserCaseUser {
	return &userCaseUser{
		db:             database.DB(),
		userrepository: repo.NewUserRepo(),
	}
}

func (u *userCaseUser) IsLogin(email string, password string) (*model.UserLogin, bool) {
	user, ok := u.userrepository.IsLogin(email, u.db)
	var checkpass bool
	if user != nil {
		checkpass = utility.CheckPasswordHash(password, user.Password)
		log.Println(user.Password)
		if ok != nil {
			log.Println(ok)
			checkpass = false
			return nil, checkpass
		}
	}
	return user, checkpass
}

func (u *userCaseUser) ForgotPassword(email string) bool {
	err, ok := u.userrepository.ForgotPassword(email, u.db)
	if err != nil {
		log.Println(err)
	}
	if ok == true {
		log.Println(ok)
		return true
	}
	return false
}

func (u *userCaseUser) ResetPassword(email string, password string) bool {
	err, ok := u.userrepository.ResetPassword(email, password, u.db)
	if err != nil {
		log.Println(err)
	}
	if ok == true {
		return true
	}
	return false
}
