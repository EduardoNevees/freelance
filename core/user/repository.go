package user

import (
	"errors"
	"fmt"

	"github.com/EduardoNevesRamos/frelancer.git/model"
	"gorm.io/gorm"
)

type IRepository interface {
	GetUsers() ([]model.User, error)
	GetUserById(id int) (*model.User, error)
	CreateUser(user *model.User) error
	UpdateUser(user *model.User, id *int) error
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

func (r *UsersRepository) GetUsers() ([]model.User, error) {
	users := []model.User{}

	err := r.db.Find(&users).Error
	if err != nil {
		return nil, fmt.Errorf("can't find users %v", err)
	}

	return users, nil
}

func (r *UsersRepository) GetUserById(id int) (*model.User, error) {
	var user model.User

	err := r.db.First(&user, id).Error
	if err != nil {
		return &model.User{}, fmt.Errorf("cannot find products by id: %v", err)
	}

	return &user, nil
}

func (r *UsersRepository) CreateUser(user *model.User) error {
	err := r.db.Create(&user).Error
	if err != nil {
		return errors.New("can't create user")
	}

	return nil
}

func (r *UsersRepository) UpdateUser(user *model.User, id *int) error {
	err := r.db.Where(&id).Save(&user).Error
	if err != nil {
		return errors.New("can't update user")
	}
	return nil
}

func (r *UsersRepository) DeleteUser(parseId *int) error {
	user := &model.User{}

	err := r.db.Where(&parseId).Delete(&user).Error
	if err != nil {
		return errors.New("can't delete user")
	}

	return nil
}
