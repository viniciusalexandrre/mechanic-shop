package customer

import "time"

type CustomerRequest struct {
	Type        string  `json:"type"`
	Document    string  `json:"document"`
	Name        string  `json:"name"`
	FantasyName *string `json:"fantasy_name"`
	Email       *string `json:"email"`
	Phone       string  `json:"phone"`
	Address     *string `json:"address"`
	Observation *string `json:"observation"`
}

type Customer struct {
	ID string `json:"id"`
	CustomerRequest
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ListCustomersParams struct {
	Search string
	SortBy string
	Order  string
	Page   int
	Limit  int
}
