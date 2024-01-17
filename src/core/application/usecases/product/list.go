package product

import (
	"github.com/CAVAh/api-tech-challenge/src/core/application/ports/repositories"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
)

type List struct {
	repository repositories.ProductRepository
}

func BuildList(repository repositories.ProductRepository) *List {
	return &List{repository: repository}
}

func (p List) Execute(categoryId uint) ([]entities.Product, error) {

	if categoryId == 0 {
		return p.repository.FindAll()
	}

	return p.repository.FindByCategoryId(categoryId)
}
