package car

import "github.com/gofiber/fiber/v2"

var (
	ErrInvalidBody          = fiber.NewError(fiber.StatusBadRequest, "JSON inválido")
	ErrCustomerIDRequired   = fiber.NewError(fiber.StatusBadRequest, "Cliente é obrigatório")
	ErrInvalidCustomerID    = fiber.NewError(fiber.StatusBadRequest, "ID do cliente inválido")
	ErrLicensePlateRequired = fiber.NewError(fiber.StatusBadRequest, "Placa é obrigatória")
	ErrInvalidLicensePlate  = fiber.NewError(fiber.StatusBadRequest, "Placa deve estar no formato ABC1234 ou ABC1D23")
	ErrBrandRequired        = fiber.NewError(fiber.StatusBadRequest, "Marca é obrigatória")
	ErrBrandTooLong         = fiber.NewError(fiber.StatusBadRequest, "Marca deve ter no máximo 50 caracteres")
	ErrModelRequired        = fiber.NewError(fiber.StatusBadRequest, "Modelo é obrigatório")
	ErrModelTooLong         = fiber.NewError(fiber.StatusBadRequest, "Modelo deve ter no máximo 100 caracteres")
	ErrInvalidModelYear     = fiber.NewError(fiber.StatusBadRequest, "Ano do modelo inválido")
	ErrColorTooLong         = fiber.NewError(fiber.StatusBadRequest, "Cor deve ter no máximo 30 caracteres")
	ErrNegativeKm           = fiber.NewError(fiber.StatusBadRequest, "Quilometragem não pode ser negativa")

	ErrCustomerNotFound   = fiber.NewError(fiber.StatusNotFound, "Cliente não encontrado")
	ErrLicensePlateExists = fiber.NewError(fiber.StatusConflict, "Já existe um carro com essa placa")
	ErrCreateFailed       = fiber.NewError(fiber.StatusInternalServerError, "Erro ao cadastrar o carro")
	ErrUpdateFailed       = fiber.NewError(fiber.StatusInternalServerError, "Erro na atualização das informações do carro")
	ErrListFailed         = fiber.NewError(fiber.StatusInternalServerError, "Erro na listagem das informações")
	ErrIdNotFound         = fiber.NewError(fiber.StatusInternalServerError, "Carro não encontrado")
)
