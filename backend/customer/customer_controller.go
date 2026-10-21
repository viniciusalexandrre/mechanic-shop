package customer

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (m *CustomerModule) CreateCustomer(c *fiber.Ctx) error {
	var req CustomerRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "JSON inválido"})
	}

	if err := m.createCustomer(req); err != nil {
		return c.Status(err.Code).JSON(fiber.Map{"error": err.Message})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "ok"})
}

func (m *CustomerModule) UpdateCustomer(c *fiber.Ctx) error {
	id := c.Query("id")
	if _, err := uuid.Parse(id); err != nil {
		return c.Status(ErrInvalidID.Code).JSON(fiber.Map{"error": ErrInvalidID.Message})
	}

	var req CustomerRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "JSON inválido"})
	}

	if err := m.updateCustomer(id, req); err != nil {
		return c.Status(err.Code).JSON(fiber.Map{"error": err.Message})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "ok"})
}

func (m *CustomerModule) DeleteCustomer(c *fiber.Ctx) error {
	id := c.Query("id")
	if _, err := uuid.Parse(id); err != nil {
		return c.Status(ErrInvalidID.Code).JSON(fiber.Map{"error": ErrInvalidID.Message})
	}

	if err := m.deleteCustomer(id); err != nil {
		return c.Status(err.Code).JSON(fiber.Map{"error": err.Message})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "ok"})
}

func (m *CustomerModule) GetCustomerByID(c *fiber.Ctx) error {
	id := c.Query("id")
	if _, err := uuid.Parse(id); err != nil {
		return c.Status(ErrInvalidID.Code).JSON(fiber.Map{"error": ErrInvalidID.Message})
	}

	customer, err := m.getCustomerByID(id)
	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{"error": err.Message})
	}

	return c.Status(fiber.StatusOK).JSON(customer)
}

func (m *CustomerModule) ListCustomers(c *fiber.Ctx) error {
	params := ListParams{
		Search: c.Query("search"),
		SortBy: c.Query("sort_by", "created_at"),
		Order:  c.Query("order", "asc"),
		Page:   c.QueryInt("page", 1),
		Limit:  c.QueryInt("page_size", 20),
	}

	customers, err := m.listCustomers(params)
	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{"error": err.Message})
	}

	return c.Status(fiber.StatusOK).JSON(customers)
}

func (m *CustomerModule) ListUserNamesCustomers(c *fiber.Ctx) error {
	parm := c.Query("prefix")

	usernames, err := m.listUserNamesCustomers(parm)
	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{"error": err.Message})
	}

	return c.Status(fiber.StatusOK).JSON(usernames)
}
