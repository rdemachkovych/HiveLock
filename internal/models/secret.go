package models

import (
	"time"

	"gorm.io/gorm"
)

type Secret struct {
	ID       uint   `gorm:"primaryKey"`
	Key      string `gorm:"uniqueIndex"`
	Value    []byte // Encrypted secret
	CreateAt time.Time
}

// MigrateSecrets sets up the database schema for secrets
func MigrateSecrets(db *gorm.DB) error {
	return db.AutoMigrate(&Secret{})
}
