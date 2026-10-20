package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type EnvVars struct {
	DB_CONNECTION_STRING string
}

func LoadEnv() (*EnvVars, error) {
	_ = godotenv.Load("../.env")
	DB_CONNECTION_STRING := os.Getenv("DB_CONNECTION_STRING")
	if DB_CONNECTION_STRING == "" {
		return nil, errors.New("Undefined environment variable: DB_CONNECTION_STRING")
	}
	return &EnvVars{DB_CONNECTION_STRING: DB_CONNECTION_STRING}, nil
}
