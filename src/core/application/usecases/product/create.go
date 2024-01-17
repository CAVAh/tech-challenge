package product

import (
	productDtos "github.com/CAVAh/api-tech-challenge/src/core/application/dtos/product"
	"github.com/CAVAh/api-tech-challenge/src/core/application/ports/repositories"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
)

type Create struct {
	repository repositories.ProductRepository
}

func BuildCreate(repository repositories.ProductRepository) *Create {
	return &Create{repository: repository}
}

func (p *Create) Execute(inputDto productDtos.PersistProductDto) (*entities.Product, error) {
	product := entities.Product{
		Name:        inputDto.Name,
		Price:       inputDto.Price,
		Description: inputDto.Description,
		CategoryID:  inputDto.CategoryID,
	}

	return p.repository.Create(&product)
}
