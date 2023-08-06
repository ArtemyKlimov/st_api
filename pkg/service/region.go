package service

import (
	"st_api/pkg/models"
	"st_api/pkg/repository"
)

type RegionService struct {
	repo repository.Regions
}

func NewRegionService(repo repository.Regions) *RegionService {
	return &RegionService{repo: repo}
}

func (s *RegionService) GetAll() ([]models.Region, error) {
	return s.repo.GetAll()
}
