package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// -------------------------------PUBLIC--------------------------------- //

func Connect() (*gorm.DB, error) {
	// Database connection parameters
	dsn := "user=postgres password=postgres dbname=godb host=localhost port=5432 sslmode=disable"
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
