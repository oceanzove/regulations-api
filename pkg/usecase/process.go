package usecase

import (
	"github.com/sirupsen/logrus"
	"regulations-api/models"
)

func (u *Usecase) CreateProcess(email string) (*models.CreateProcessOutput, ErrorCode) {
	output, err := u.services.Process.Create(email)
	if err != nil {
		logrus.Error(err)
		return nil, InternalServerError
	}

	return output, Success
}

func (u *Usecase) GetProcesses(email string) (*models.GetProcessesOutput, ErrorCode) {
	output, err := u.services.Process.GetPrivate(email)
	if err != nil {
		logrus.Error(err)
		return nil, InternalServerError
	}

	return output, Success
}

func (u *Usecase) UpdateProcess(input models.UpdateProcessInput, email string) ErrorCode {
	err := u.services.Process.UpdatePrivate(input, email)
	if err != nil {
		logrus.Error(err)
		return InternalServerError
	}

	return Success
}
