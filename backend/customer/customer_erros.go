package customer

import "github.com/gofiber/fiber/v2"

var (
	ErrNameRequired = fiber.NewError(fiber.StatusBadRequest, "Nome é obrigatório")
	ErrInvalidType  = fiber.NewError(fiber.StatusBadRequest, "Type deve ser PF ou PJ")
	ErrCreateFailed = fiber.NewError(fiber.StatusInternalServerError, "Erro ao cadastrar o cliente")

	ErrInvalidID        = fiber.NewError(fiber.StatusBadRequest, "ID do cliente inválido")
	ErrCustomerNotFound = fiber.NewError(fiber.StatusNotFound, "Cliente não encontrado")
	ErrUpdateFailed     = fiber.NewError(fiber.StatusInternalServerError, "Erro ao atualizar o cliente")
	ErrListFailed       = fiber.NewError(fiber.StatusInternalServerError, "Erro ao listar os clientes")
	ErrDeleteFailed     = fiber.NewError(fiber.StatusInternalServerError, "Erro ao deletar o cliente")
	ErrCustomerHasCars  = fiber.NewError(fiber.StatusConflict, "Cliente possui carros cadastrados e não pode ser deletado")
)

func ErrDocumentExists(doc string) *fiber.Error {
	return fiber.NewError(fiber.StatusConflict, "documento "+doc+" já cadastrado")
}
