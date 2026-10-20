package database

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/viniciusalexandrre/oficina/backend/config"

	_ "github.com/lib/pq"
)

func DatabaseConnection(env *config.EnvVars) *sqlx.DB {
	db, err := sqlx.Connect("postgres", env.DB_CONNECTION_STRING)
	if err != nil {
		log.Panic(err.Error())
	}
	return db
}
