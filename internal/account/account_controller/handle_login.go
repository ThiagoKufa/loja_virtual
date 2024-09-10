package account_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thiagokufa/loja_virtual/internal/account/account_entity"
	"github.com/thiagokufa/loja_virtual/internal/account/account_service"
)

func HandleLogin(c *gin.Context) {
	request := account_entity.LoginAccountRequestDTO{}
	c.BindJSON(&request)

	data, err := account_service.Login(request)
	if err.Error != "" {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, data)
}
