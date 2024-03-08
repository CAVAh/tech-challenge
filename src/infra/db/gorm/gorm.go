package gorm

import (
	"fmt"
	models2 "github.com/CAVAh/api-tech-challenge/src/infra/db/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDB() {

	conectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=America/Fortaleza",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))

	DB, err = gorm.Open(postgres.Open(conectionString))

	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	productCategories := []models2.ProductCategory{
		{Description: "Lanche"},
		{Description: "Acompanhamento"},
		{Description: "Bebida"},
		{Description: "Sobremesa"},
	}

	if !DB.Migrator().HasTable("product_categories") {
		DB.Migrator().CreateTable(&productCategories)
		DB.Create(&productCategories)
	}

	DB.AutoMigrate(&models2.Customer{}, &models2.Product{}, &models2.Order{}, &models2.OrderProduct{})

	if !DB.Migrator().HasColumn(&models2.OrderProduct{}, "Quantity") {
		DB.Migrator().AddColumn(&models2.OrderProduct{}, "Quantity")
		DB.Migrator().AddColumn(&models2.OrderProduct{}, "Observation")
	}
}
