package usecase

import (
	"github.com/sirupsen/logrus"
	"regulations-api/models"
)

func (u *Usecase) CreateSteps(input *models.CreateStepsInput) ErrorCode {
	err := u.services.Step.Create(input)
	if err != nil {
		logrus.Error(err)
		return InternalServerError
	}

	return Success
}
