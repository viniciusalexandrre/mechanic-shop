package car

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func (m *CarModule) createCar(req CarRequest) *fiber.Error {
	if err := validateCar(req); err != nil {
		return err
	}

	query := `
		INSERT INTO customers(customer_id, license_plate, model, brand, model_year, color, km, observation, image)
		VALUES (:customer_id, :license_plate, :model, :brand, :model_year, :color, :km, :observation, :image )
	`

	if _, err := m.Db.NamedExec(query, req); err != nil {
		log.Println("Inserir Carro:", err)
		return ErrCreateFailed
	}

	return nil
}

func (m *CarModule) updateCar(id string, req CarRequest) *fiber.Error {
	if err := validateCar(req); err != nil {
		return err
	}

	query := `
		UPDATE car SET
		license_plate = :license_plate, model = :model, brand = :brand, model_year = :model_year, color = :color, km = :km, observation = :observation, image = :image, updated_at = now()
		WHERE customer_id = : customer_id
	`

	result, err := m.Db.NamedExec(query, Car{ID: id, CarRequest: req})

	if err != nil {
		log.Println("Atualizar cliente:", err)
		return ErrUpdateFailed
	}

	rows, err := result.RowsAffected()

	switch {
	case err != nil:
		return ErrUpdateFailed
	case rows == 0:
		return ErrIdNotFound
	}

	return nil
}

// func (m *CarModule) deleteCar(id string) *fiber.Error {
// 	result, err := m.Db.Exec(`DELETE FROM car WHERE id = $1`, id)
// 	if err != nil {

// 	}
// }
