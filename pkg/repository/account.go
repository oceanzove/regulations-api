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

func (r *AccountPostgres) GetByID(id string) (*models.Account, error) {
	var account models.Account
	if err := r.db.Get(&account, `SELECT * FROM "Account" WHERE id=$1`, id); err != nil {
		logrus.Error(err.Error())
		return nil, clerr.ErrorServer
	}
	return &account, nil
}

func NewAccountPostgres(db *sqlx.DB) *AccountPostgres {
	return &AccountPostgres{db: db}
}

func (r *AccountPostgres) Get(login string) (*models.Account, error) {
	var account models.Account

	if err := r.db.Get(&account, `SELECT * FROM "Account" WHERE login=$1`, login); err != nil {
		logrus.Error(err.Error())
		return nil, clerr.ErrorServer
	}

	return &account, nil
}
