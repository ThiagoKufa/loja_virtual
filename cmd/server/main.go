package main

import (
	"github.com/thiagokufa/loja_virtual/internal/db"
	"github.com/thiagokufa/loja_virtual/internal/server"
	"github.com/thiagokufa/loja_virtual/pkg/utils"
)

func main() {
	utils.EnvLoad()
	db.Init()
	server.Init()

}
