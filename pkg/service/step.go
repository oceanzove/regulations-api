package service

import (
	"regulations-api/models"
	"regulations-api/pkg/repository"
)

type StepService struct {
	repo repository.Step
}

func NewStepService(repo repository.Step) *StepService {
	return &StepService{repo: repo}
}

func (o *StepService) Create(input *models.CreateStepsInput) error {
	return o.repo.CreateSteps(input)
}

//func (o *StepService) GetPrivate(email string) (*models.GetProcessesOutput, error) {
//	return o.repo.GetPrivate(email)
//}
//
//func (o *StepService) UpdatePrivate(input models.UpdateProcessInput, email string) error {
//	return o.repo.UpdatePrivate(input, email)
//}
