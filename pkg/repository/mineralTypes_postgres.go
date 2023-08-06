package repository

import (
	"fmt"
	"st_api/pkg/models"

	"github.com/jmoiron/sqlx"
)

type MineralTypesPostgres struct {
	db *sqlx.DB
}

func NewMineralTypesPostgres(db *sqlx.DB) *MineralTypesPostgres {
	return &MineralTypesPostgres{db: db}
}

func (m *MineralTypesPostgres) GetAll() ([]models.MineralTypes, error) {
	var mineraltypes []models.MineralTypes

	query := fmt.Sprintf(`SELECT DISTINCT gr.grnameid AS id, g.name AS name 
	FROM %s gr
	JOIN %s p ON gr.propertyid = p.id 
	JOIN %s g ON gr.grnameid = g.id 
	ORDER BY gr.grnameid`,
		mineralTypesPropertiesTable, propertiesTable, mineralTypesTable)

	if err := m.db.Select(&mineraltypes, query); err != nil {
		return nil, err
	}
	return mineraltypes, nil
}

func (m *MineralTypesPostgres) GetAllInRegion(regionId int) ([]models.MineralTypes, error) {
	var mineraltypes []models.MineralTypes

	query := fmt.Sprintf(`SELECT DISTINCT gr.grnameid AS id, g.name AS name 
	FROM %s gr
	JOIN %s p ON gr.propertyid = p.id 
	JOIN %s g ON gr.grnameid = g.id 
	WHERE p.sub_rf = %d 
	ORDER BY gr.grnameid`,
		mineralTypesPropertiesTable, propertiesTable, mineralTypesTable, regionId)

	if err := m.db.Select(&mineraltypes, query); err != nil {
		return nil, err
	}
	return mineraltypes, nil
}
