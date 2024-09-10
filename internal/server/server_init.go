package server

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/thiagokufa/server_task_app/internal/route"
)

func Init() {
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	r := gin.Default()
	route.Init(r)

	r.Run(port)
}
