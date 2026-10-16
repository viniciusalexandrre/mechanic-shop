package database

import (
	"database/sql"
	"fmt"

	"github.com/viniciusalexandrre/oficina/backend/config"

	_ "github.com/lib/pq"
)

// DatabaseConnection abre o pool de conexões com o PostgreSQL e confirma que o banco responde.
func DatabaseConnection(env *config.EnvVars) (*sql.DB, error) {
	db, err := sql.Open("postgres", env.DBConnectionString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("ping no banco: %w", err)
	}

	return db, nil
}
