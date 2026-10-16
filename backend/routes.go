package main

import (
	"net/http"

	"github.com/viniciusalexandrre/oficina/backend/customer"
)

func Routes() *http.ServeMux {

	router := http.NewServeMux()
	router.HandleFunc("POST /create_customers", customer.CreateCustomer)

	return router
}
