package repositories

import (
	"errors"
	"github.com/CAVAh/api-tech-challenge/src/adapters/driven/db/gorm"
	"github.com/CAVAh/api-tech-challenge/src/adapters/driven/db/models"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
	"strings"
)

type OrderRepository struct {
}

func (r OrderRepository) Create(costumerId int, productIds []int) (*entities.Order, error) {
	var order = models.Order{}

	gorm.DB.Where("id = ?", costumerId).Find(&order.Customer)
	gorm.DB.Where("id IN (?)", productIds).Find(&order.Products)

	if err := gorm.DB.Create(&order).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return nil, errors.New("pedido já existe no sistema")
		} else {
			return nil, errors.New("ocorreu um erro desconhecido ao criar o pedido")
		}
	}

	result := order.ToDomain()

	return &result, nil
}

func (r OrderRepository) List() ([]entities.Order, error) {
	var orders []models.Order
	var response []entities.Order

	gorm.DB.Preload("Customer").Preload("Products").Find(&orders)

	for _, order := range orders {
		response = append(response, order.ToDomain())
	}

	return response, nil
}
