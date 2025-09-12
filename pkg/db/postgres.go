package db

import (
	"backend-go/config"
	"backend-go/internal/models"
	"backend-go/internal/seeders"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB(cfg config.DatabaseConfig) *gorm.DB {
	var dsn string
	// Prioriza a DATABASE_URL se ela for fornecida (ideal para o Railway)
	if cfg.URL != "" {
		dsn = cfg.URL
	} else {
		// Monta a DSN para o ambiente local se a URL não estiver presente
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port)
	}

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
		// === TABELAS BASE (sem dependências) ===
		&models.University{},
		&models.Sport{},

		// === TABELAS COM DEPENDÊNCIAS SIMPLES ===
		&models.Course{},   // depende de University
		&models.Position{}, // depende de Sport
		&models.User{},     // depende de University e Course
		&models.Role{},     // pode depender de Athletic (circular, mas opcional)

		// === TABELAS DE NEGÓCIO ===
		&models.Athletic{}, // depende de University e User (creator)
		&models.Team{},     // depende de Athletic

		// === CONTEÚDO ===
		&models.News{},         // depende de Athletic e User (author)
		&models.Championship{}, // depende de Athletic
		&models.Tournament{},   // depende de Championship e Sport

		// === JOGADORES E PARTIDAS ===
		&models.Player{}, // depende de Team, Position, User
		&models.Match{},  // depende de Tournament e Team

		// === ESTATÍSTICAS E RESULTADOS ===
		&models.Result{},      // depende de Match
		&models.Lineup{},      // depende de Match e Player
		&models.PlayerStats{}, // depende de Player e Match

		// === SISTEMA SOCIAL ===
		&models.Follow{},  // depende de User e Athletic
		&models.Like{},    // depende de User (polimórfico)
		&models.Comment{}, // depende de User (polimórfico e self-referencing)

		// === SISTEMA DE PERMISSÕES ===
		&models.UserRoleAthletic{}, // depende de User, Role, Athletic

		// === SISTEMA DE NOTIFICAÇÕES ===
		&models.Notification{}, // depende de User e Athletic

		// === TABELAS ASSOCIATIVAS ===
		&models.TournamentMatch{}, // depende de Tournament e Match
	)
}
