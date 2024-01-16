package repositories

import (
	"github.com/CAVAh/api-tech-challenge/src/adapters/driven/db/models"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
)

type ProductRepository interface {
	Create(product *models.Product) (*entities.Product, error)
	List() ([]entities.Product, error)
	FindById(ids []int) ([]entities.Product, error)
}
