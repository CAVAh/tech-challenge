package gorm

import (
	"log"

	"github.com/CAVAh/api-tech-challenge/src/adapters/driven/db/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDB() {
	conectionString := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable TimeZone=America/Fortaleza"
	DB, err = gorm.Open(postgres.Open(conectionString))

	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	DB.AutoMigrate(&models.Customer{}, &models.Order{})
}
