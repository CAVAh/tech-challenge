package models

import (
	"fmt"

	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	CustomerId int
	Customer   Customer `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (c Order) ToDomain() entities.Order {
	fmt.Println(c.ID)
	fmt.Println(c.CustomerId)
	fmt.Println(c.Customer)

	return entities.Order{
		ID:        c.ID,
		CreatedAt: c.CreatedAt.Format("2006-01-02 15:04:05"),
		Customer:  c.Customer.ToDomain(),
		Products:  nil,
	}
}
