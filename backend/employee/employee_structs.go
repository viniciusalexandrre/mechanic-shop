package employee

import "time"

type Employee struct {
	ID string `json:"id"`
	EmployeeRequest
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type EmployeeRequest struct {
	Name    string `json:"name"`
	HiredAt string `json:"hired_at"`
	Cpf     string `json:"cpf"`
}

type ListEmployeesParams struct {
	Search string
	SortBy string
	Order  string
	Page   int
	Limit  int
}
