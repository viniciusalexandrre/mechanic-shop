package car

import (
	"github.com/jmoiron/sqlx"
	"github.com/viniciusalexandrre/oficina/backend/config"
)

type CarModule struct {
	Db  *sqlx.DB
	Env *config.EnvVars
}

func NewCarModule(db *sqlx.DB, env *config.EnvVars) *CarModule {
	return &CarModule{
		Db:  db,
		Env: env,
	}
}
