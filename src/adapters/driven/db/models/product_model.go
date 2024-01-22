package models

import (
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string `gorm:"unique;index"`
	Price       float64
	Description string
	CategoryID  uint
}

func (c Product) ToDomain() entities.Product {
	return entities.Product{
		ID:          c.ID,
		Name:        c.Name,
		Price:       c.Price,
		Description: c.Description,
		CategoryId:  c.CategoryID,
		CreatedAt:   c.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}
