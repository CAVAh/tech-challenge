package models

import (
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	CustomerID uint
	Customer   Customer `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Products   []OrderProduct
	Status     string
}

func (o Order) ToDomain() entities.Order {
	var products []entities.ProductInsideOrder
	//TODO: tem que pegar quantity e observation
	//for _, p := range o.Products {
	//	products = append(products, entities.ProductInsideOrder{Product: p.ToDomain()})
	//}

	return entities.Order{
		ID:        o.ID,
		CreatedAt: o.CreatedAt.Format("2006-01-02 15:04:05"),
		Customer:  o.Customer.ToDomain(),
		Status:    o.Status,
		Products:  products,
	}
}
