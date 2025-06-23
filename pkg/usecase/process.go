package usecase

import (
	"github.com/sirupsen/logrus"
	"regulations-api/models"
)

func (u *Usecase) CreateProcess(accountId string, input *models.CreateProcessInput) ErrorCode {
	err := u.services.Process.Create(accountId, input)
	if err != nil {
		logrus.Error(err)
		return InternalServerError
	}

	return Success
}

func (u *Usecase) GetProcesses(accountId string) (*models.GetProcessesOutput, ErrorCode) {
	output, err := u.services.Process.GetPrivate(accountId)
	if err != nil {
		logrus.Error(err)
		return nil, InternalServerError
	}

	return output, Success
}

func (u *Usecase) GetProcessByID(accountID, processID string) (*models.Process, ErrorCode) {
	process, err := u.services.Process.GetByID(accountID, processID)
	if err != nil {
		return nil, InternalServerError
	}
	return process, Success
}

func (u *Usecase) UpdateProcess(input *models.UpdateProcessInput) ErrorCode {
	err := u.services.Process.UpdatePrivate(input)
	if err != nil {
		logrus.Error(err)
		return InternalServerError
	}

	return Success
}

func (u *Usecase) UpdateStepById(input *models.Step) ErrorCode {
	err := u.services.Process.UpdateStepById(input)
	if err != nil {
		logrus.Error(err)
		return InternalServerError
	}

	return Success
}

func (u *Usecase) LinkRegulationToProcess(processID, regulationID string) ErrorCode {
	err := u.services.Process.LinkRegulationToProcess(processID, regulationID)
	if err != nil {
		logrus.Error(err)
		return InternalServerError
	}
	return Success
}

func (u *Usecase) UnlinkRegulationToProcess(processID, regulationID string) ErrorCode {
	err := u.services.Process.UnlinkRegulationToProcess(processID, regulationID)
	if err != nil {
		logrus.Error(err)
		return InternalServerError
	}
	return Success
}

func (u *Usecase) GetStepsByProcess(processID string) ([]*models.Step, error) {
	return u.services.Process.GetStepsByProcess(processID)
}

func (u *Usecase) GetRegulationsByProcess(processID string) ([]*models.Regulation, error) {
	return u.services.Process.GetRegulationsByProcess(processID)
}

func (u *Usecase) CreateStep(input *models.Step) ErrorCode {
	err := u.services.Process.CreateStep(input)
	if err != nil {
		logrus.Error(err)
		return InternalServerError
	}
	return Success
}

func (u *Usecase) DeleteProcessById(processId string) ErrorCode {
	err := u.services.Process.DeleteProcessById(processId)
	if err != nil {
		logrus.Error(err)
		return InternalServerError
	}

	return Success
}

func (u *Usecase) DeleteStepById(stepId string) ErrorCode {
	err := u.services.Process.DeleteStepById(stepId)
	if err != nil {
		logrus.Error(err)
		return InternalServerError
	}

	return Success
}
