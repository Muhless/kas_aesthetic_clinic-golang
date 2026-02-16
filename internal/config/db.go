package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectionDB() *gorm.DB {
	dsn := "host=localhost user=root password= dbname=kas_aesthetic_clinic port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed connect to database")
	}
	return db
}
