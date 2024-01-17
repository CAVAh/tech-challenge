package product

import (
	"github.com/CAVAh/api-tech-challenge/src/core/application/ports/repositories"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
)

type Read struct {
	repository repositories.ProductRepository
}

func BuildRead(repository repositories.ProductRepository) *Read {
	return &Read{repository: repository}
}

func (p Read) Execute(id uint) (*entities.Product, error) {
	return p.repository.FindById(id)
}
