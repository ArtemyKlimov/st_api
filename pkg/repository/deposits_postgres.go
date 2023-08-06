package repository

import (
	"fmt"
	"st_api/pkg/models"

	"github.com/jmoiron/sqlx"
)

type DepositsPostgres struct {
	db *sqlx.DB
}

func NewDepositsPostgres(db *sqlx.DB) *DepositsPostgres {
	return &DepositsPostgres{db: db}
}

func (d *DepositsPostgres) GetInRegion(regionId int) ([]models.Deposit, error) {
	var deposits []models.Deposit
	query := fmt.Sprintf(`SELECT DISTINCT name 
	FROM %s
	WHERE sub_rf = %d`,
		propertiesTable, regionId)
	if err := d.db.Select(&deposits, query); err != nil {
		return nil, err
	}
	return deposits, nil
}
