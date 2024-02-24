package usecases

import (
	productDtos "github.com/CAVAh/api-tech-challenge/src/core/domain/dtos/product"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
	repositories2 "github.com/CAVAh/api-tech-challenge/src/gateways/repositories"
	"log"
)

type CreateProductUsecase struct {
	repository                repositories2.ProductRepository
	productCategoryRepository repositories2.ProductCategoryRepository
}

func BuildCreateProductUsecase(repository repositories2.ProductRepository,
	productCategoryRepository repositories2.ProductCategoryRepository) *CreateProductUsecase {
	return &CreateProductUsecase{repository: repository, productCategoryRepository: productCategoryRepository}
}

func (p *CreateProductUsecase) Execute(inputDto productDtos.PersistProductDto) (*entities.Product, error) {
	product := &entities.Product{
		Name:        inputDto.Name,
		Price:       inputDto.Price,
		Description: inputDto.Description,
		CategoryId:  inputDto.CategoryID,
	}

	category, err := p.productCategoryRepository.FindById(product.CategoryId)

	if err != nil {
		return nil, nil
	}

	if !category.IsExistingProductCategory() {
		log.Println("Product category doesn't exist! - productCategoryId=", product.CategoryId)
		return product, nil
	}

	return p.repository.Create(product)
}
