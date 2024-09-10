package account_controller

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/thiagokufa/loja_virtual/internal/account/account_entity"
	"github.com/thiagokufa/loja_virtual/internal/account/account_service"
)

func HandleCreate(c *gin.Context) {
	var request account_entity.CreateAccountRequestDTO
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	account := account_entity.Account{
		ID:        uuid.New().String(),
		Name:      request.Name,
		Email:     request.Email,
		Password:  request.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := account_service.Create(account)
	if err.Field != "" {
		c.JSON(err.Code, gin.H{
			"error": err.Error,
			"field": err.Field,
		})
		return
	}

	response := account_entity.CreateAccountResponseDTO{
		Message: "Account created successfully",
		Data: account_entity.DataAccount{
			ID:    account.ID,
			Name:  account.Name,
			Email: account.Email,
		},
	}
	c.JSON(200, response)
}
