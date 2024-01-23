package usecases

import (
	dtosProd "github.com/CAVAh/api-tech-challenge/src/core/application/dtos/product"
	"github.com/CAVAh/api-tech-challenge/src/core/application/ports/repositories"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
)

type EditProductUsecase struct {
	repository repositories.ProductRepository
}

func BuildEditProductUsecase(repository repositories.ProductRepository) *EditProductUsecase {
	return &EditProductUsecase{repository: repository}
}

func (p *EditProductUsecase) Execute(inputDto dtosProd.PersistProductDto) (*entities.Product, error) {
	retrievedProduct, err := p.repository.FindById(inputDto.ID)

	retrievedProduct.PatchFields(inputDto.Name, inputDto.Price, inputDto.Description, inputDto.CategoryID)

	if err == nil && retrievedProduct.IsExistingProduct() {
		return p.repository.Edit(retrievedProduct)
	}

	emptyProduct := entities.Product{}

	return &emptyProduct, err
}
