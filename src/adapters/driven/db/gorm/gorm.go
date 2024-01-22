package gorm

import (
	"github.com/CAVAh/api-tech-challenge/src/adapters/driven/db/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
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
  
  DB.AutoMigrate(&models.Customer{})
	DB.AutoMigrate(&models.Product{})
	DB.AutoMigrate(&models.Order{})
	DB.AutoMigrate(&models.OrderProduct{})
}
