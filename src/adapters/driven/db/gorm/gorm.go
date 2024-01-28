package gorm

import (
	"fmt"
	"github.com/CAVAh/api-tech-challenge/src/adapters/driven/db/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var (
	DB *gorm.DB
)

func ConnectDB() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	conectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=America/Fortaleza",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
	DB, err = gorm.Open(postgres.Open(conectionString))

	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	productCategories := []models.ProductCategory{
		{Description: "Lanche"},
		{Description: "Acompanhamento"},
		{Description: "Bebida"},
		{Description: "Sobremesa"},
	}

	if !DB.Migrator().HasTable("product_categories") {
		DB.Migrator().CreateTable(&productCategories)
		DB.Create(&productCategories)
	}

	DB.AutoMigrate(&models.Customer{}, &models.Product{}, &models.Order{}, &models.OrderProduct{})

	if !DB.Migrator().HasColumn(&models.OrderProduct{}, "Quantity") {
		DB.Migrator().AddColumn(&models.OrderProduct{}, "Quantity")
		DB.Migrator().AddColumn(&models.OrderProduct{}, "Observation")
	}
}
