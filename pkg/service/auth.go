package service

import (
	"github.com/sirupsen/logrus"
	"regulations-api/models"
	"regulations-api/pkg/repository"
	"regulations-api/pkg/utils"
)

type AuthService struct {
	repo repository.Auth
}

func NewAuthService(repo repository.Auth) *AuthService {
	return &AuthService{repo: repo}
}

func (a AuthService) SignIn(input *models.SignInInput, accountPassword string) error {
	if err := utils.ComparePasswords(accountPassword, input.Password); err != nil {
		logrus.Warning(err.Error())
		return err
	}
	return nil
}
