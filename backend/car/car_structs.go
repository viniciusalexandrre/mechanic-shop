package car

import "time"

type CarRequest struct {
	CustomerID   string  `json:"customer_id"`
	LicensePlate string  `json:"license_plate"`
	Model        string  `json:"model"`
	Brand        string  `json:"brand"`
	ModelYear    int     `json:"model_year"`
	Color        *string `json:"color"`
	Km           *int    `json:"km"`
	Observation  *string `json:"observation"`
	Image        *string `json:"image"`
}

type Car struct {
	ID string `json:"id"`
	CarRequest
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ListCarsParams struct {
	Search string
	SortBy string
	Order  string
	Page   int
	Limit  int
}
