package user

import (
	"errors"

	"github.com/EduardoNevesRamos/frelancer.git/model"
	"gorm.io/gorm"
)

type IRepository interface {
	GetUsers() ([]*model.User, error)
	GetUserById(id int) (*model.User, error)
	CreateUser() (*model.User, error)
	UpdateUser() (*model.User, error)
	DeleteUser(parseId int) error
}

type UsersRepository struct {
	db *gorm.DB
}

func NewUserRepository(dataBase *gorm.DB) IRepository {
	return &UsersRepository{
		db: dataBase,
	}
}

func (r *UsersRepository) GetUsers() ([]*model.User, error) {
	users := []*model.User{}

	err := r.db.Find(&users).Error
	if err != nil {
		return nil, errors.New("can't find users at database")
	}

	return users, nil
}

func (r *UsersRepository) GetUserById(id int) (*model.User, error) {
	user := model.User{}

	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, errors.New("can't find user by id at database")
	}

	return &user, nil
}

func (r *UsersRepository) CreateUser() (*model.User, error) {
	user := &model.User{}

	err := r.db.Create(user).Error
	if err != nil {
		return nil, errors.New("can't create user")
	}

	return user, nil
}

func (r *UsersRepository) UpdateUser() (*model.User, error) {
	user := &model.User{}

	err := r.db.Save(user).Error
	if err != nil {
		return nil, errors.New("can't update user")
	}

	return user, nil
}

func (r *UsersRepository) DeleteUser(parseId int) error {
	user := &model.User{}

	err := r.db.Delete(user, parseId).Error
	if err != nil {
		return errors.New("can't delete user")
	}

	return nil
}
