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

func (c Order) ToDomain() entities.Order {

	return entities.Order{
		ID:        c.ID,
		CreatedAt: c.CreatedAt.Format("2006-01-02 15:04:05"),
		Customer:  c.Customer.ToDomain(),
		Products:  nil,
	}
}
