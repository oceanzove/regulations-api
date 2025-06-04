package usecase

import (
	"github.com/sirupsen/logrus"
	"regulations-api/models"
)

func (u *Usecase) CreateRegulation(accountId string, input *models.CreateRegulationInput) ErrorCode {
	err := u.services.Regulation.Create(accountId, input)
	if err != nil {
		logrus.Error(err)
		return InternalServerError
	}

	return Success
}

func (u *Usecase) GetRegulation(accountId string) (*models.GetRegulationsOutput, ErrorCode) {
	output, err := u.services.Regulation.GetPrivate(accountId)
	if err != nil {
		logrus.Error(err)
		return nil, InternalServerError
	}

	return output, Success
}

func (u *Usecase) GetRegulationByID(accountID, regulationID string) (*models.Regulation, ErrorCode) {
	regulation, err := u.services.Regulation.GetByID(accountID, regulationID)
	if err != nil {
		return nil, InternalServerError
	}
	return regulation, Success
}

func (u *Usecase) UpdateRegulation(input models.UpdateRegulationInput, accountId string) ErrorCode {
	err := u.services.Regulation.UpdatePrivate(input, accountId)
	if err != nil {
		logrus.Error(err)
		return InternalServerError
	}

	return Success
}
