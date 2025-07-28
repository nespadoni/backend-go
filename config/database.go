package config

import (
	"backend-go/handler"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "athleticmanager"
)

func InitDB() *gorm.DB {
	connectionString := fmt.Sprintf("host=%s port=%d user-service=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := gorm.Open(
		postgres.Open(connectionString),
		&gorm.Config{},
	)

	handler.New("DB_CONN_ERROR", "Error connecting to the database: ", err)

	return db
}
