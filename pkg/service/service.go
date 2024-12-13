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

type Regulation interface {
	GetPrivate(email string) (*models.GetRegulationsOutput, error)
	UpdatePrivate(input models.UpdateRegulationInput, email string) error
}

type JWTToken interface {
	GenerateAccessToken(email string) (string, error)
	GenerateRefreshToken(email string) (string, error)
	ParseToken(tokenString string) (*models.JWTClaims, error)
}

type Service struct {
	Account
	Auth
	Regulation
	JWTToken
}

func NewService(repos *repository.Repository, config *models.ConfigService) *Service {
	return &Service{
		Account:    NewAccountService(repos.Account),
		Auth:       NewAuthService(repos.Auth),
		JWTToken:   NewJWTTokenService(config.Server),
		Regulation: NewRegulationService(repos.Regulation),
	}
}
