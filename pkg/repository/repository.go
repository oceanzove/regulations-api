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
	UpdatePrivate(input models.UpdateRegulationInput, email string) error
	Create(accountID string, input *models.CreateRegulationInput) error
	CreateSection(accountID string, input *models.CreateSectionInput) error
	GetSections(accountID string) (*models.GetSectionsOutput, error)
}

type Process interface {
	GetPrivate(accountId string) (*models.GetProcessesOutput, error)
	GetByID(accountId string, processId string) (*models.Process, error)
	UpdatePrivate(input *models.UpdateProcessInput, accountId string) error
	Create(accountId string, input *models.CreateProcessInput) error
	LinkRegulationToProcess(accountId string, processId string) error
	CreateStep(input *models.Step) error
	GetStepsByProcess(processId string) ([]*models.Step, error)
	GetRegulationsByProcess(processId string) ([]*models.Regulation, error)
}

type Step interface {
	CreateSteps(input *models.CreateStepsInput) error
}

type Organization interface {
	GetDepartments(accountId string) (*models.GetDepartmentOutput, error)
	GetPositions(accountId string) (*models.GetPositionOutput, error)
	GetPositionsByDepartment(accountId string, departmentId string) (*models.GetPositionOutput, error)
	GetEmployees(accountId string) (*models.GetEmployeesOutput, error)
	CreateEmployee(input *models.CreateEmployeeInput) error
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
