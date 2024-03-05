package usecases

import (
	"github.com/CAVAh/api-tech-challenge/src/adapters/gateways"
	dtosProd "github.com/CAVAh/api-tech-challenge/src/core/domain/dtos/product"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
)

type EditProductUsecase struct {
	repository gateways.ProductRepository
}

func BuildEditProductUsecase(repository gateways.ProductRepository) *EditProductUsecase {
	return &EditProductUsecase{repository: repository}
}

func (p *EditProductUsecase) Execute(inputDto dtosProd.PersistProductDto) (*entities.Product, error) {
	retrievedProduct, err := p.repository.FindById(inputDto.ID)

	if err == nil && retrievedProduct.IsExistingProduct() {
		retrievedProduct.PatchFields(inputDto.Name, inputDto.Price, inputDto.Description, inputDto.CategoryID)

		return p.repository.Edit(retrievedProduct)
	}

	emptyProduct := entities.Product{}

	return &emptyProduct, err
}
