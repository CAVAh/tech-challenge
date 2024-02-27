package usecases

import (
	dtosProd "github.com/CAVAh/api-tech-challenge/src/core/domain/dtos/product"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
	"github.com/CAVAh/api-tech-challenge/src/gateways/repositories"
)

type EditProductUsecase struct {
	repository repositories.ProductRepository
}

func BuildEditProductUsecase(repository repositories.ProductRepository) *EditProductUsecase {
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