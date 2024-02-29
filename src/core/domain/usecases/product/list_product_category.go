package usecases

import (
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
	"github.com/CAVAh/api-tech-challenge/src/gateways"
)

type ListProductCategoryUsecase struct {
	repository gateways.ProductCategoryRepository
}

func BuildListProductCategoryUsecase(repository gateways.ProductCategoryRepository) *ListProductCategoryUsecase {
	return &ListProductCategoryUsecase{repository: repository}
}

func (p ListProductCategoryUsecase) Execute() ([]entities.ProductCategory, error) {
	return p.repository.FindAll()
}
