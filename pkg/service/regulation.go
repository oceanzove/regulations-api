package service

import (
	"regulations-api/models"
	"regulations-api/pkg/repository"
)

type RegulationService struct {
	repo repository.Regulation
}

func NewRegulationService(repo repository.Regulation) *RegulationService {
	return &RegulationService{repo: repo}
}

//func (o *OfferService) GetPrivateByID(offerID int, email string) (*models.Offer, error) {
//	return o.repo.GetPrivateByID(offerID, email)
//}
//
//func (o *OfferService) Get(input *models.OfferGetInput) (*models.OfferGetActiveOutput, error) {
//	return o.repo.Get(input)
//}

func (o *RegulationService) Create(accountId string, input *models.CreateRegulationInput) error {
	return o.repo.Create(accountId, input)
}

func (o *RegulationService) CreateSection(accountId string, input *models.CreateSectionInput) error {
	return o.repo.CreateSection(accountId, input)
}

func (o *RegulationService) GetPrivate(accountId string) (*models.GetRegulationsOutput, error) {
	return o.repo.GetPrivate(accountId)
}

func (o *RegulationService) GetSections(accountId string) (*models.GetSectionsOutput, error) {
	return o.repo.GetSections(accountId)
}

func (o *RegulationService) GetByID(accountID, regulationID string) (*models.Regulation, error) {
	return o.repo.GetByID(accountID, regulationID)
}

func (o *RegulationService) UpdatePrivate(input models.UpdateRegulationInput, accountId string) error {
	return o.repo.UpdatePrivate(input, accountId)
}

//func (o *OfferService) Create(input *models.OfferCreateInput, email string) (*models.OfferCreateOutput, error) {
//	return o.repo.Create(input, email)
//}
//
//func NewOfferService(repo repository.Offer) *OfferService {
//	return &OfferService{repo: repo}
//}
//
//func (o *OfferService) UpdateStatus(input *models.OfferUpdateStatusInput) (*models.Offer, error) {
//	return o.repo.UpdateStatus(input)
//}
