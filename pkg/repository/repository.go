package repository

import (
	"st_api/pkg/models"

	"github.com/jmoiron/sqlx"
)

type Deposits interface {
	GetInRegion(regionId int) ([]models.Deposit, error)
}

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

type Response interface {
	GetResponseByRegionForMineral(regionId int, mineralId int) ([]models.Response, error)
}

type Repository struct {
	Regions
	Deposits
	Minerals
	MineralTypes
	Response
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Regions:      NewRegionsPostgres(db),
		MineralTypes: NewMineralTypesPostgres(db),
		Minerals:     NewMineralsPostgres(db),
		Deposits:     NewDepositsPostgres(db),
		Response:     NewResponsePostgres(db),
	}
}
