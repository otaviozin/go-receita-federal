package main

import (
	"log"

	"github.com/otaviozin/go-receita-federal/database"
	"github.com/otaviozin/go-receita-federal/types"
)

func main() {
	db := database.DBConnection()

	err := db.AutoMigrate(&types.Empresa{})
	if err != nil {
		log.Fatalf("Erro ao fazer AutoMigrate: %v", err)
	}

	log.Println("Banco de dados migrado com sucesso!")
}
