package repositories

import "github.com/CAVAh/api-tech-challenge/src/core/domain/entities"

type ProductRepository interface {
	Create(entity *entities.Product) (*entities.Product, error)
	FindById(id int) (*entities.Product, error)
	FindByIds(ids []int) ([]entities.Product, error)
	FindAll() ([]entities.Product, error)
	FindByCategoryId(categoryId int) ([]entities.Product, error)
	Edit(entity *entities.Product) (*entities.Product, error)
	DeleteById(id int) error
}
