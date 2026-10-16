package main

import (
	"log"
	"net/http"

	"github.com/viniciusalexandrre/oficina/backend/config"
	"github.com/viniciusalexandrre/oficina/backend/database"
)

func main() {
	router := Routes()
	env, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.DatabaseConnection(env)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Println("servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}
