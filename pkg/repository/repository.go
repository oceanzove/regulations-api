package repository

import (
	"github.com/jmoiron/sqlx"
	"regulations-api/models"
)

type Sources struct {
	BusinessDB *sqlx.DB
}

type Auth interface {
}

type Account interface {
	Get(email string) (*models.Account, error)
}

type Regulation interface {
	GetPrivate(email string) (*models.GetRegulationsOutput, error)
	UpdatePrivate(input models.UpdateRegulationInput, email string) error
	Create(email string) (*models.CreateRegulationOutput, error)
}

type Process interface {
	GetPrivate(email string) (*models.GetProcessesOutput, error)
	UpdatePrivate(input models.UpdateProcessInput, email string) error
	Create(email string, input *models.CreateProcessInput) error
}

type Repository struct {
	Account
	Auth
	Regulation
	Process
}

func NewRepository(sources *Sources) *Repository {
	return &Repository{
		Account:    NewAccountPostgres(sources.BusinessDB),
		Auth:       NewAuthPostgres(sources.BusinessDB),
		Regulation: NewRegulationPostgres(sources.BusinessDB),
		Process:    NewProcessPostgres(sources.BusinessDB),
	}
}
