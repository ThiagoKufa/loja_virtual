package account_service

import (
	"net/http"

	"github.com/thiagokufa/loja_virtual/internal/account/account_entity"
	"github.com/thiagokufa/loja_virtual/internal/account/account_repository"
	"github.com/thiagokufa/loja_virtual/pkg/utils"
)

func Login(a account_entity.LoginAccountRequestDTO) (account_entity.LoginAccountResponseDTO, utils.ErrorResponse) {
	account, err := account_repository.FindByEmail(a.Email)

	if err != nil {
		return account_entity.LoginAccountResponseDTO{}, utils.ErrorResponse{
			Field: "email",
			Error: "Email not found",
			Code:  http.StatusNotFound,
		}
	}

	validPassword := utils.ComparePassword(account.Password, a.Password)
	if !validPassword {
		return account_entity.LoginAccountResponseDTO{}, utils.ErrorResponse{
			Field: "password",
			Error: "Password does not match",
			Code:  http.StatusBadRequest,
		}
	}

	token, err := utils.GenerateToken(account.ID)
	if err != nil {
		return account_entity.LoginAccountResponseDTO{}, utils.ErrorResponse{
			Field: "token",
			Error: "Error generating token",
			Code:  http.StatusInternalServerError,
		}
	}

	return account_entity.LoginAccountResponseDTO{
		Data: account_entity.DataAccount{
			ID:    account.ID,
			Name:  account.Name,
			Email: account.Email,
		},
		Message:  "Login successful",
		JWTToken: token,
	}, utils.ErrorResponse{}
}
