package repoimlp

import (
	"log"

	"github.com/jinzhu/gorm"

	model "Hybrid-blog/models"
	repo "Hybrid-blog/repositories"
)

type userRepo struct{}

func NewUserRepo() repo.User {
	return &userRepo{}
}

func (u *userRepo) SelectUsers(db *gorm.DB) (*model.User, error) {
	var user *model.User
	result := db.Raw("SELECT * FROM public.user").Scan(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (u *userRepo) IsLogin(email string, db *gorm.DB) (*model.UserLogin, error) {
	var list []*model.UserLogin
	var result *gorm.DB
	var user *model.UserLogin
	var err error
	result = db.Table("public.users").
		Select("users.email, users.password ,images.base64 ,users.image_id , images.id").Joins("INNER JOIN public.images ON images.id = users.image_id").
		Where("users.email = ?", email)
	result.Scan(&list)

	if len(list) > 0 {
		user = list[0]
		err = nil
	} else {
		result = db.Table("public.users").
			Select("users.id, users.email, users.password").
			Where("users.email = ?", email)
		result.Scan(&list)
		if len(list) > 0 {
			user = list[0]
			err = nil
		}
	}

	if result.Error != nil {
		user = nil
		err = result.Error
	}
	return user, err
}

func (u *userRepo) ForgotPassword(email string, db *gorm.DB) (error, bool) {
	var err error
	var checkFounds bool = false
	var list []*model.User

	var result *gorm.DB
	result = db.Table("public.users").Select("email").Where("email = ?", email).Limit(1).Scan(&list)

	if len(list) > 0 {
		checkFounds = true
		err = nil
	} else if result.Error != nil {
		err = result.Error
	}
	return err, checkFounds
}

//UPDATE public.users SET password= 'sdasdasdasdas' WHERE email = 'sonkehb096@gmail.com'
func (u *userRepo) ResetPassword(email string, pass string, db *gorm.DB) (error, bool) {
	var checkBool = true
	rback := db.Begin()
	err := db.Table("public.users").Where("email = ?", email).Update("password", pass)
	if err.Error != nil {
		checkBool = false
		log.Println("RoolBack")
		rback.Rollback()
		return err.Error, checkBool
	}
	return nil, checkBool
}
