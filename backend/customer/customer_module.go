package customer

import (
	"github.com/jmoiron/sqlx"
	"github.com/viniciusalexandrre/oficina/backend/config"
)

type CustomerModule struct {
	Db  *sqlx.DB
	Env *config.EnvVars
}

func NewCustomerModule(db *sqlx.DB, env *config.EnvVars) *CustomerModule {
	return &CustomerModule{
		Db:  db,
		Env: env,
	}
}
