package user

import (
	"errors"
	"fmt"

	"github.com/EduardoNevesRamos/frelancer.git/dto"
	"gorm.io/gorm"
)

type IRepository interface {
	GetUsers() ([]dto.User, error)
	GetUserById(id int) (*dto.User, error)
	CreateUser(user *dto.User) error
	UpdateUser(user *dto.User, id *int) error
	DeleteUser(parseId *int) error
}

type UsersRepository struct {
	db *gorm.DB
}

func NewUserRepository(dataBase *gorm.DB) IRepository {
	return &UsersRepository{
		db: dataBase,
	}
}

func (r *UsersRepository) GetUsers() ([]dto.User, error) {
	users := []dto.User{}

	err := r.db.Find(&users).Error
	if err != nil {
		return nil, fmt.Errorf("can't find users %v", err)
	}

	return users, nil
}

func (r *UsersRepository) GetUserById(id int) (*dto.User, error) {
	var user dto.User

	err := r.db.First(&user, id).Error
	if err != nil {
		return &dto.User{}, fmt.Errorf("cannot find products by id: %v", err)
	}

	return &user, nil
}

func (r *UsersRepository) CreateUser(user *dto.User) error {
	err := r.db.Create(&user).Error
	if err != nil {
		return errors.New("can't create user")
	}

	return nil
}

func (r *UsersRepository) UpdateUser(user *dto.User, id *int) error {
	err := r.db.Where(&id).Save(&user).Error
	if err != nil {
		return errors.New("can't update user")
	}
	return nil
}

func (r *UsersRepository) DeleteUser(parseId *int) error {
	user := &dto.User{}

	err := r.db.Where(&parseId).Delete(&user).Error
	if err != nil {
		return errors.New("can't delete user")
	}

	return nil
}
