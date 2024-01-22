package product

import (
	"github.com/CAVAh/api-tech-challenge/src/core/application/ports/repositories"
)

type DeleteProductUsecase struct {
	repository repositories.ProductRepository
}

func BuildDeleteProductUsecase(repository repositories.ProductRepository) *DeleteProductUsecase {
	return &DeleteProductUsecase{repository: repository}
}

func (p *DeleteProductUsecase) Execute(id int) error {
	return p.repository.DeleteById(id)
}
