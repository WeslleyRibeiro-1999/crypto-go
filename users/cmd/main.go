package main

import (
	"fmt"

	"github.com/WeslleyRibeiro-1999/crypto-go/users/db"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Erro ao carregar arquivo .env: ", err)
		return
	}

	database, err := db.NewDatabase()
	if err != nil {
		panic("Falha ao conectar ao banco de dados")
	}

	fmt.Println("CHEGOU AQUI: ", database)
}
