package models

import "github.com/CAVAh/api-tech-challenge/src/core/domain/entities"

type OrderProduct struct {
	OrderID     uint `gorm:"primaryKey"`
	ProductID   uint `gorm:"primaryKey"`
	Order       Order
	Product     Product
	Quantity    int
	Observation string
}

func (c OrderProduct) ToDomain() entities.OrderProduct {
	return entities.OrderProduct{
		ProductID:   c.ProductID,
		Quantity:    c.Quantity,
		Observation: c.Observation,
	}
}
