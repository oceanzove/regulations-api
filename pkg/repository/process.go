package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"regulations-api/models"
)

type ProcessPostgres struct {
	db *sqlx.DB
}

func NewProcessPostgres(db *sqlx.DB) *ProcessPostgres {
	return &ProcessPostgres{db: db}
}

func (t *ProcessPostgres) Create(accountId string, input *models.CreateProcessInput) error {
	if input == nil {
		err := errors.New("input input is nil")
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
	err := t.db.Get(&count, `SELECT COUNT(*) FROM "Process" WHERE account_id = $1`, accountId)
	if err != nil {
		logrus.Error("Error while counting input: ", err.Error())
		return err
	}

	// 2. Генерируем название, если не задано
	title := input.Title
	if title == "" {
		title = fmt.Sprintf("Процесс %d", count+1)
	}

	// 3. Используем ID и описание из input
	id := input.ID
	description := input.Description

	// 4. Вставляем новый процесс в таблицу
	var newProcessID string
	err = t.db.Get(&newProcessID, `
		INSERT INTO "Process" (id, title, description, account_id)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`, id, title, description, accountId)
	if err != nil {
		logrus.Error("Error while inserting new input: ", err.Error())
		return err
	}

	return nil
}

func (t *ProcessPostgres) GetPrivate(accountId string) (*models.GetProcessesOutput, error) {
	var output models.GetProcessesOutput

	err := t.db.Select(&output.Processes, `SELECT id, title, description FROM "Process" WHERE  account_id = $1`, accountId)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return &output, nil
}

func (t *ProcessPostgres) GetByID(accountID, processID string) (*models.Process, error) {
	var process models.Process
	err := t.db.Get(&process, `SELECT * FROM "Process" WHERE id=$1 AND account_id=$2`, processID, accountID)
	if err != nil {
		return nil, err
	}
	return &process, nil
}

func (t *ProcessPostgres) UpdatePrivate(input *models.UpdateProcessInput, accountId string) error {
	_, err := t.db.Exec(`UPDATE "Process" SET title = $1, description = $2 WHERE  id = $3 AND account_id = $4`, input.Title, input.Description, input.ID, accountId)
	if err != nil {
		logrus.Error(err.Error())
		return err
	}

	return nil
}

func (t *ProcessPostgres) LinkRegulationToProcess(processID, regulationID string) error {
	_, err := t.db.Exec(`
		INSERT INTO "ProcessRegulation" (process_id, regulation_id)
		VALUES ($1, $2)
		ON CONFLICT DO NOTHING
	`, processID, regulationID)
	return err
}

func (t *ProcessPostgres) CreateStep(input *models.Step) error {
	if input == nil {
		return errors.New("input is nil")
	}
	_, err := t.db.Exec(`
		INSERT INTO "Step" (id, name, description, process_id, responsible, "order")
		VALUES ($1, $2, $3, $4, $5, $6)
	`, input.ID, input.Title, input.Description, input.ProcessID, input.Responsible, input.Order)
	return err
}

func (t *ProcessPostgres) GetStepsByProcess(processID string) ([]*models.Step, error) {
	var steps []*models.Step

	err := t.db.Select(
		&steps,
		`SELECT id, name, description, "order", process_id, responsible
		 FROM "Step"
		 WHERE process_id = $1`,
		processID,
	)

	if err != nil {
		logrus.Error("Error while getting steps by process: ", err)
		return nil, err
	}

	return steps, nil
}

func (t *ProcessPostgres) GetRegulationsByProcess(processID string) ([]*models.Regulation, error) {
	var regulations []*models.Regulation

	err := t.db.Select(
		&regulations,
		`SELECT r.id, r.title, r.content
		 FROM "Regulation" r
		 JOIN "ProcessRegulation" pr ON pr.regulation_id = r.id
		 WHERE pr.process_id = $1`,
		processID,
	)

	if err != nil {
		logrus.Error("Error while getting regulations by process: ", err)
		return nil, err
	}

	return regulations, nil
}
