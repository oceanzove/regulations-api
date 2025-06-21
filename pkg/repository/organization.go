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

func (t *OrganizationPostgres) GetEmployeeDepartment(accountId string) (*models.GetEmployeeDepartmentOutput, error) {
	var output models.GetEmployeeDepartmentOutput

	err := t.db.Select(&output.EmployeeDepartment, `
		SELECT DISTINCT ed.employee_id, ed.department_id
		FROM "EmployeeDepartment" ed
		JOIN "Department" d ON ed.department_id = d.id
		WHERE d.organization_id = (
    	SELECT d2.organization_id
    	FROM "EmployeeDepartment" ed2
    	JOIN "Department" d2 ON ed2.department_id = d2.id
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

func (t *OrganizationPostgres) GetEmployeePosition(accountId string) (*models.GetEmployeePositionOutput, error) {
	var output models.GetEmployeePositionOutput

	err := t.db.Select(&output.EmployeePosition, `
	SELECT ep."employee_id", ep."position_id"
	FROM "EmployeePosition" ep
	JOIN "Employee" e ON ep."employee_id" = e."id"
	JOIN "EmployeeDepartment" ed ON e."id" = ed."employee_id"
	JOIN "Department" d ON ed."department_id" = d."id"
	JOIN "Organization" o ON d."organization_id" = o."id"
	WHERE o."id" = (
		SELECT d."organization_id"
		FROM "Department" d
		JOIN "EmployeeDepartment" ed ON d."id" = ed."department_id"
		WHERE ed."employee_id" = $1
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

func (t *OrganizationPostgres) CreatePosition(input *models.CreatePositionInput) error {
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

	// 1. Insert Position
	_, err = tx.Exec(`
		INSERT INTO "Position" (
			id, name
		) VALUES ($1, $2)
	`,
		input.ID,
		input.Name,
	)
	if err != nil {
		tx.Rollback()
		logrus.Error("failed to insert Position:", err)
		return err
	}

	// 2. Insert DepartmentPosition
	_, err = tx.Exec(`
		INSERT INTO "DepartmentPosition" (
			department_id, position_id
		) VALUES ($1, $2)
	`,
		input.DepartmentID,
		input.ID,
	)
	if err != nil {
		tx.Rollback()
		logrus.Error("failed to insert DepartmentPosition:", err)
		return err
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		logrus.Error("failed to commit transaction:", err)
		return err
	}

	return nil
}

func (t *OrganizationPostgres) CreateDepartment(accountID string, input *models.CreateDepartmentInput) error {
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

	var organization models.Organization

	err = t.db.Get(&organization, `
		SELECT o.*
		FROM "Organization" o
	JOIN "Department" d ON d.organization_id = o.id
	JOIN "EmployeeDepartment" ed ON ed.department_id = d.id
	JOIN "Employee" e ON e.id = ed.employee_id
	WHERE e.id = $1
	LIMIT 1;
	`, accountID)

	// 1. Insert Department
	_, err = tx.Exec(`
		INSERT INTO "Department" (
			id, name, organization_id
		) VALUES ($1, $2, $3)
	`,
		input.ID,
		input.Name,
		organization.ID,
	)
	if err != nil {
		tx.Rollback()
		logrus.Error("failed to insert Department:", err)
		return err
	}

	if err != nil {
		tx.Rollback()
		logrus.Error("failed to insert DepartmentPosition:", err)
		return err
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		logrus.Error("failed to commit transaction:", err)
		return err
	}

	return nil
}

func (t *OrganizationPostgres) GetEmployeeById(employeeId string) (*models.Employee, error) {
	var output models.Employee

	err := t.db.Get(&output, `
	SELECT *
	FROM "Employee" WHERE id = $1
`, employeeId)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return &output, nil
}

func (t *OrganizationPostgres) GetDepartmentByEmployeeId(employeeId string) (*models.Department, error) {
	var output models.Department

	err := t.db.Get(&output, `
		SELECT d.id, d.name
		FROM "EmployeeDepartment" ed
		JOIN "Department" d ON ed.department_id = d.id
		WHERE ed.employee_id = $1
		LIMIT 1
	`, employeeId)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return &output, nil
}

func (t *OrganizationPostgres) GetPositionByEmployeeId(employeeId string) (*models.Position, error) {
	var output models.Position

	err := t.db.Get(&output, `
		SELECT p.*
		FROM "EmployeePosition" ep
		JOIN "Position" p ON ep.position_id = p.id
		WHERE ep.employee_id = $1
		LIMIT 1
	`, employeeId)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return &output, nil
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

func (t *OrganizationPostgres) GetDepartmentPosition(accountId string) (*models.GetDepartmentPositionOutput, error) {
	var output models.GetDepartmentPositionOutput

	err := t.db.Select(&output.DepartmentPosition, `
			SELECT dp.*
		FROM "DepartmentPosition" dp
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

func (t *OrganizationPostgres) UpdateEmployee(input *models.Employee) error {
	_, err := t.db.Exec(`UPDATE "Employee" SET
                      Id = $1,
                      full_name = $2,
                      phone_number = $3,
                      birth_date = $4,
                      employment_date = $5,
                      residential_address = $6,
                      marital_status = $7,
                      email = $8
                      WHERE Id = $1`,
		input.ID,
		input.FullName,
		input.PhoneNumber,
		input.BirthDate,
		input.EmploymentDate,
		input.ResidentialAddress,
		input.MaritalStatus,
		input.Email,
	)
	if err != nil {
		logrus.Error(err.Error())
		return err
	}

	return nil
}

func (t *OrganizationPostgres) UpdatePositionById(input *models.UpdatePositionInput) error {
	_, err := t.db.Exec(`UPDATE "Position" SET
                      id = $1,
                      name = $2
                      WHERE id = $1`,
		input.ID,
		input.Name,
	)

	_, err = t.db.Exec(`
		INSERT INTO "DepartmentPosition" (
			department_id, position_id
		) VALUES ($1, $2)
	`,
		input.DepartmentID,
		input.ID,
	)

	if err != nil {
		logrus.Error(err.Error())
		return err
	}

	return nil
}

func (t *OrganizationPostgres) UpdateDepartmentById(accountId string, input *models.UpdateDepartmentInput) error {
	var organization models.Organization

	err := t.db.Get(&organization, `
		SELECT o.*
		FROM "Organization" o
	JOIN "Department" d ON d.organization_id = o.id
	JOIN "EmployeeDepartment" ed ON ed.department_id = d.id
	JOIN "Employee" e ON e.id = ed.employee_id
	WHERE e.id = $1
	LIMIT 1;
	`, accountId)

	_, err = t.db.Exec(`UPDATE "Department" SET
                      id = $1,
                      name = $2,
                      organization_id = $3
                      WHERE id = $1`,
		input.ID,
		input.Name,
		organization.ID,
	)
	if err != nil {
		logrus.Error(err.Error())
		return err
	}

	return nil
}

func (t *OrganizationPostgres) UpdateAccount(input *models.Account) error {
	_, err := t.db.Exec(`UPDATE "Account" SET
                      login = $1,
                      password = $2,
                      role = $3
                      WHERE Id = $4`,
		input.Login,
		input.Password,
		input.Role,
		input.ID,
	)
	if err != nil {
		logrus.Error(err.Error())
		return err
	}

	return nil
}

func (t *OrganizationPostgres) UpdateEmployeeDepartment(input *models.UpdateEmployeeDepartment) error {
	_, err := t.db.Exec(`UPDATE "EmployeeDepartment" SET
                      department_id = $1
                      WHERE employee_id = $2`,
		input.DepartmentID,
		input.EmployeeID,
	)
	if err != nil {
		logrus.Error(err.Error())
		return err
	}

	return nil
}

func (t *OrganizationPostgres) UpdateEmployeePosition(input *models.UpdateEmployeePosition) error {
	_, err := t.db.Exec(`UPDATE "EmployeePosition" SET
                      position_id = $1
                      WHERE employee_id = $2`,
		input.PositionID,
		input.EmployeeID,
	)
	if err != nil {
		logrus.Error(err.Error())
		return err
	}

	return nil
}

func (t *OrganizationPostgres) DeleteEmployeeById(employeeId string) error {
	tx, err := t.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	queries := []string{
		`DELETE FROM "EmployeePosition" WHERE "employee_id" = $1`,
		`DELETE FROM "EmployeeDepartment" WHERE "employee_id" = $1`,
		`DELETE FROM "Account" WHERE "id" = $1`,
		`DELETE FROM "Employee" WHERE "id" = $1`,
	}

	for _, query := range queries {
		if _, execErr := tx.Exec(query, employeeId); execErr != nil {
			err = execErr
			return err
		}
	}

	return nil
}
