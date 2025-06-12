package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"regulations-api/models"
)

type OrganizationPostgres struct {
	db *sqlx.DB
}

func NewOrganizationPostgres(db *sqlx.DB) *OrganizationPostgres {
	return &OrganizationPostgres{db: db}
}

func (t *OrganizationPostgres) GetDepartments(accountId string) (*models.GetDepartmentOutput, error) {
	var output models.GetDepartmentOutput

	err := t.db.Select(&output.Departments, `
		SELECT d.id, d.name
		FROM "Department" d
		WHERE d.organization_id = (
			SELECT dep.organization_id
			FROM "EmployeeDepartment" ed
			JOIN "Department" dep ON dep.id = ed.department_id
			WHERE ed.employee_id = $1
			LIMIT 1
		)
	`, accountId)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return &output, nil
}

func (t *OrganizationPostgres) GetDepartmentByID(accountId string, departmentId string) (*models.Department, error) {
	var output models.Department

	err := t.db.Get(&output, `
	SELECT id, name
	FROM "Department" WHERE id = $1
`, departmentId)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return &output, nil
}

func (t *OrganizationPostgres) GetPositions(accountId string) (*models.GetPositionOutput, error) {
	var output models.GetPositionOutput

	err := t.db.Select(&output.Positions, `
		SELECT DISTINCT p.id, p.name
		FROM "Position" p
		JOIN "DepartmentPosition" dp ON dp.position_id = p.id
		JOIN "Department" d ON d.id = dp.department_id
		WHERE d.organization_id = (
			SELECT dep.organization_id
			FROM "EmployeeDepartment" ed
			JOIN "Department" dep ON dep.id = ed.department_id
			WHERE ed.employee_id = $1
			LIMIT 1
		)
	`, accountId)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return &output, nil
}

func (t *OrganizationPostgres) GetPositionsByDepartment(accountId string, departmentId string) (*models.GetPositionOutput, error) {
	var output models.GetPositionOutput

	err := t.db.Select(&output.Positions, `
		SELECT DISTINCT p.id, p.name
		FROM "Position" p
		JOIN "DepartmentPosition" dp ON dp.position_id = p.id
		JOIN "Department" d ON d.id = dp.department_id
		WHERE d.id = $2
		AND d.organization_id = (
			SELECT dep.organization_id
			FROM "EmployeeDepartment" ed
			JOIN "Department" dep ON dep.id = ed.department_id
			WHERE ed.employee_id = $1
			LIMIT 1
		)
	`, accountId, departmentId)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return &output, nil
}

func (t *OrganizationPostgres) CreateEmployee(input *models.CreateEmployeeInput) error {
	tx, err := t.db.Beginx()
	if err != nil {
		logrus.Error("failed to begin transaction:", err)
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			logrus.Error("panic recovered:", r)
		}
	}()

	// 1. Insert Employee
	_, err = tx.Exec(`
		INSERT INTO "Employee" (
			id, full_name, phone_number, birth_date, employment_date,
			residential_address, marital_status, email
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`,
		input.Employee.ID,
		input.Employee.FullName,
		input.Employee.PhoneNumber,
		input.Employee.BirthDate,
		input.Employee.EmploymentDate,
		input.Employee.ResidentialAddress,
		input.Employee.MaritalStatus,
		input.Employee.Email,
	)
	if err != nil {
		tx.Rollback()
		logrus.Error("failed to insert employee:", err)
		return err
	}

	// 2. Insert Account
	_, err = tx.Exec(`
		INSERT INTO "Account" (
			id, login, password, role
		) VALUES ($1, $2, $3, $4)
	`,
		input.Account.ID,
		input.Account.Login,
		input.Account.Password,
		input.Account.Role,
	)
	if err != nil {
		tx.Rollback()
		logrus.Error("failed to insert account:", err)
		return err
	}

	// 3. Insert EmployeeDepartment
	_, err = tx.Exec(`
		INSERT INTO "EmployeeDepartment" (employee_id, department_id)
		VALUES ($1, $2)
	`,
		input.Employee.ID,
		input.DepartmentID,
	)
	if err != nil {
		tx.Rollback()
		logrus.Error("failed to insert employee_department:", err)
		return err
	}

	// 4. Insert EmployeePosition
	_, err = tx.Exec(`
		INSERT INTO "EmployeePosition" (employee_id, position_id)
		VALUES ($1, $2)
	`,
		input.Employee.ID,
		input.PositionID,
	)
	if err != nil {
		tx.Rollback()
		logrus.Error("failed to insert employee_position:", err)
		return err
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		logrus.Error("failed to commit transaction:", err)
		return err
	}

	return nil
}

func (t *OrganizationPostgres) GetEmployees(accountId string) (*models.GetEmployeesOutput, error) {
	var output models.GetEmployeesOutput

	err := t.db.Select(&output.Employees, `
		SELECT e.*
		FROM "Employee" e
		JOIN "EmployeeDepartment" ed ON ed.employee_id = e.id
		JOIN "Department" d ON d.id = ed.department_id
		WHERE d.organization_id = (
    	SELECT dep.organization_id
    	FROM "EmployeeDepartment" ed2
    	JOIN "Department" dep ON dep.id = ed2.department_id
    	WHERE ed2.employee_id = $1
    	LIMIT 1
		);
	`, accountId)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return &output, nil
}
