package user

import (
	"errors"
	"strconv"

	"github.com/EduardoNevesRamos/frelancer.git/dto"
)

type IService interface {
	GetUsers() ([]dto.User, error)
	GetUserById(id string) (*dto.User, error)
	CreateUser(user *dto.User) error
	UpdateUser(user *dto.User, id string) error
	DeleteUser(id string) error
}
type UserService struct {
	repository IRepository
}

func NewUserService(repository IRepository) IService {
	return &UserService{repository}
}

func (s *UserService) GetUsers() ([]dto.User, error) {
	userResponse, err := s.repository.GetUsers()
	if err != nil {
		return nil, err
	}

	return userResponse, nil
}

func (s *UserService) GetUserById(id string) (*dto.User, error) {
	parseId, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.New("ID has to be integer")
	}

	userResponse, err := s.repository.GetUserById(parseId)
	if err != nil {
		return nil, err
	}

	return userResponse, nil
}

func (s *UserService) CreateUser(user *dto.User) error {
	err := s.repository.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) UpdateUser(user *dto.User, id string) error {
	parseId, err := strconv.Atoi(id)
	if err != nil {
		return errors.New("ID has to be integer")
	}

	err = s.repository.UpdateUser(user, &parseId)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) DeleteUser(id string) error {
	parseId, err := strconv.Atoi(id)
	if err != nil {
		return errors.New("ID has to be integer")
	}

	err = s.repository.DeleteUser(&parseId)
	if err != nil {
		return err
	}

	return nil
}
