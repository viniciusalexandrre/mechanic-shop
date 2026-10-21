package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/viniciusalexandrre/oficina/backend/customer"
)

func RegisterRoutes(app *fiber.App, customerModule *customer.CustomerModule) {
	app.Post("/create_customers", customerModule.CreateCustomer)
	app.Put("/edit_customer", customerModule.UpdateCustomer)
	app.Delete("/delete_customer", customerModule.DeleteCustomer)
	app.Get("/get_customer", customerModule.GetCustomerByID)
	app.Get("/list_customers", customerModule.ListCustomers)
	app.Get("/list_usernames_customers", customerModule.ListUserNamesCustomers)
}
