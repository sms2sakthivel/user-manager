package model

import (
	"gorm.io/gorm"
)

// Migrate the database schema to match the models.
func Automigrate(DB *gorm.DB) error {
	return DB.AutoMigrate(&User{})
}
