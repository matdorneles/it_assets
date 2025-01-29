package config

import (
	"log"

	"github.com/matdorneles/gestao_ti/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDB() {
	connectionString := "host=localhost user=root password=root dbname=gestao_ti port=5432 sslmode=disable timezone=America/Sao_Paulo"
	DB, err = gorm.Open(postgres.Open(connectionString))
	if err != nil {
		log.Panic("Error connecting to database")
	}
	log.Println("Connected to PostgreSQL")
	DB.AutoMigrate(&models.Computer{})
}
