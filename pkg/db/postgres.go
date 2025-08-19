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

	// IMPORTANTE: Executar seed das roles após as migrations
	if err := seeders.SeedRoles(db); err != nil {
		log.Fatalf("Erro ao executar seed das roles: %v", err)
	}

	log.Println("Conectado ao banco de dados com sucesso!")
	return db
}

func autoMigrate(db *gorm.DB) error {
	log.Println("Executando migrations...")

	return db.AutoMigrate(
		// Tabelas base
		&models.University{},
		&models.Role{},
		&models.Sport{},

		// Tabelas com dependências
		&models.User{},
		&models.Athletic{},
		&models.Position{},

		// Sistema social
		&models.Follow{},
		&models.Like{},
		&models.Comment{},

		// Conteúdo
		&models.News{},
		&models.Championship{},
		&models.Tournament{},

		// Times e jogadores
		&models.Team{},
		&models.Player{},
		&models.Match{},

		// Estatísticas
		&models.Result{},
		&models.Lineup{},
		&models.PlayerStats{},
		&models.TournamentMatch{},

		// Sistema de permissões
		&models.UserRoleAthletic{},
		&models.Notification{},
	)
}
