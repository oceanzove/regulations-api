package service

import (
	"regulations-api/models"
	"regulations-api/pkg/repository"
)

type ProcessService struct {
	repo repository.Process
}

func NewProcessService(repo repository.Process) *ProcessService {
	return &ProcessService{repo: repo}
}

func (o *ProcessService) Create(accountId string, input *models.CreateProcessInput) error {
	return o.repo.Create(accountId, input)
}

func (o *ProcessService) GetPrivate(accountId string) (*models.GetProcessesOutput, error) {
	return o.repo.GetPrivate(accountId)
}

func (o *ProcessService) GetByID(accountID, processID string) (*models.Process, error) {
	return o.repo.GetByID(accountID, processID)
}

func (o *ProcessService) UpdatePrivate(input *models.UpdateProcessInput) error {
	return o.repo.UpdatePrivate(input)
}

func (o *ProcessService) UpdateStepById(input *models.Step) error {
	return o.repo.UpdateStepById(input)
}

func (o *ProcessService) LinkRegulationToProcess(processID, regulationID string) error {
	return o.repo.LinkRegulationToProcess(processID, regulationID)
}

func (o *ProcessService) UnlinkRegulationToProcess(processID, regulationID string) error {
	return o.repo.UnlinkRegulationToProcess(processID, regulationID)
}

func (o *ProcessService) GetStepsByProcess(processID string) ([]*models.Step, error) {
	return o.repo.GetStepsByProcess(processID)
}

func (o *ProcessService) GetRegulationsByProcess(processID string) ([]*models.Regulation, error) {
	return o.repo.GetRegulationsByProcess(processID)
}

func (o *ProcessService) CreateStep(input *models.Step) error {
	return o.repo.CreateStep(input)
}

func (o *ProcessService) DeleteProcessById(processId string) error {
	return o.repo.DeleteProcessById(processId)
}

func (o *ProcessService) DeleteStepById(stepId string) error {
	return o.repo.DeleteStepById(stepId)
}
