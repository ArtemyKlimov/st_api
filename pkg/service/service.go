package service

import (
	"st_api/pkg/models"
	"st_api/pkg/repository"
)

type Regions interface {
	GetAll() ([]models.Region, error)
}

type MineralTypes interface {
	GetAll() ([]models.MineralTypes, error)
	GetAllInRegion(regionId int) ([]models.MineralTypes, error)
}

type Minerals interface {
	GetAll() ([]models.Minerals, error)
	GetAllInRegion(regionId int) ([]models.Minerals, error)
	GetAllInRegionByType(regionId, mineralTypeId int) ([]models.Minerals, error)
	GetByType(mineralTypeId int) ([]models.Minerals, error)
}

type Deposits interface {
	GetInRegion(regionId int) ([]models.Deposit, error)
}

type Response interface {
	GetResponseByRegionForMineral(regionId int, mineralId int) ([]models.Response, error)
}

type Service struct {
	Regions
	MineralTypes
	Minerals
	Deposits
	Response
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Regions:      NewRegionService(repos.Regions),
		MineralTypes: NewMineralTypeService(repos.MineralTypes),
		Minerals:     NewMineralsService(repos.Minerals),
		Deposits:     NewDepositsService(repos.Deposits),
		Response:     NewResponseService(repos.Response),
	}
}
