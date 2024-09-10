package utils

import (
	"fmt"

	"github.com/joho/godotenv"
)

func EnvLoad() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Erro ao carregar o arquivo .env: %v", err)
	}
	fmt.Println("Arquivo .env carregado com sucesso")
}
