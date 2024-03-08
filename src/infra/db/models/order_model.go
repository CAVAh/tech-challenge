package models

import (
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
	"github.com/CAVAh/api-tech-challenge/src/utils"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	CustomerID    uint
	Customer      Customer `gorm:"foreignKey:CustomerID;references:ID"`
	Products      []OrderProduct
	Status        string
	PaymentStatus string
}

func (o Order) ToDomain(products []entities.ProductInsideOrder) entities.Order {
	return entities.Order{
		ID:            o.ID,
		Customer:      o.Customer.ToDomain(),
		Products:      products,
		Status:        o.Status,
		PaymentStatus: o.PaymentStatus,
		CreatedAt:     o.CreatedAt.Format(utils.CompleteEnglishDateFormat),
		UpdatedAt:     o.UpdatedAt.Format(utils.CompleteEnglishDateFormat),
	}
}
