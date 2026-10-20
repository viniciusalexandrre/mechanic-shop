package customer

import (
	"github.com/jmoiron/sqlx"
	"github.com/viniciusalexandrre/oficina/backend/config"
)

// CustomerModule guarda as dependências do pacote customer.
// Os handlers (customer_controller.go) e as queries (customer_service.go) são métodos dele.
type CustomerModule struct {
	Db  *sqlx.DB
	Env *config.EnvVars
}

// NewCustomerModule é chamado pelo fx, que entrega o banco e a config já prontos.
func NewCustomerModule(db *sqlx.DB, env *config.EnvVars) *CustomerModule {
	return &CustomerModule{
		Db:  db,
		Env: env,
	}
}
