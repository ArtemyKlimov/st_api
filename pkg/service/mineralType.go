package service

import (
	"st_api/pkg/models"
	"st_api/pkg/repository"
)

type MineralTypeService struct {
	repo repository.MineralTypes
}

func NewMineralTypeService(repo repository.MineralTypes) *MineralTypeService {
	return &MineralTypeService{repo: repo}
}

func (s *MineralTypeService) GetAll() ([]models.MineralTypes, error) {
	return s.repo.GetAll()
}

func (s *MineralTypeService) GetAllInRegion(regionId int) ([]models.MineralTypes, error) {
	return s.repo.GetAllInRegion(regionId)
}
