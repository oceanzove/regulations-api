package service

import (
	"regulations-api/models"
	"regulations-api/pkg/repository"
)

type Account interface {
	Get(email string) (*models.Account, error)
}

type Auth interface {
	SignIn(input *models.SignInInput, accountPassword string) error
}

type JWTToken interface {
	GenerateAccessToken(email string) (string, error)
	GenerateRefreshToken(email string) (string, error)
	ParseToken(tokenString string) (*models.JWTClaims, error)
}

type Service struct {
	Account
	Auth
	JWTToken
}

func NewService(repos *repository.Repository, config *models.ConfigService) *Service {
	return &Service{
		Account:  NewAccountService(repos),
		Auth:     NewAuthService(repos),
		JWTToken: NewJWTTokenService(config.Server),
	}
}
