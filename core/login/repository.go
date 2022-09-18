package login

import (
	"errors"

	"github.com/EduardoNevesRamos/frelancer.git/model"
	"gorm.io/gorm"
)

type ILoginRepository interface {
	Login(login *model.Login, user *model.User) error
}

type LoginRepository struct {
	db *gorm.DB
}

func NewLoginRepository(dataBase *gorm.DB) ILoginRepository {
	return &LoginRepository{
		db: dataBase,
	}
}

func (r *LoginRepository) Login(login *model.Login, user *model.User) error {
	dbError := r.db.Where("email = ?", login.Email).First(&user).Error
	if dbError != nil {
		return errors.New("error: Invalid Email")
	}

	return nil
}
