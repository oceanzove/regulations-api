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

type Repository struct {
	Account
	Auth
}

func NewRepository(sources *Sources) *Repository {
	return &Repository{
		Account: NewAccountPostgres(sources.BusinessDB),
		Auth:    NewAuthPostgres(sources.BusinessDB),
	}
}
