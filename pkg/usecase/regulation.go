package usecase

import (
	"github.com/sirupsen/logrus"
	"regulations-api/models"
)

func (u *Usecase) CreateRegulation(email string) (*models.CreateRegulationOutput, ErrorCode) {
	output, err := u.services.Regulation.Create(email)
	if err != nil {
		logrus.Error(err)
		return nil, InternalServerError
	}

	return output, Success
}

func (u *Usecase) GetRegulation(email string) (*models.GetRegulationsOutput, ErrorCode) {
	output, err := u.services.Regulation.GetPrivate(email)
	if err != nil {
		logrus.Error(err)
		return nil, InternalServerError
	}

	return output, Success
}

func (u *Usecase) UpdateRegulation(input models.UpdateRegulationInput, email string) ErrorCode {
	err := u.services.Regulation.UpdatePrivate(input, email)
	if err != nil {
		logrus.Error(err)
		return InternalServerError
	}

	return Success
}
