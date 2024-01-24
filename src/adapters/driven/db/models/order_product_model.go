package models

import "github.com/CAVAh/api-tech-challenge/src/core/domain/entities"

type OrderProduct struct {
	ID          uint `gorm:"primaryKey"`
	OrderID     uint
	ProductID   uint
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
