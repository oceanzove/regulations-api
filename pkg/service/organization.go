package service

import (
	"regulations-api/models"
	"regulations-api/pkg/repository"
)

type OrganizationService struct {
	repo repository.Organization
}

func NewOrganizationService(repo repository.Organization) *OrganizationService {
	return &OrganizationService{repo: repo}
}

func (o *OrganizationService) GetDepartments(accountId string) (*models.GetDepartmentOutput, error) {
	return o.repo.GetDepartments(accountId)
}

func (o *OrganizationService) GetDepartmentByID(accountId string, departmentId string) (*models.Department, error) {
	return o.repo.GetDepartmentByID(accountId, departmentId)
}

func (o *OrganizationService) GetPositions(accountId string) (*models.GetPositionOutput, error) {
	return o.repo.GetPositions(accountId)
}

func (o *OrganizationService) GetPositionsByDepartment(accountId string, departmentId string) (*models.GetPositionOutput, error) {
	return o.repo.GetPositionsByDepartment(accountId, departmentId)
}

func (o *OrganizationService) GetEmployees(accountId string) (*models.GetEmployeesOutput, error) {
	return o.repo.GetEmployees(accountId)
}

func (o *OrganizationService) GetEmployeeById(employeeId string) (*models.Employee, error) {
	return o.repo.GetEmployeeById(employeeId)
}

func (o *OrganizationService) GetDepartmentByEmployeeId(employeeId string) (*models.Department, error) {
	return o.repo.GetDepartmentByEmployeeId(employeeId)
}

func (o *OrganizationService) GetPositionByEmployeeId(employeeId string) (*models.Position, error) {
	return o.repo.GetPositionByEmployeeId(employeeId)
}

func (o *OrganizationService) CreateEmployee(accountId string, input *models.CreateEmployeeInput) error {
	return o.repo.CreateEmployee(input)
}

func (o *OrganizationService) UpdateEmployee(input *models.Employee) error {
	return o.repo.UpdateEmployee(input)
}

func (o *OrganizationService) UpdateAccount(input *models.Account) error {
	return o.repo.UpdateAccount(input)
}

func (o *OrganizationService) UpdateEmployeeDepartment(input *models.UpdateEmployeeDepartment) error {
	return o.repo.UpdateEmployeeDepartment(input)
}

func (o *OrganizationService) UpdateEmployeePosition(input *models.UpdateEmployeePosition) error {
	return o.repo.UpdateEmployeePosition(input)
}
