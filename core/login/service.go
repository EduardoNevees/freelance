package login

import (
	"errors"

	"github.com/EduardoNevesRamos/frelancer.git/common"
	"github.com/EduardoNevesRamos/frelancer.git/dto"
)

type ILoginService interface {
	Login(login *dto.Login) (string, error)
}

type LoginService struct {
	repository ILoginRepository
}

func NewLoginService(Repository ILoginRepository) ILoginService {
	return &LoginService{Repository}
}

func (s *LoginService) Login(login *dto.Login) (string, error) {
	var user dto.User

	err := s.repository.Login(login, &user)
	if err != nil {
		return "", err
	}

	if common.SHA256Enconder(user.Password) != common.SHA256Enconder(login.Password) {
		return "", errors.New("invalid credentials")
	}

	token, err := common.NewJWTService().GenerateToken(user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}
