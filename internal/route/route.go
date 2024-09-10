package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thiagokufa/loja_virtual/internal/account"
)

func Init(r *gin.Engine, version string) {
	r.GET("/versions", func(c *gin.Context) {
		c.String(http.StatusOK, version)
	})
	account.InitRoute(r)
}
