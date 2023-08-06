package service

import (
	"st_api/pkg/models"
	"st_api/pkg/repository"
)

type MineralsService struct {
	repo repository.Minerals
}

func NewMineralsService(repo repository.Minerals) *MineralsService {
	return &MineralsService{repo: repo}
}

func (m *MineralsService) GetAll() ([]models.Minerals, error) {
	return m.repo.GetAll()
}
func (m *MineralsService) GetAllInRegion(regionId int) ([]models.Minerals, error) {
	return m.repo.GetAllInRegion(regionId)
}
func (m *MineralsService) GetAllInRegionByType(regionId, mineralTypeId int) ([]models.Minerals, error) {
	return m.repo.GetAllInRegionByType(regionId, mineralTypeId)
}
func (m *MineralsService) GetByType(mineralTypeId int) ([]models.Minerals, error) {
	return m.repo.GetByType(mineralTypeId)
}
