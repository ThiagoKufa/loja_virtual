package route

import (
	"github.com/gin-gonic/gin"
	"github.com/thiagokufa/server_task_app/internal/account"
)

func Init(r *gin.Engine) {
	account.InitRoute(r)
}
