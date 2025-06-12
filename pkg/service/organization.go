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

func (o *OrganizationService) GetPositions(accountId string) (*models.GetPositionOutput, error) {
	return o.repo.GetPositions(accountId)
}

func (o *OrganizationService) GetPositionsByDepartment(accountId string, departmentId string) (*models.GetPositionOutput, error) {
	return o.repo.GetPositionsByDepartment(accountId, departmentId)
}

func (o *OrganizationService) GetEmployees(accountId string) (*models.GetEmployeesOutput, error) {
	return o.repo.GetEmployees(accountId)
}

func (o *OrganizationService) CreateEmployee(accountId string, input *models.CreateEmployeeInput) error {
	return o.repo.CreateEmployee(input)
}
