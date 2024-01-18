package repositories

import (
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
)

type ProductRepository interface {
	Create(product *entities.Product) (*entities.Product, error)
	List() ([]entities.Product, error)
	FindById(ids []int) ([]entities.Product, error)
}
