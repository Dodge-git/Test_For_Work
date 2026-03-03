package database

import (
	"fmt"
	"log"
	"time"

	"github.com/Dodge-git/Test_For_Work/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgres(cfg *config.Config) *gorm.DB {

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBName,
		cfg.DBPort,
	)

	var db *gorm.DB
	var err error

	for i := 0; i < 25; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("Connected to database")
			return db
		}

		log.Println("Database not ready, retrying...")
		time.Sleep(2 * time.Second)
	}

	log.Fatal("failed to connect database after retries:", err)
	return nil
}