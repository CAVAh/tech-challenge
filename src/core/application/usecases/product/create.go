package product

import (
	productDtos "github.com/CAVAh/api-tech-challenge/src/core/application/dtos/product"
	"github.com/CAVAh/api-tech-challenge/src/core/application/ports/repositories"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
)

type Create struct {
	repository                repositories.ProductRepository
	productCategoryRepository repositories.ProductCategoryRepository
}

func BuildCreate(repository repositories.ProductRepository,
	productCategoryRepository repositories.ProductCategoryRepository) *Create {
	return &Create{repository: repository, productCategoryRepository: productCategoryRepository}
}

func (p *Create) Execute(inputDto productDtos.PersistProductDto) (*entities.Product, error) {
	product := entities.Product{
		Name:        inputDto.Name,
		Price:       inputDto.Price,
		Description: inputDto.Description,
		CategoryID:  inputDto.CategoryID,
	}

	category, err := p.productCategoryRepository.FindById(product.CategoryID)

	if err != nil {
		return nil, nil
	}

	if !category.IsExistingProductCategory() {
		return &product, nil
	}

	return p.repository.Create(&product)
}
