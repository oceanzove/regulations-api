package service

import (
	"regulations-api/models"
	"regulations-api/pkg/repository"
)

type Account interface {
	Get(login string) (*models.Account, error)
	GetByID(id string) (*models.Account, error)
}

type Auth interface {
	SignIn(input *models.SignInInput, accountPassword string) error
}

type Regulation interface {
	GetPrivate(accountId string) (*models.GetRegulationsOutput, error)
	GetByID(accountId string, regulationID string) (*models.Regulation, error)
	UpdatePrivate(input models.UpdateRegulationInput, accountId string) error
	Create(accountId string, input *models.CreateRegulationInput) error
	CreateSection(accountId string, input *models.CreateSectionInput) error
	GetSections(accountId string) (*models.GetSectionsOutput, error)
}

type Process interface {
	GetPrivate(accountId string) (*models.GetProcessesOutput, error)
	GetByID(accountId string, processId string) (*models.Process, error)
	UpdatePrivate(input *models.UpdateProcessInput, accountId string) error
	Create(accountId string, input *models.CreateProcessInput) error
	LinkRegulationToProcess(accountId string, processId string) error
	GetStepsByProcess(processId string) ([]*models.Step, error)
	GetRegulationsByProcess(processId string) ([]*models.Regulation, error)
	CreateStep(input *models.Step) error
}

type Step interface {
	Create(input *models.CreateStepsInput) error
}

type Organization interface {
	GetDepartments(accountId string) (*models.GetDepartmentOutput, error)
	GetDepartmentByID(accountId string, departmentId string) (*models.Department, error)
	GetPositions(accountId string) (*models.GetPositionOutput, error)
	GetPositionsByDepartment(accountId string, departmentId string) (*models.GetPositionOutput, error)
	GetEmployees(accountId string) (*models.GetEmployeesOutput, error)
	GetEmployeeById(employee string) (*models.Employee, error)
	GetDepartmentByEmployeeId(employeeId string) (*models.Department, error)
	GetPositionByEmployeeId(employeeId string) (*models.Position, error)
	CreateEmployee(accountId string, input *models.CreateEmployeeInput) error
	UpdateEmployee(input *models.Employee) error
	UpdateAccount(input *models.Account) error
	UpdateEmployeeDepartment(input *models.UpdateEmployeeDepartment) error
	UpdateEmployeePosition(input *models.UpdateEmployeePosition) error
}

type JWTToken interface {
	GenerateAccessToken(account *models.Account) (string, error)
	GenerateRefreshToken(account *models.Account) (string, error)
	GenerateAccessFromRefresh(email string) (string, error)
	ParseToken(tokenString string) (*models.JWTClaims, error)
}

type Service struct {
	Account
	Auth
	Regulation
	Process
	JWTToken
	Step
	Organization
}

func NewService(repos *repository.Repository, config *models.ConfigService) *Service {
	return &Service{
		Account:      NewAccountService(repos.Account),
		Auth:         NewAuthService(repos.Auth),
		JWTToken:     NewJWTTokenService(config.Server, repos.Account),
		Regulation:   NewRegulationService(repos.Regulation),
		Process:      NewProcessService(repos.Process),
		Organization: NewOrganizationService(repos.Organization),
	}
}
