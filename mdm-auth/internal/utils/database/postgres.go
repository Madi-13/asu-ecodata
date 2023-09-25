package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func InitPostgresDb(dbUrl string) (*sqlx.DB, error) {
	dbConn, err := sqlx.Connect("postgres", dbUrl)
	if err != nil {
		return nil, err
	}

	return dbConn, nil
}
