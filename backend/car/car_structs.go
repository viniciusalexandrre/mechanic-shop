package car

import "time"

type Car struct {
	ID string `json:"id"`
	CarRequest
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type CarRequest struct {
	CustomerID   string  `json:"customer_id" db:"customer_id"`
	LicensePlate string  `json:"license_plate" db:"license_plate"`
	Model        string  `json:"model" db:"model"`
	Brand        string  `json:"brand" db:"brand"`
	ModelYear    int     `json:"model_year" db:"model_year"`
	Color        *string `json:"color" db:"color"`
	Km           *int    `json:"km" db:"km"`
	Observation  *string `json:"observation" db:"observation"`
	Image        *string `json:"image" db:"image"`
}
