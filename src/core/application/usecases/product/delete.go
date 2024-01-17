package product

import (
	"github.com/CAVAh/api-tech-challenge/src/core/application/ports/repositories"
)

type Delete struct {
	repository repositories.ProductRepository
}

func BuildDelete(repository repositories.ProductRepository) *Delete {
	return &Delete{repository: repository}
}

func (p *Delete) Execute(id uint) error {
	return p.repository.DeleteById(id)
}
