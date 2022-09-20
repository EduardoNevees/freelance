package login

import (
	"errors"

	"github.com/EduardoNevesRamos/frelancer.git/dto"
	"gorm.io/gorm"
)

type ILoginRepository interface {
	Login(login *dto.Login, user *dto.User) error
}

type LoginRepository struct {
	db *gorm.DB
}

func NewLoginRepository(dataBase *gorm.DB) ILoginRepository {
	return &LoginRepository{
		db: dataBase,
	}
}

func (r *LoginRepository) Login(login *dto.Login, user *dto.User) error {
	dbError := r.db.Where("email = ?", login.Email).First(&user).Error
	if dbError != nil {
		return errors.New("error: Invalid Email")
	}

	return nil
}
