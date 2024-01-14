package models

import (
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	customer Customer
	products []Product
}

func (c Order) ToDomain() entities.Order {

	return entities.Order{
		ID:        c.ID,
		CreatedAt: c.CreatedAt.Format("2006-01-02 15:04:05"),
		Customer:  c.customer.ToDomain(),
		Products:  nil,
	}
}
