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
	Create(email string) (*models.CreateRegulationOutput, error)
}

type Process interface {
	GetPrivate(email string) (*models.GetProcessesOutput, error)
	UpdatePrivate(input *models.UpdateProcessInput, email string) error
	Create(email string, input *models.CreateProcessInput) error
}

type Step interface {
	Create(input *models.CreateStepsInput) error
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
	Process
	JWTToken
	Step
}

func NewService(repos *repository.Repository, config *models.ConfigService) *Service {
	return &Service{
		Account:    NewAccountService(repos.Account),
		Auth:       NewAuthService(repos.Auth),
		JWTToken:   NewJWTTokenService(config.Server),
		Regulation: NewRegulationService(repos.Regulation),
		Process:    NewProcessService(repos.Process),
		Step:       NewStepService(repos.Step),
	}
}
