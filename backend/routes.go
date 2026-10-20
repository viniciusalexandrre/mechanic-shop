package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/viniciusalexandrre/oficina/backend/customer"
)

// RegisterRoutes liga cada rota ao método do módulo correspondente.
// O fx chama esta função entregando o app e os módulos já criados.
func RegisterRoutes(app *fiber.App, customerModule *customer.CustomerModule) {
	app.Post("/create_customers", customerModule.CreateCustomer)
	app.Put("/edit_customer/:id", customerModule.UpdateCustomer)
	app.Delete("/delete_customer/:id", customerModule.DeleteCustomer)
	app.Get("/list_customers", customerModule.ListCustomers)
}
