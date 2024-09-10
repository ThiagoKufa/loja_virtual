package account_validate

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/thiagokufa/server_task_app/internal/account/account_entity"
	"github.com/thiagokufa/server_task_app/pkg/utils"
)

func RegisterValidation(a *account_entity.Account) utils.ErrorResponse {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(a.Email) {
		err := utils.ErrorResponse{
			Field: "email",
			Error: "invalid email",
			Code:  http.StatusBadRequest,
		}
		return err
	}

	if len(a.Name) < 3 {
		err := utils.ErrorResponse{
			Field: "name",
			Error: "name is required and must be at least 3 characters",
			Code:  http.StatusBadRequest,
		}
		return err
	}

	// Validação da senha
	passwordErrors := []string{}
	if len(a.Password) < 8 {
		passwordErrors = append(passwordErrors, "password must contain at least 8 characters")
	}
	if !strings.ContainsAny(a.Password, "!@#$%^&*") {
		passwordErrors = append(passwordErrors, "password must contain at least one special character")
	}
	if !strings.ContainsAny(a.Password, "0123456789") {
		passwordErrors = append(passwordErrors, "password must contain at least one number")
	}
	if !strings.ContainsAny(a.Password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		passwordErrors = append(passwordErrors, "password must contain at least one uppercase letter")
	}
	if !strings.ContainsAny(a.Password, "abcdefghijklmnopqrstuvwxyz") {
		passwordErrors = append(passwordErrors, "password must contain at least one lowercase letter")
	}

	if len(passwordErrors) > 0 {
		err := utils.ErrorResponse{
			Field: "password",
			Error: strings.Join(passwordErrors, "; "),
			Code:  http.StatusBadRequest,
		}
		return err
	}

	return utils.ErrorResponse{}
}
