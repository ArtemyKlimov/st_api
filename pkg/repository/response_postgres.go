package repository

import (
	"fmt"
	"st_api/pkg/models"

	"github.com/jmoiron/sqlx"
)

type ResponsePostgres struct {
	db *sqlx.DB
}

func NewResponsePostgres(db *sqlx.DB) *ResponsePostgres {
	return &ResponsePostgres{db: db}
}

func (r *ResponsePostgres) GetResponseByRegionForMineral(regionId int, mineralId int) ([]models.Response, error) {
	var responses []models.Response
	query := fmt.Sprintf(`SELECT p.Response AS Response, 
									p.t_trend AS t_trend, 
									p.g_temp AS g_temp, 
									p.Oopt AS oopt, 
									p.coast_dist AS coast_dist,
									p.coast_rate AS coast_rate,
									p.name AS name,
									p.perm_distr AS perm_distr,
									p.yed_dist AS yed_dist
							FROM %s p
							JOIN %s kp ON p.id = kp.propertyId
							WHERE p.sub_rf = %d AND kp.kindnameid = %d`,
		propertiesTable, mineralsPropertiesTable, regionId, mineralId)
	if err := r.db.Select(&responses, query); err != nil {
		return nil, err
	}
	return responses, nil
}
