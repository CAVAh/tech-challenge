package models

import (
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name              string
	Price             float64
	Description       string
	ProductCategoryID int
	Category          ProductCategory `gorm:"foreignKey:ProductCategoryID;references:ID"`
}

func (p Product) ToDomain() entities.Product {

	return *entities.NewProduct(int(p.ID),
		p.Name,
		p.Price,
		p.Description,
		p.ProductCategoryID,
		p.CreatedAt.Format("2006-01-02 15:04:05"))
}

func (p *Product) PatchFields(name string, price float64, description string, categoryId int) {
	p.Name = name
	p.Price = price
	p.Description = description
	p.ProductCategoryID = categoryId
}
