package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"regulations-api/models"
)

type RegulationPostgres struct {
	db *sqlx.DB
}

func NewRegulationPostgres(db *sqlx.DB) *RegulationPostgres {
	return &RegulationPostgres{db: db}
}

func (t *RegulationPostgres) Create(accountId string, input *models.CreateRegulationInput) error {
	if input == nil {
		err := errors.New("input is nil")
		logrus.Error(err, err.Error())
		return err
	}

	if input.ID == "" {
		err := errors.New("input ID is required")
		logrus.Error(err, err.Error())
		return err
	}

	// 1. Получаем количество существующих регламентов для данного пользователя.
	var count int
	err := t.db.Get(&count, `SELECT COUNT(*) FROM "Regulation" WHERE account_id = $1`, accountId)
	if err != nil {
		logrus.Error("Error while counting regulations: ", err.Error())
		return err
	}

	// 2. Генерируем название, если не задано
	title := input.Title
	if title == "" {
		title = fmt.Sprintf("Регламент %d", count+1)
	}

	// 3. Используем ID и content из input
	id := input.ID
	content := input.Content

	// 4. Вставляем новый регламент в таблицу
	var newRegulationID string
	err = t.db.Get(&newRegulationID, `
        INSERT INTO "Regulation" (id, title, content, account_id)
        VALUES ($1, $2, $3, $4)
        RETURNING id
    `, id, title, content, accountId)
	if err != nil {
		logrus.Error("Error while inserting new regulation: ", err.Error())
		return err
	}

	return nil
}

func (t *RegulationPostgres) GetPrivate(accountId string) (*models.GetRegulationsOutput, error) {
	var output models.GetRegulationsOutput

	err := t.db.Select(&output.Regulations, `SELECT id, title, content FROM "Regulation" WHERE  account_id = $1`, accountId)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return &output, nil
}

func (t *RegulationPostgres) UpdatePrivate(input models.UpdateRegulationInput, email string) error {
	_, err := t.db.Exec(`UPDATE "Regulation" SET title = $1, content = $2 WHERE  id = $3 AND account_id = $4`, input.Title, input.Content, input.ID, email)
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
