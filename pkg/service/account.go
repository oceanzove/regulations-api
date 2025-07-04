package service

import (
	"regulations-api/models"
	"regulations-api/pkg/repository"
)

type AccountService struct {
	repo repository.Account
}

func (a *AccountService) Get(login string) (*models.Account, error) {
	return a.repo.Get(login)
}

func (a *AccountService) GetByID(id string) (*models.Account, error) {
	return a.repo.GetByID(id)
}

func NewAccountService(repo repository.Account) *AccountService {
	return &AccountService{repo: repo}
}
