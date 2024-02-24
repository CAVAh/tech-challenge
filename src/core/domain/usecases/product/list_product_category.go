package usecases

import (
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
	"github.com/CAVAh/api-tech-challenge/src/gateways/repositories"
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
