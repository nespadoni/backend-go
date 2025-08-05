package db

import (
	"backend-go/internal/championship"
	"backend-go/internal/match"
	"backend-go/internal/result"
	"backend-go/internal/sport"
	"backend-go/internal/tournament"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	// Carrega variáveis do .env
	if err := godotenv.Load(); err != nil {
		log.Println("Aviso: Arquivo .env não encontrado, usando variáveis de ambiente do sistema.")
	}

	dsn := os.Getenv("dsn")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Falha ao conectar no banco: %v", err)
	}

	Migrate(db)

	return db
}

func Migrate(database *gorm.DB) {

	if err := database.Debug().AutoMigrate(&championship.Championship{}, &sport.Sport{}, &tournament.Tournament{}, &match.Match{}, &result.Result{}); err != nil {
		log.Fatalf("Erro ao Executar o Migrations: %v", err)
	}

}
