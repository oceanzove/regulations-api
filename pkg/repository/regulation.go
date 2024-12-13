package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"regulations-api/models"
)

type RegulationPostgres struct {
	db *sqlx.DB
}

func NewRegulationPostgres(db *sqlx.DB) *RegulationPostgres {
	return &RegulationPostgres{db: db}
}

func (t *RegulationPostgres) GetPrivate(email string) (*models.GetRegulationsOutput, error) {
	var output models.GetRegulationsOutput

	err := t.db.Select(&output.Regulations, `SELECT id, title, content FROM "Regulation" WHERE  account_email = $1`, email)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return &output, nil
}

func (t *RegulationPostgres) UpdatePrivate(input models.UpdateRegulationInput, email string) error {
	var output models.GetRegulationsOutput

	err := t.db.Select(&output.Regulations, `UPDATE "Regulation" SET title = $1, content = $2 WHERE  id = $3 AND account_email = $4`, input.Title, input.Content, input.ID, email)
	if err != nil {
		logrus.Error(err.Error())
		return err
	}

	return nil
}

//func (t *OfferPostgres) Create(input *models.OfferCreateInput, email string) (*models.OfferCreateOutput, error) {
//	var output models.OfferCreateOutput
//	err := t.db.Get(&output, `INSERT INTO "Offer" (title, description, chapter) VALUES ($1, $2, $3) RETURNING *`, input.Title, input.Description, input.Chapter)
//	if err != nil {
//		logrus.Error(err.Error())
//		return nil, err
//	}
//
//	_, err = t.db.Exec(`INSERT INTO "AccountOffer" (offer, account, is_creator) VALUES ($1, $2, true)`, output.ID, email)
//	if err != nil {
//		logrus.Error(err.Error())
//		return nil, err
//	}
//
//	return &output, nil
//}

//func (t *OfferPostgres) Get(input *models.OfferGetInput) (*models.OfferGetActiveOutput, error) {
//	var output models.OfferGetActiveOutput
//
//	err := t.db.Select(&output.List, `SELECT id, title, description, status, chapter  FROM "Offer" WHERE status != 'создан'`)
//	if err != nil {
//		logrus.Error(err.Error())
//		return nil, err
//	}
//
//	return &output, nil
//}
