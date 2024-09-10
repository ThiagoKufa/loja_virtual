package account_service

import (
	"fmt"
	"net/http"

	"github.com/thiagokufa/loja_virtual/internal/account/account_entity"
	"github.com/thiagokufa/loja_virtual/internal/account/account_repository"
	"github.com/thiagokufa/loja_virtual/internal/account/account_validate"
	"github.com/thiagokufa/loja_virtual/pkg/utils"
)

// Create cria uma nova conta de usuário.
// Recebe um objeto Account e retorna uma ErrorResponse.
// Se a validação falhar ou o email já existir, retorna um erro apropriado.
func Create(a account_entity.Account) utils.ErrorResponse {
	// Valida os dados da conta usando a função de validação registrada.
	validate := account_validate.RegisterValidation(&a)
	fmt.Println(validate)

	// Se a validação falhar, retorna o erro correspondente.
	if validate.Error != "" || validate.Code != 0 {
		return validate
	}

	// Verifica se o email já está em uso.
	emailExists, _ := account_repository.FindByEmail(a.Email)
	if emailExists.Email != "" {
		return utils.ErrorResponse{
			Field: "email",
			Error: "email already exists",
			Code:  http.StatusBadRequest,
		}
	}

	// Hash da senha antes de armazená-la.
	a.Password = utils.HashPassword(a.Password)

	// Tenta criar a conta no repositório.
	err := account_repository.Create(a)
	if err != nil {
		return utils.ErrorResponse{
			Field: "error",
			Error: err.Error(),
			Code:  http.StatusInternalServerError,
		}
	}

	// Retorna uma resposta de erro vazia se tudo ocorrer bem.
	return utils.ErrorResponse{}
}
