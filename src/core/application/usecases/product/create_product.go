package product

import (
	productDtos "github.com/CAVAh/api-tech-challenge/src/core/application/dtos/product"
	"github.com/CAVAh/api-tech-challenge/src/core/application/ports/repositories"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
	"log"
)

type CreateProductUsecase struct {
	repository                repositories.ProductRepository
	productCategoryRepository repositories.ProductCategoryRepository
}

func BuildCreateProductUsecase(repository repositories.ProductRepository,
	productCategoryRepository repositories.ProductCategoryRepository) *CreateProductUsecase {
	return &CreateProductUsecase{repository: repository, productCategoryRepository: productCategoryRepository}
}

func (p *CreateProductUsecase) Execute(inputDto productDtos.PersistProductDto) (*entities.Product, error) {
	product := entities.NewProduct(0, inputDto.Name, inputDto.Price,
		inputDto.Description, inputDto.CategoryID, "")

	category, err := p.productCategoryRepository.FindById(product.CategoryID)

	if err != nil {
		return nil, nil
	}

	if !category.IsExistingProductCategory() {
		log.Println("Product category doesn't exist! - productCategoryId=", product.CategoryID)
		return product, nil
	}

	return p.repository.Create(product)
}
