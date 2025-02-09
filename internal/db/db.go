package db

import (
	"hivelock/internal/models"
	"log"

	"gorm.io/driver/sqlite" // Use "gorm.io/driver/postgres" for PostgreSQL
	"gorm.io/gorm"
)

// DB is a global database connection
var DB *gorm.DB

func InitDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("secrets.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto-migrate the Secret model
	DB.AutoMigrate(&models.Secret{})
}
