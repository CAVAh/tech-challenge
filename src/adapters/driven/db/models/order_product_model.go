package models

import (
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
	"gorm.io/gorm"
)

type OrderProduct struct {
	gorm.Model
	OrderID     uint
	ProductID   uint
	Quantity    int
	Observation string
	Product     Product
}

func (c OrderProduct) ToDomain() entities.OrderProduct {
	return entities.OrderProduct{
		ProductID:   c.ProductID,
		Quantity:    c.Quantity,
		Observation: c.Observation,
	}
}
