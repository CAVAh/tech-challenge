package product

import (
	dtosProd "github.com/CAVAh/api-tech-challenge/src/core/application/dtos/product"
	"github.com/CAVAh/api-tech-challenge/src/core/application/ports/repositories"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
)

type Edit struct {
	repository repositories.ProductRepository
}

func BuildEdit(repository repositories.ProductRepository) *Edit {
	return &Edit{repository: repository}
}

func (p *Edit) Execute(inputDto dtosProd.PersistProductDto, id uint) (*entities.Product, error) {
	retrievedProduct, err := p.repository.FindById(id)

	retrievedProduct.PatchFields(inputDto.Name, inputDto.Price, inputDto.Description, inputDto.CategoryID)

	if err == nil && retrievedProduct.IsExistingProduct() {
		return p.repository.Edit(retrievedProduct)
	}

	emptyProduct := entities.Product{}

	return &emptyProduct, err
}
