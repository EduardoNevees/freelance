package login

import (
	"errors"

	"github.com/EduardoNevesRamos/frelancer.git/common"
	"github.com/EduardoNevesRamos/frelancer.git/model"
)

type ILoginService interface {
	Login(login *model.Login) (string, error)
}

type LoginService struct {
	repository ILoginRepository
}

func NewLoginService(Repository ILoginRepository) ILoginService {
	return &LoginService{Repository}
}

func (s *LoginService) Login(login *model.Login) (string, error) {
	var user model.User

	err := s.repository.Login(login, &user)
	if err != nil {
		return "", err
	}

	if common.SHA256Enconder(user.Password) != common.SHA256Enconder(login.Password) {
		return "", errors.New("invalid credentials")
	}

	token, err := common.NewJWTService().GenerateToken(uint(user.ID))
	if err != nil {
		return "", err
	}

	return token, nil
}
