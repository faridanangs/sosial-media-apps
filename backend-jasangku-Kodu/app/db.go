package app

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Name     string
	Password string
	Schema   string
}

func ConnectToDatabase(dbConfig DbConfig) *gorm.DB {
	dns := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable TimeZone=Asia/Jakarta",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name,
		dbConfig.User,
		dbConfig.Password,
	)
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{
		Logger:      logger.Default.LogMode(logger.Info),
		PrepareStmt: true,
	})
	if err != nil {
		log.Fatalf("Could not connect to database\n, error: %v", err)
	}
	schema := fmt.Sprintf("set search_path to %s", dbConfig.Schema)
	db.Exec(schema)

	return db
}
