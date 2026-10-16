package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

// EnvVars guarda as configurações da aplicação lidas das variáveis de ambiente.
type EnvVars struct {
	DBConnectionString string
}

// Load lê as variáveis de ambiente e retorna erro se faltar alguma obrigatória.
func Load() (*EnvVars, error) {
	// Em desenvolvimento, carrega o .env da raiz do repositório (rodando de dentro de backend/).
	// No Docker o arquivo não existe e as variáveis vêm do compose, então o erro é ignorado.
	_ = godotenv.Load("../.env")

	dbConnectionString := os.Getenv("DB_CONNECTION_STRING")
	if dbConnectionString == "" {
		return nil, errors.New("variável de ambiente DB_CONNECTION_STRING não definida")
	}

	return &EnvVars{
		DBConnectionString: dbConnectionString,
	}, nil
}
