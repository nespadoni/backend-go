package db

import (
	"backend-go/internal/models"
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

	if err := database.Debug().AutoMigrate(&models.Championship{},
		&models.Sport{},
		&models.Tournament{},
		&models.Role{},
		&models.University{},
		&models.User{},
		&models.Userchampionship{},
		&models.Match{},
		&models.Result{}); err != nil {
		log.Fatalf("Erro ao Executar o Migrations: %v", err)
	}

	var countRoles int64
	var countUniversities int64

	database.Model(&models.Role{}).Count(&countRoles)
	database.Model(&models.University{}).Count(&countUniversities)

	if countRoles == 0 {
		adminRole := models.Role{
			Name: "ADM", Description: "MASTER ROLE", Admin: true,
		}
		database.Create(&adminRole)

		userRole := models.Role{
			Name: "USER", Description: "NORMAL ROLE", Admin: false,
		}
		database.Create(&userRole)
	}

	if countUniversities == 0 {
		univag := models.University{
			Name: "UNIVAG",
		}
		database.Create(&univag)

		unic := models.University{
			Name: "UNIC",
		}
		database.Create(&unic)

		ufmt := models.University{
			Name: "UFMT",
		}
		database.Create(&ufmt)
	}
}
