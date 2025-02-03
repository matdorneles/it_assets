package models

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

// connects to PostgreSQL, panics if an error is returned
func ConnectDB() {
	connectionString := "host=localhost user=root password=root dbname=gestao_ti port=5432 sslmode=disable timezone=America/Sao_Paulo"
	DB, err = gorm.Open(postgres.Open(connectionString))
	if err != nil {
		log.Panic("Error connecting to database")
	}
	log.Println("Connected to PostgreSQL")
	DB.AutoMigrate(&Computer{})
}
