package service

import (
	"st_api/pkg/models"
	"st_api/pkg/repository"
)

type DepositsService struct {
	repo repository.Deposits
}

func NewDepositsService(repo repository.Deposits) *DepositsService {
	return &DepositsService{repo: repo}
}

func (d *DepositsService) GetInRegion(regionId int) ([]models.Deposit, error) {
	return d.repo.GetInRegion(regionId)
}
