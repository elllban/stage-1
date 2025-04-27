package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"stage-1/internal/service"
)

var db *gorm.DB

func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost user=task password=task dbname=task port=5432 sslmode=disable"
	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	if err := db.AutoMigrate(&service.RequestBody{}); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}

	return db, nil
}
