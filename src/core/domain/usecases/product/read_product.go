package usecases

import (
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
	"github.com/CAVAh/api-tech-challenge/src/gateways/repositories"
)

type ReadProductUsecase struct {
	repository repositories.ProductRepository
}

func BuildReadProductUsecase(repository repositories.ProductRepository) *ReadProductUsecase {
	return &ReadProductUsecase{repository: repository}
}

func (p ReadProductUsecase) Execute(id uint) (*entities.Product, error) {
	return p.repository.FindById(id)
}