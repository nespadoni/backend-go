package main

import (
	"backend-go/config"
	"backend-go/pkg/db"
	"backend-go/routes"
	"log"
)

func main() {
	// 1. Carrega as configurações
	cfg := config.Load()
	log.Printf("Iniciando servidor na porta: %s", cfg.Port)

	// 2. Inicializa o banco de dados com as configurações
	database := db.InitDB(cfg.Database)

	// 3. Inicializa o router passando database e config
	routes.InitRouter(database, cfg)
}
