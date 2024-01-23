package models

import (
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	CustomerId int
	Customer   Customer  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Products   []Product `gorm:"many2many:order_products;"`
}

func (o Order) ToDomain() entities.Order {
	var products []entities.Product
	for _, p := range o.Products {
		products = append(products, p.ToDomain())
	}

	return entities.Order{
		ID:        o.ID,
		CreatedAt: o.CreatedAt.Format("2006-01-02 15:04:05"),
		Customer:  o.Customer.ToDomain(),
		Products:  products,
	}
}
