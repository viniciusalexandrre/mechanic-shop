package customer

import "time"

type ListParams struct {
	Search string
	SortBy string
	Order  string
	Page   int
	Limit  int
}

type CustomerRequest struct {
	Type        string  `json:"type" db:"type"`
	Document    string  `json:"document" db:"document"`
	Name        string  `json:"name" db:"name"`
	FantasyName *string `json:"fantasy_name" db:"fantasy_name"`
	Email       *string `json:"email" db:"email"`
	Phone       string  `json:"phone" db:"phone"`
	Address     *string `json:"address" db:"address"`
	Observation *string `json:"observation" db:"observation"`
}

type Customer struct {
	ID string `json:"id" db:"id"`
	CustomerRequest
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type CustomerName struct {
	ID   string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
