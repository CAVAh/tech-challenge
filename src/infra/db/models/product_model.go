package models

import (
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
	"github.com/CAVAh/api-tech-challenge/src/utils"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string `gorm:"unique;index"`
	Price       float64
	Description string
	CategoryID  uint
	Category    ProductCategory `gorm:"foreignKey:CategoryID;references:ID"`
}

func (c Product) ToDomain() entities.Product {
	return entities.Product{
		ID:          c.ID,
		Name:        c.Name,
		Price:       c.Price,
		Description: c.Description,
		CategoryId:  c.CategoryID,
		CreatedAt:   c.CreatedAt.Format(utils.CompleteEnglishDateFormat),
	}
}

func (p *Product) PatchFields(name string, price float64, description string, categoryId uint) {
	p.Name = name
	p.Price = price
	p.Description = description
	p.CategoryID = categoryId
}
