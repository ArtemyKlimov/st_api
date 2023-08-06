package service

import (
	"st_api/pkg/models"
	"st_api/pkg/repository"
)

type ResponseService struct {
	repo repository.Response
}

func NewResponseService(repo repository.Response) *ResponseService {
	return &ResponseService{repo: repo}
}

func (r *ResponseService) GetResponseByRegionForMineral(regionId int, mineralId int) ([]models.Response, error) {
	return r.repo.GetResponseByRegionForMineral(regionId, mineralId)
}
