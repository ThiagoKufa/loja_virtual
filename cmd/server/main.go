package main

import (
	"fmt"
	"os"

	"github.com/thiagokufa/loja_virtual/internal/db"
	"github.com/thiagokufa/loja_virtual/internal/server"
	"github.com/thiagokufa/loja_virtual/pkg/utils"
)

var Version = ""

func main() {
	utils.EnvLoad()
	db.Init()
	server.Init(Version)

	fmt.Printf("Server v%s running on port %s", Version, os.Getenv("PORT"))
}
