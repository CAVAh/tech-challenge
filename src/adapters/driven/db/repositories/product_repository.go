package repositories

import (
	"errors"
	"strings"

	"github.com/CAVAh/api-tech-challenge/src/adapters/driven/db/gorm"
	"github.com/CAVAh/api-tech-challenge/src/adapters/driven/db/models"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
)

type ProductRepository struct {
}

func (r ProductRepository) Create(product *models.Product) (*entities.Product, error) {
	if err := gorm.DB.Create(&product).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return nil, errors.New("produto já existe no sistema")
		} else {
			return nil, errors.New("ocorreu um erro desconhecido ao criar o produto")
		}
	}

	result := product.ToDomain()

	return &result, nil
}

func (r ProductRepository) List() ([]entities.Product, error) {
	var products []models.Product
	var response []entities.Product

	gorm.DB.Find(&products)

	for _, product := range products {
		response = append(response, product.ToDomain())
	}

	return response, nil
}
