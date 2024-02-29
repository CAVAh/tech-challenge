package gateways

import "github.com/CAVAh/api-tech-challenge/src/core/domain/entities"

type ProductCategoryRepository interface {
	FindAll() ([]entities.ProductCategory, error)
	FindById(id uint) (*entities.ProductCategory, error)
}
