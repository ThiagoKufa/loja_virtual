package account

import (
	"github.com/gin-gonic/gin"
	"github.com/thiagokufa/server_task_app/internal/account/account_controller"
)

func InitRoute(r *gin.Engine) {
	accountRoute := r.Group("/account")
	{
		accountRoute.POST("/create", account_controller.HandleCreate)
		accountRoute.POST("/login", account_controller.HandleLogin)
	}
}
