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

func (u *Usecase) CreateSection(accountId string, input *models.CreateSectionInput) ErrorCode {
	err := u.services.Regulation.CreateSection(accountId, input)
	if err != nil {
		logrus.Error(err)
		return InternalServerError
	}

	return Success
}

func (u *Usecase) DeleteRegulationById(regulationId string) ErrorCode {
	err := u.services.Regulation.DeleteRegulationById(regulationId)
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

func (u *Usecase) GetSections(accountId string) (*models.GetSectionsOutput, ErrorCode) {
	output, err := u.services.Regulation.GetSections(accountId)
	if err != nil {
		logrus.Error(err)
		return nil, InternalServerError
	}

	return output, Success
}

func (u *Usecase) GetSectionById(regulationId string) (*models.GetSectionByRegulationOutput, ErrorCode) {
	output, err := u.services.Regulation.GetSectionById(regulationId)
	if err != nil {
		logrus.Error(err)
		return nil, InternalServerError
	}

	return output, Success
}

func (u *Usecase) LinkSectionToRegulation(input *models.LinkSectionToRegulation) ErrorCode {
	err := u.services.Regulation.LinkSectionToRegulation(input)
	if err != nil {
		logrus.Error(err)
		return InternalServerError
	}
	return Success
}

func (u *Usecase) UnlinkSectionToRegulation(regulationID, sectionID string) ErrorCode {
	err := u.services.Regulation.UnlinkSectionToRegulation(regulationID, sectionID)
	if err != nil {
		logrus.Error(err)
		return InternalServerError
	}
	return Success
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
