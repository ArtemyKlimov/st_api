package repository

import (
	"fmt"
	"st_api/pkg/models"

	"github.com/jmoiron/sqlx"
)

type RegionsPostgres struct {
	db *sqlx.DB
}

func NewRegionsPostgres(db *sqlx.DB) *RegionsPostgres {
	return &RegionsPostgres{db: db}
}

func (r *RegionsPostgres) GetAll() ([]models.Region, error) {
	var regions []models.Region

	query := fmt.Sprintf(`SELECT * FROM %s`, regionsTable)

	if err := r.db.Select(&regions, query); err != nil {
		return nil, err
	}
	return regions, nil
}
