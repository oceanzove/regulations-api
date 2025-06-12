package usecase

import (
	"github.com/sirupsen/logrus"
	"regulations-api/models"
)

func (u *Usecase) GetAccount(accountID string) (*models.Account, ErrorCode) {
	employee, err := u.services.Account.GetByID(accountID)
	if err != nil {
		return nil, InternalServerError
	}
	return employee, Success
}

func (u *Usecase) GetDepartments(accountID string) (*models.GetDepartmentOutput, ErrorCode) {
	departments, err := u.services.Organization.GetDepartments(accountID)
	if err != nil {
		return nil, InternalServerError
	}
	return departments, Success
}

func (u *Usecase) GetDepartmentById(accountID string, departmentId string) (*models.Department, ErrorCode) {
	department, err := u.services.Organization.GetDepartmentByID(accountID, departmentId)
	if err != nil {
		return nil, InternalServerError
	}
	return department, Success
}

func (u *Usecase) GetPositions(accountID string) (*models.GetPositionOutput, ErrorCode) {
	positions, err := u.services.Organization.GetPositions(accountID)
	if err != nil {
		return nil, InternalServerError
	}
	return positions, Success
}

func (u *Usecase) GetPositionsByDepartment(accountId string, departmentId string) (*models.GetPositionOutput, ErrorCode) {
	positions, err := u.services.Organization.GetPositionsByDepartment(accountId, departmentId)
	if err != nil {
		return nil, InternalServerError
	}
	return positions, Success
}

func (u *Usecase) GetEmployees(accountID string) (*models.GetEmployeesOutput, ErrorCode) {
	employees, err := u.services.Organization.GetEmployees(accountID)
	if err != nil {
		return nil, InternalServerError
	}
	return employees, Success
}

func (u *Usecase) CreateEmployee(accountId string, input *models.CreateEmployeeInput) ErrorCode {
	err := u.services.Organization.CreateEmployee(accountId, input)
	if err != nil {
		logrus.Error(err)
		return InternalServerError
	}

	return Success
}
