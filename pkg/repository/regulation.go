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

	var departmentID string
	err := t.db.Get(&departmentID, `
		SELECT ed.department_id
		FROM "EmployeeDepartment" ed
		WHERE ed.employee_id = $1
		LIMIT 1
	`, accountId)
	if err != nil {
		logrus.Error("Failed to find department for employee/account: ", err)
		return err
	}

	// 1. Получаем количество существующих регламентов для данного пользователя.
	var count int
	err = t.db.Get(&count, `SELECT COUNT(*) FROM "Regulation" WHERE department_id = $1`, departmentID)
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
        INSERT INTO "Regulation" (id, title, content, department_id)
        VALUES ($1, $2, $3, $4)
        RETURNING id
    `, id, title, content, departmentID)
	if err != nil {
		logrus.Error("Error while inserting new regulation: ", err.Error())
		return err
	}

	return nil
}

func (t *RegulationPostgres) CreateSection(accountId string, input *models.CreateSectionInput) error {
	if input == nil {
		err := errors.New("input is nil")
		logrus.Error(err)
		return err
	}
	if input.ID == "" {
		err := errors.New("input ID is required")
		logrus.Error(err)
		return err
	}

	// Получить количество секций (можно убрать, если не нужен автотайтл)
	var count int
	err := t.db.Get(&count, `SELECT COUNT(*) FROM "Section"`)
	if err != nil {
		logrus.Error("Error while counting sections: ", err)
		return err
	}

	// Заголовок по умолчанию
	title := input.Title
	if title == "" {
		title = fmt.Sprintf("Секция %d", count+1)
	}

	// Получить department_id сотрудника
	var departmentID string
	err = t.db.Get(&departmentID, `
		SELECT ed.department_id
		FROM "EmployeeDepartment" ed
		WHERE ed.employee_id = $1
		LIMIT 1
	`, accountId)
	if err != nil {
		logrus.Error("Failed to find department for employee/account: ", err)
		return err
	}

	// Вставка новой секции
	var newSectionID string
	err = t.db.Get(&newSectionID, `
        INSERT INTO "Section" (id, title, content, department_id)
        VALUES ($1, $2, $3, $4)
        RETURNING id
    `, input.ID, title, input.Content, departmentID)
	if err != nil {
		logrus.Error("Error while inserting new section: ", err)
		return err
	}

	return nil
}

func (t *RegulationPostgres) GetByID(accountID string, regulationID string) (*models.Regulation, error) {

	var departmentID string
	err := t.db.Get(&departmentID, `
	SELECT ed.department_id
	FROM "EmployeeDepartment" ed
	WHERE ed.employee_id = $1
	LIMIT 1
	`, accountID)

	var regulation models.Regulation

	err = t.db.Get(&regulation, `SELECT * FROM "Regulation" WHERE id=$1 AND department_id=$2`, regulationID, departmentID)
	if err != nil {
		return nil, err
	}
	return &regulation, nil
}

func (t *RegulationPostgres) GetPrivate(accountId string) (*models.GetRegulationsOutput, error) {
	var output models.GetRegulationsOutput

	var departmentID string
	err := t.db.Get(&departmentID, `
	SELECT ed.department_id
	FROM "EmployeeDepartment" ed
	WHERE ed.employee_id = $1
	LIMIT 1
	`, accountId)

	err = t.db.Select(&output.Regulations, `SELECT id, title, content FROM "Regulation" WHERE  department_id = $1`, departmentID)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return &output, nil
}

func (t *RegulationPostgres) GetSections(accountID string) (*models.GetSectionsOutput, error) {
	var output models.GetSectionsOutput

	err := t.db.Select(&output.Sections, `SELECT id, title, content FROM "Section"`)
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
