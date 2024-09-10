package route

import (
	"github.com/gin-gonic/gin"
	"github.com/thiagokufa/loja_virtual/internal/account"
)

func Init(r *gin.Engine) {
	account.InitRoute(r)
}
