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

func (u *Usecase) GetEmployeeDepartment(accountId string) (*models.GetEmployeeDepartmentOutput, ErrorCode) {
	employees, err := u.services.Organization.GetEmployeeDepartment(accountId)
	if err != nil {
		return nil, InternalServerError
	}
	return employees, Success
}

func (u *Usecase) GetEmployeePosition(accountId string) (*models.GetEmployeePositionOutput, ErrorCode) {
	employees, err := u.services.Organization.GetEmployeePosition(accountId)
	if err != nil {
		return nil, InternalServerError
	}
	return employees, Success
}

func (u *Usecase) GetEmployeeById(employeeId string) (*models.Employee, ErrorCode) {
	employee, err := u.services.Organization.GetEmployeeById(employeeId)
	if err != nil {
		return nil, InternalServerError
	}
	return employee, Success
}

func (u *Usecase) GetDepartmentByEmployeeId(employeeId string) (*models.Department, ErrorCode) {
	employee, err := u.services.Organization.GetDepartmentByEmployeeId(employeeId)
	if err != nil {
		return nil, InternalServerError
	}
	return employee, Success
}

func (u *Usecase) GetDepartmentPosition(accountId string) (*models.GetDepartmentPositionOutput, ErrorCode) {
	employee, err := u.services.Organization.GetDepartmentPosition(accountId)
	if err != nil {
		return nil, InternalServerError
	}
	return employee, Success
}

func (u *Usecase) GetPositionByEmployeeId(employeeId string) (*models.Position, ErrorCode) {
	employee, err := u.services.Organization.GetPositionByEmployeeId(employeeId)
	if err != nil {
		return nil, InternalServerError
	}
	return employee, Success
}

func (u *Usecase) CreateEmployee(accountId string, input *models.CreateEmployeeInput) ErrorCode {
	err := u.services.Organization.CreateEmployee(accountId, input)
	if err != nil {
		logrus.Error(err)
		return InternalServerError
	}

	return Success
}

func (u *Usecase) CreatePosition(input *models.CreatePositionInput) ErrorCode {
	err := u.services.Organization.CreatePosition(input)
	if err != nil {
		logrus.Error(err)
		return InternalServerError
	}

	return Success
}

func (u *Usecase) CreateDepartment(accountId string, input *models.CreateDepartmentInput) ErrorCode {
	err := u.services.Organization.CreateDepartment(accountId, input)
	if err != nil {
		logrus.Error(err)
		return InternalServerError
	}

	return Success
}

func (u *Usecase) UpdateEmployee(input *models.Employee) ErrorCode {
	err := u.services.Organization.UpdateEmployee(input)
	if err != nil {
		logrus.Error(err)
		return InternalServerError
	}

	return Success
}

func (u *Usecase) UpdatePositionById(input *models.UpdatePositionInput) ErrorCode {
	err := u.services.Organization.UpdatePositionById(input)
	if err != nil {
		logrus.Error(err)
		return InternalServerError
	}

	return Success
}

func (u *Usecase) UpdateDepartmentById(accountId string, input *models.UpdateDepartmentInput) ErrorCode {
	err := u.services.Organization.UpdateDepartmentById(accountId, input)
	if err != nil {
		logrus.Error(err)
		return InternalServerError
	}

	return Success
}

func (u *Usecase) UpdateAccount(input *models.Account) ErrorCode {
	err := u.services.Organization.UpdateAccount(input)
	if err != nil {
		logrus.Error(err)
		return InternalServerError
	}

	return Success
}

func (u *Usecase) DeleteEmployeeById(employeeId string) ErrorCode {
	err := u.services.Organization.DeleteEmployeeById(employeeId)
	if err != nil {
		logrus.Error(err)
		return InternalServerError
	}

	return Success
}

func (u *Usecase) UpdateEmployeeDepartment(input *models.UpdateEmployeeDepartment) ErrorCode {
	err := u.services.Organization.UpdateEmployeeDepartment(input)
	if err != nil {
		logrus.Error(err)
		return InternalServerError
	}

	return Success
}

func (u *Usecase) UpdateEmployeePosition(input *models.UpdateEmployeePosition) ErrorCode {
	err := u.services.Organization.UpdateEmployeePosition(input)
	if err != nil {
		logrus.Error(err)
		return InternalServerError
	}

	return Success
}
