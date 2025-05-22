package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"regulations-api/models"
)

type StepPostgres struct {
	db *sqlx.DB
}

func NewStepPostgres(db *sqlx.DB) *StepPostgres {
	return &StepPostgres{db: db}
}

func (t *StepPostgres) CreateSteps(input *models.CreateStepsInput) error {
	if len(input.Steps) == 0 {
		err := errors.New("input slice is empty")
		logrus.Error(err, err.Error())
		return err
	}

	for _, step := range input.Steps {
		if step.ProcessID == "" {
			err := errors.New("step.ProcessID is required")
			logrus.WithField("step", step).Error(err)
			return err
		}

		if step.Title == "" {
			err := errors.New("step.Name is required")
			logrus.WithField("step", step).Error(err)
			return err
		}

		_, err := t.db.Exec(`
			INSERT INTO "Step" (id, name, description, process_id, "order", responsible)
			VALUES ($1, $2, $3, $4, $5, $6)
		`, step.ID, step.Title, step.Description, step.ProcessID, step.Order, step.Responsible)

		if err != nil {
			logrus.WithField("step", step).Error("Error inserting step: ", err.Error())
			return err
		}
	}

	return nil
}

//func (t *ProcessPostgres) GetPrivate(email string) (*models.GetProcessesOutput, error) {
//	var output models.GetProcessesOutput
//
//	err := t.db.Select(&output.Processes, `SELECT id, title, description FROM "Process" WHERE  account_email = $1`, email)
//	if err != nil {
//		logrus.Error(err.Error())
//		return nil, err
//	}
//
//	return &output, nil
//}
//
//func (t *ProcessPostgres) UpdatePrivate(input models.UpdateProcessInput, email string) error {
//	_, err := t.db.Exec(`UPDATE "Process" SET title = $1, description = $2 WHERE  id = $3 AND account_email = $4`, input.Title, input.Description, input.ID, email)
//	if err != nil {
//		logrus.Error(err.Error())
//		return err
//	}
//
//	return nil
//}
