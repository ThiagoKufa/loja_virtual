package main

import (
	"github.com/thiagokufa/server_task_app/internal/db"
	"github.com/thiagokufa/server_task_app/internal/server"
	"github.com/thiagokufa/server_task_app/pkg/utils"
)

func main() {
	utils.EnvLoad()
	db.Init()
	server.Init()

}
