package car

import "github.com/gofiber/fiber/v2"

func validateCar(req CarRequest) *fiber.Error {
	switch {
	case req.CustomerID == "":
		return ErrCustomerIDRequired
	case req.LicensePlate == "":
		return ErrLicensePlateRequired
	case req.Model == "":
		return ErrModelRequired
	}

	return nil
}
