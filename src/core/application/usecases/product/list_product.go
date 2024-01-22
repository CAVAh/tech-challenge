package usecases

import (
	"github.com/CAVAh/api-tech-challenge/src/core/application/ports/repositories"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
)

type ListProductUsecase struct {
	repository repositories.ProductRepository
}

func BuildListProductUsecase(repository repositories.ProductRepository) *ListProductUsecase {
	return &ListProductUsecase{repository: repository}
}

func (p ListProductUsecase) Execute(categoryId int) ([]entities.Product, error) {

	if categoryId == 0 {
		return p.repository.FindAll()
	}

	return p.repository.FindByCategoryId(categoryId)
}
