package customer

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
)

const pgForeignKeyViolation = "23503"

func validateCustomer(req CustomerRequest) *fiber.Error {
	switch req.Type {
	case "PF":
		if len(req.Document) != 11 {
			return fiber.NewError(fiber.StatusBadRequest, "CPF deve ter 11 dígitos")
		}
	case "PJ":
		if len(req.Document) != 14 {
			return fiber.NewError(fiber.StatusBadRequest, "CNPJ deve ter 14 dígitos")
		}
	default:
		return ErrInvalidType
	}

	switch {
	case req.Name == "":
		return ErrNameRequired
	}

	return nil
}

func (m *CustomerModule) createCustomer(req CustomerRequest) *fiber.Error {
	if err := validateCustomer(req); err != nil {
		return err
	}

	query := `
		INSERT INTO customers (type, document, name, fantasy_name, email, phone, address, observation)
		VALUES (:type, :document, :name, :fantasy_name, :email, :phone, :address, :observation)`

	if _, err := m.Db.NamedExec(query, req); err != nil {
		log.Println("inserir cliente:", err)
		return ErrCreateFailed
	}

	return nil
}

func (m *CustomerModule) updateCustomer(id string, req CustomerRequest) *fiber.Error {
	if err := validateCustomer(req); err != nil {
		return err
	}

	query := `
		UPDATE customers SET
			type = :type, document = :document, name = :name, fantasy_name = :fantasy_name,
			email = :email, phone = :phone, address = :address, observation = :observation,
			updated_at = now()
		WHERE id = :id`

	result, err := m.Db.NamedExec(query, Customer{ID: id, CustomerRequest: req})
	if err != nil {
		log.Println("atualizar cliente:", err)
		return ErrUpdateFailed
	}

	rows, err := result.RowsAffected()
	if err != nil {
		log.Println("atualizar cliente:", err)
		return ErrUpdateFailed
	}
	if rows == 0 {
		return ErrCustomerNotFound
	}

	return nil
}

func (m *CustomerModule) deleteCustomer(id string) *fiber.Error {
	result, err := m.Db.Exec(`DELETE FROM customers WHERE id = $1`, id)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == pgForeignKeyViolation {
			return ErrCustomerHasCars
		}

		log.Println("deletar cliente:", err)
		return ErrDeleteFailed
	}

	rows, err := result.RowsAffected()
	if err != nil {
		log.Println("deletar cliente:", err)
		return ErrDeleteFailed
	}
	if rows == 0 {
		return ErrCustomerNotFound
	}

	return nil
}

var sortableColumns = map[string]string{
	"name":       "name",
	"created_at": "created_at",
	"updated_at": "updated_at",
}

const maxPageSize = 100

func (m *CustomerModule) listCustomers(params ListParams) ([]Customer, *fiber.Error) {
	sortColumn, ok := sortableColumns[params.SortBy]

	if !ok {
		sortColumn = "created_at"
	}
	if params.Order != "asc" && params.Order != "desc" {
		params.Order = "asc"
	}
	if params.Page < 1 {
		params.Page = 1
	}
	if params.Limit < 1 || params.Limit > maxPageSize {
		params.Limit = 20
	}

	offset := (params.Page - 1) * params.Limit

	query := fmt.Sprintf(`
		SELECT *
		FROM customers
		WHERE name ILIKE $1 OR document LIKE $1
		ORDER BY %s %s
		LIMIT $2 OFFSET $3`, sortColumn, params.Order)

	customers := []Customer{}
	if err := m.Db.Select(&customers, query, "%"+params.Search+"%", params.Limit, offset); err != nil {
		log.Println("listar clientes:", err)
		return nil, ErrListFailed
	}

	return customers, nil
}

const maxNameResults = 20

func (m *CustomerModule) listUserNamesCustomers(prefix string) ([]CustomerName, *fiber.Error) {
	query := `
		SELECT id, name
		FROM customers
		WHERE name ILIKE $1
		ORDER BY name
		LIMIT $2`

	names := []CustomerName{}
	if err := m.Db.Select(&names, query, prefix+"%", maxNameResults); err != nil {
		log.Println("listar nomes dos clientes:", err)
		return nil, ErrListFailed
	}

	return names, nil
}

func (m *CustomerModule) getCustomerByID(id string) (Customer, *fiber.Error) {
	var customer Customer

	err := m.Db.Get(&customer, `SELECT * FROM customers WHERE id = $1`, id)
	if errors.Is(err, sql.ErrNoRows) {
		return Customer{}, ErrCustomerNotFound
	}
	if err != nil {
		log.Println("buscar cliente:", err)
		return Customer{}, ErrGetFailed
	}

	return customer, nil
}
