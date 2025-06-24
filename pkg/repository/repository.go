package repository

import (
	"github.com/jmoiron/sqlx"
	"regulations-api/models"
)

type Sources struct {
	BusinessDB *sqlx.DB
}

type Auth interface {
}

type Account interface {
	Get(login string) (*models.Account, error)
	GetByID(id string) (*models.Account, error)
}

type Regulation interface {
	GetPrivate(accountID string) (*models.GetRegulationsOutput, error)
	GetByID(accountID string, regulationID string) (*models.Regulation, error)
	UpdatePrivate(input models.UpdateRegulationInput) error
	Create(accountID string, input *models.CreateRegulationInput) error
	CreateSection(accountID string, input *models.CreateSectionInput) error
	GetSections(accountID string) (*models.GetSectionsOutput, error)
	DeleteRegulationById(regulationId string) error
	LinkSectionToRegulation(input *models.LinkSectionToRegulation) error
	UnlinkSectionToRegulation(regulationID, sectionID string) error
	GetSectionById(regulationID string) (*models.GetSectionByRegulationOutput, error)
}

type Process interface {
	GetPrivate(accountId string) (*models.GetProcessesOutput, error)
	GetByID(accountId string, processId string) (*models.Process, error)
	UpdatePrivate(input *models.UpdateProcessInput) error
	Create(accountId string, input *models.CreateProcessInput) error
	LinkRegulationToProcess(processID, regulationID string) error
	UnlinkRegulationToProcess(processID, regulationID string) error
	CreateStep(input *models.Step) error
	GetStepsByProcess(processId string) ([]*models.Step, error)
	GetRegulationsByProcess(processId string) ([]*models.Regulation, error)
	DeleteProcessById(processId string) error
	UpdateStepById(input *models.Step) error
	DeleteStepById(stepId string) error
}

type Step interface {
	CreateSteps(input *models.CreateStepsInput) error
}

type Organization interface {
	GetDepartments(accountId string) (*models.GetDepartmentOutput, error)
	GetDepartmentByID(accountId string, departmentId string) (*models.Department, error)
	GetPositions(accountId string) (*models.GetPositionOutput, error)
	GetPositionsByDepartment(accountId string, departmentId string) (*models.GetPositionOutput, error)
	GetEmployees(accountId string) (*models.GetEmployeesOutput, error)
	GetEmployeeById(employeeId string) (*models.Employee, error)
	GetDepartmentByEmployeeId(employeeId string) (*models.Department, error)
	GetPositionByEmployeeId(employeeId string) (*models.Position, error)
	CreateEmployee(input *models.CreateEmployeeInput) error
	UpdateEmployee(input *models.Employee) error
	UpdateAccount(input *models.Account) error
	UpdateEmployeeDepartment(input *models.UpdateEmployeeDepartment) error
	UpdateEmployeePosition(input *models.UpdateEmployeePosition) error
	GetEmployeeDepartment(accountId string) (*models.GetEmployeeDepartmentOutput, error)
	GetEmployeePosition(accountId string) (*models.GetEmployeePositionOutput, error)
	DeleteEmployeeById(employeeId string) error
	GetDepartmentPosition(accountId string) (*models.GetDepartmentPositionOutput, error)
	CreatePosition(input *models.CreatePositionInput) error
	CreateDepartment(accountID string, input *models.CreateDepartmentInput) error
	UpdatePositionById(input *models.UpdatePositionInput) error
	UpdateDepartmentById(accountId string, input *models.UpdateDepartmentInput) error
}

type Repository struct {
	Account
	Auth
	Regulation
	Process
	Step
	Organization
}

func NewRepository(sources *Sources) *Repository {
	return &Repository{
		Account:      NewAccountPostgres(sources.BusinessDB),
		Auth:         NewAuthPostgres(sources.BusinessDB),
		Regulation:   NewRegulationPostgres(sources.BusinessDB),
		Process:      NewProcessPostgres(sources.BusinessDB),
		Organization: NewOrganizationPostgres(sources.BusinessDB),
	}
}
