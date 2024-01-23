package usecases

import (
	"github.com/CAVAh/api-tech-challenge/src/core/application/dtos"
	"github.com/CAVAh/api-tech-challenge/src/core/application/ports/repositories"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
)

type CreateProductUsecase struct {
	ProductRepository repositories.ProductRepository
}

func (r *CreateProductUsecase) Execute(inputDto dtos.CreateProductDto) (*entities.Product, error) {
	product := entities.Product{
		Name:        inputDto.Name,
		Price:       inputDto.Price,
		Description: inputDto.Description,
		CategoryID:  inputDto.CategoryId,
	}

	return r.ProductRepository.Create(&product)
}
