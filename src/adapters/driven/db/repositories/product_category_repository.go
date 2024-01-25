package repositories

import (
	"github.com/CAVAh/api-tech-challenge/src/adapters/driven/db/gorm"
	"github.com/CAVAh/api-tech-challenge/src/adapters/driven/db/models"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
)

type ProductCategoryRepository struct {
}

func (r ProductCategoryRepository) FindAll() ([]entities.ProductCategory, error) {
	var categories []models.ProductCategory

	err := checkError(gorm.DB.Find(&categories))

	if err != nil {
		return []entities.ProductCategory{}, err
	}

	productEntities := []entities.ProductCategory{}

	for _, category := range categories {
		productEntities = append(productEntities, category.ToDomain())
	}

	return productEntities, nil
}

func (r ProductCategoryRepository) FindById(id int) (*entities.ProductCategory, error) {
	var productCategory models.ProductCategory

	err := checkError(gorm.DB.Find(&productCategory, id))

	if err != nil {
		return &entities.ProductCategory{}, err
	}

	result := productCategory.ToDomain()

	return &result, nil
}
