package usecases

import (
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
	"github.com/CAVAh/api-tech-challenge/src/gateways"
)

type ReadProductUsecase struct {
	repository gateways.ProductRepository
}

func BuildReadProductUsecase(repository gateways.ProductRepository) *ReadProductUsecase {
	return &ReadProductUsecase{repository: repository}
}

func (p ReadProductUsecase) Execute(id uint) (*entities.Product, error) {
	return p.repository.FindById(id)
}
