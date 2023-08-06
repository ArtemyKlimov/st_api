package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	schemaName                  = "Ecology"
	usersTable                  = schemaName + ".users"
	regionsTable                = schemaName + ".SUB_RF"
	mineralsTable               = schemaName + ".KIND_NAME"
	mineralTypesTable           = schemaName + ".GR_NAME"
	propertiesTable             = schemaName + ".Properties"
	mineralTypesPropertiesTable = schemaName + ".GRnameProperties"
	mineralsPropertiesTable     = schemaName + ".KindNameProperties"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
