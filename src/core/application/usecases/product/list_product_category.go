package usecases

import (
	"github.com/CAVAh/api-tech-challenge/src/core/application/ports/repositories"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
)

type ListProductCategoryUsecase struct {
	repository repositories.ProductCategoryRepository
}

func BuildListProductCategoryUsecase(repository repositories.ProductCategoryRepository) *ListProductCategoryUsecase {
	return &ListProductCategoryUsecase{repository: repository}
}

func (p ListProductCategoryUsecase) Execute() ([]entities.ProductCategory, error) {
	return p.repository.FindAll()
}
