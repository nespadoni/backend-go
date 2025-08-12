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
		&models.Athletic{},
		&models.UserRoleAthletic{},
		&models.Match{},
		&models.Result{}); err != nil {
		log.Fatalf("Erro ao Executar o Migrations: %v", err)
	}
	//Cria os principais Registros
	mainData(database)
}

func mainData(database *gorm.DB) {
	var countRoles int64
	var countUniversities int64
	var countUsers int64

	database.Model(&models.Role{}).Count(&countRoles)
	database.Model(&models.University{}).Count(&countUniversities)
	database.Model(&models.User{}).Count(&countUsers)

	if countRoles == 0 && countUsers == 0 {
		//Role ADM
		adminRole := models.Role{
			Name: "ADM", Description: "MASTER ROLE", Admin: true,
		}

		result := database.Create(&adminRole)
		if result.Error != nil {
			log.Fatalf("Erro ao criar o registro: %v", result.Error)
		}

		//Role NORMAL
		userRole := models.Role{
			Name: "USER", Description: "NORMAL ROLE", Admin: false,
		}

		result = database.Create(&userRole)
		if result.Error != nil {
			log.Fatalf("Erro ao criar o registro: %v", result.Error)
		}

		//USUÁRIO
		user := models.User{
			Name:      "admin",
			Email:     "admin",
			Password:  "admin",
			Telephone: "99999999"}

		result = database.Create(&user)
		if result.Error != nil {
			log.Fatalf("Erro ao criar o registro: %v", result.Error)
		}

		//USUARIO ROLE, SEM ATLETICA PORQUE SERA O USUARIO MASTER
		var athleticID *int
		user_role_athletic := models.UserRoleAthletic{
			UserID:     user.ID,
			RoleID:     adminRole.ID,
			AthleticID: athleticID,
		}

		result = database.Create(&user_role_athletic)
		if result.Error != nil {
			log.Fatalf("Erro ao criar o registro: %v", result.Error)
		}
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

		result := database.Create(&ufmt)
		if result.Error != nil {
			log.Fatalf("Erro ao criar o registro: %v", result.Error)
		}
	}
}
