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

func (r ProductRepository) Create(entity *entities.Product) (*entities.Product, error) {
	product := models.Product{
		Name:        entity.Name,
		Price:       entity.Price,
		Description: entity.Description,
		CategoryId:  entity.CategoryId,
	}

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

func (r ProductRepository) FindById(ids []int) ([]entities.Product, error) {
	var products []models.Product
	var response []entities.Product

	gorm.DB.Where("id in (?)", ids).Find(&products)

	for _, product := range products {
		response = append(response, product.ToDomain())
	}

	return response, nil
}
