package repositories

import "github.com/CAVAh/api-tech-challenge/src/core/domain/entities"

type ProductRepository interface {
	Create(entity *entities.Product) (*entities.Product, error)
	FindById(id uint) (*entities.Product, error)
	FindByIds(ids []uint) ([]entities.Product, error)
	FindAll() ([]entities.Product, error)
	FindByCategoryId(categoryId uint) ([]entities.Product, error)
	Edit(entity *entities.Product) (*entities.Product, error)
	DeleteById(id uint) error
}
