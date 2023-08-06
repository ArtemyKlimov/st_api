package repository

import (
	"fmt"
	"st_api/pkg/models"

	"github.com/jmoiron/sqlx"
)

type MineralsPostgres struct {
	db *sqlx.DB
}

func NewMineralsPostgres(db *sqlx.DB) *MineralsPostgres {
	return &MineralsPostgres{db: db}
}

func (m *MineralsPostgres) GetAll() ([]models.Minerals, error) {
	var minerals []models.Minerals
	query := fmt.Sprintf(`SELECT DISTINCT kn.kindnameid AS id, k.name AS name 
	FROM %s kn
	JOIN %s p ON kn.propertyid = p.id 
	JOIN %s k ON kn.kindnameid = k.id
	ORDER BY kn.kindnameid`,
		mineralsPropertiesTable, propertiesTable, mineralsTable)
	if err := m.db.Select(&minerals, query); err != nil {
		return nil, err
	}
	return minerals, nil
}

func (m *MineralsPostgres) GetAllInRegion(regionId int) ([]models.Minerals, error) {
	var minerals []models.Minerals
	query := fmt.Sprintf(`SELECT DISTINCT kn.kindnameid AS id, k.name AS name 
	FROM %s kn
	JOIN %s p ON kn.propertyid = p.id 
	JOIN %s k ON kn.kindnameid = k.id 
	WHERE p.sub_rf = %d  
	ORDER BY kn.kindnameid`,
		mineralsPropertiesTable, propertiesTable, mineralsTable, regionId)
	if err := m.db.Select(&minerals, query); err != nil {
		return nil, err
	}
	return minerals, nil
}
func (m *MineralsPostgres) GetAllInRegionByType(regionId, mineralTypeId int) ([]models.Minerals, error) {
	var minerals []models.Minerals
	query := fmt.Sprintf(`SELECT DISTINCT kn.kindnameid AS id, k.name AS name 
	FROM %s kn
	JOIN %s p ON kn.propertyid = p.id 
	JOIN %s k ON kn.kindnameid = k.id 
	JOIN %s g ON g.propertyid = p.id 
	WHERE p.sub_rf = %d AND g.grnameid = %d 
	ORDER BY kn.kindnameid`,
		mineralsPropertiesTable, propertiesTable, mineralsTable, mineralTypesPropertiesTable, regionId, mineralTypeId)
	if err := m.db.Select(&minerals, query); err != nil {
		return nil, err
	}
	return minerals, nil
}
func (m *MineralsPostgres) GetByType(mineralTypeId int) ([]models.Minerals, error) {
	var minerals []models.Minerals
	query := fmt.Sprintf(`SELECT DISTINCT kn.kindnameid AS id, k.name AS name 
	FROM %s kn
	JOIN %s p ON kn.propertyid = p.id 
	JOIN %s k ON kn.kindnameid = k.id 
	JOIN %s g ON g.propertyid = p.id 
	WHERE g.grnameid = %d 
	ORDER BY kn.kindnameid`,
		mineralsPropertiesTable, propertiesTable, mineralsTable, mineralTypesPropertiesTable, mineralTypeId)
	if err := m.db.Select(&minerals, query); err != nil {
		return nil, err
	}
	return minerals, nil
}
