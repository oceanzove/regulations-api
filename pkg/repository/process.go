package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"regulations-api/models"
)

type ProcessPostgres struct {
	db *sqlx.DB
}

func NewProcessPostgres(db *sqlx.DB) *ProcessPostgres {
	return &ProcessPostgres{db: db}
}

func (t *ProcessPostgres) Create(email string) (*models.CreateProcessOutput, error) {
	// 1. Получаем количество существующих регламентов для данного пользователя.
	var count int
	err := t.db.Get(&count, `SELECT COUNT(*) FROM "Process" WHERE account_email = $1`, email)
	if err != nil {
		logrus.Error("Error while counting process: ", err.Error())
		return nil, err
	}

	// 2. Генерируем новое название регламента
	title := "Процесс " + fmt.Sprintf("%d", count+1)

	description := "Описание процесса " + fmt.Sprintf("%d", count+1)

	// 3. Вставляем новый регламент в таблицу
	var newProcessID string
	err = t.db.Get(&newProcessID, `INSERT INTO "Process" (title, description, account_email)
       VALUES ($1, $2 , $3) RETURNING id`, title, description, email)
	if err != nil {
		logrus.Error("Error while inserting new process: ", err.Error())
		return nil, err
	}

	// 4. Формируем и возвращаем результат
	output := &models.CreateProcessOutput{
		ID:          newProcessID,
		Title:       title,
		Description: description,
	}

	return output, nil
}

func (t *ProcessPostgres) GetPrivate(email string) (*models.GetProcessesOutput, error) {
	var output models.GetProcessesOutput

	err := t.db.Select(&output.Processes, `SELECT id, title, description FROM "Process" WHERE  account_email = $1`, email)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return &output, nil
}

func (t *ProcessPostgres) UpdatePrivate(input models.UpdateProcessInput, email string) error {
	_, err := t.db.Exec(`UPDATE "Process" SET title = $1, description = $2 WHERE  id = $3 AND account_email = $4`, input.Title, input.Description, input.ID, email)
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
