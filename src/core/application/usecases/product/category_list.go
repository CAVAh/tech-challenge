package product

import (
	"github.com/CAVAh/api-tech-challenge/src/core/application/ports/repositories"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
)

type ListCategory struct {
	repository repositories.ProductCategoryRepository
}

func BuildListCategory(repository repositories.ProductCategoryRepository) *ListCategory {
	return &ListCategory{repository: repository}
}

func (p ListCategory) Execute() ([]entities.ProductCategory, error) {
	return p.repository.FindAll()
}
