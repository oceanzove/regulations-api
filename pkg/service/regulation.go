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

func (o *RegulationService) GetPrivate(email string) (*models.GetRegulationsOutput, error) {
	return o.repo.GetPrivate(email)
}

func (o *RegulationService) UpdatePrivate(input models.UpdateRegulationInput, email string) error {
	return o.repo.UpdatePrivate(input, email)
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
