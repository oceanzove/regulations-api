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

func (o *ProcessService) Create(email string) (*models.CreateProcessOutput, error) {
	return o.repo.Create(email)
}

func (o *ProcessService) GetPrivate(email string) (*models.GetProcessesOutput, error) {
	return o.repo.GetPrivate(email)
}

func (o *ProcessService) UpdatePrivate(input models.UpdateProcessInput, email string) error {
	return o.repo.UpdatePrivate(input, email)
}
