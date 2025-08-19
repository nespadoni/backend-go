package db

import (
	"backend-go/config"
	"backend-go/internal/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

func InitDB(cfg config.DatabaseConfig) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port)

	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		log.Fatalf("Erro ao conectar com o banco de dados: %v", err)
	}

	if err := autoMigrate(db); err != nil {
		log.Fatalf("Erro ao executar migrations: %v", err)
	}

	log.Println("Conectado ao banco de dados com sucesso!")
	return db
}

func autoMigrate(db *gorm.DB) error {
	// Ordem das migrations é importante por causa das foreign keys
	// Migra primeiro as tabelas sem dependências, depois as dependentes

	log.Println("Executando migrations...")

	return db.AutoMigrate(
		// Tabelas base (sem foreign keys)
		&models.University{},
		&models.Role{},
		&models.Sport{},
		&models.Championship{},

		// Tabelas com uma dependência
		&models.User{},       // depende de University
		&models.Athletic{},   // depende de University
		&models.Position{},   // depende de Sport
		&models.Tournament{}, // depende de Championship e Sport

		// Tabelas com múltiplas dependências
		&models.Team{},   // depende de Athletic
		&models.News{},   // depende de Athletic
		&models.Follow{}, // depende de User
		&models.Player{}, // depende de Team e User
		&models.Match{},  // depende de Tournament

		// Tabelas de relacionamento e estatísticas
		&models.UserRoleAthletic{}, // depende de User, Role e Athletic
		&models.Result{},           // depende de Match
		&models.Lineup{},           // depende de Match e Player
		&models.PlayerStats{},      // depende de Player e Match
		&models.TournamentMatch{},  // depende de Tournament e Match
		&models.Notification{},     // depende de User
	)
}
