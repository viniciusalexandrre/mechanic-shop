package employee

import "time"

type Employee struct {
	ID string `json:"id" db:"id"`
	EmployeeRequest
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type EmployeeRequest struct {
	Name  string `json:"name" db:"name"`
	Shift string `json:"shift" db:"shift"`
	CPF   string `json:"cpf" db:"cpf"`
}

type ListEmployeesParams struct {
	Search string
	SortBy string
	Order  string
	Page   int
	Limit  int
}
