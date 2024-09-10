package main

import (
	"fmt"
	"os"
	"time"

	"github.com/thiagokufa/loja_virtual/internal/db"
	"github.com/thiagokufa/loja_virtual/internal/server"
	"github.com/thiagokufa/loja_virtual/pkg/utils"
)

var Version = ""

func main() {
	utils.EnvLoad()
	err := db.Init()
	if err != nil {
		fmt.Println("Error connecting to database: ", err)
		//esperar 10 segundos e tentar novamente
		time.Sleep(10 * time.Second)
		main()
		return
	}
	server.Init(Version)

	fmt.Printf("Server v%s running on port %s", Version, os.Getenv("PORT"))
}
