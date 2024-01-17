package repositories

import "github.com/CAVAh/api-tech-challenge/src/core/domain/entities"

type ProductCategoryRepository interface {
	FindAll() ([]entities.ProductCategory, error)
}
