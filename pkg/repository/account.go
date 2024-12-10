package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"regulations-api/clerr"
	"regulations-api/models"
)

type AccountPostgres struct {
	db *sqlx.DB
}

func NewAccountPostgres(db *sqlx.DB) *AccountPostgres {
	return &AccountPostgres{db: db}
}

func (r *AccountPostgres) Get(email string) (*models.Account, error) {
	var account models.Account

	if err := r.db.Get(&account, `SELECT * FROM "Account" WHERE email=$1`, email); err != nil {
		logrus.Error(err.Error())
		return nil, clerr.ErrorServer
	}

	return &account, nil
}
