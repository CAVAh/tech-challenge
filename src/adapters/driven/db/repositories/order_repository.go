package repositories

import (
	"errors"
	"strings"

	"github.com/CAVAh/api-tech-challenge/src/adapters/driven/db/gorm"
	"github.com/CAVAh/api-tech-challenge/src/adapters/driven/db/models"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
)

type OrderRepository struct {
}

func (r OrderRepository) Create(order *models.Order, productIds []int) (*entities.Order, error) {
	gorm.DB.Where("id IN (?)", productIds).Find(&order.Products)

	if err := gorm.DB.Create(&order).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return nil, errors.New("pedido já existe no sistema")
		} else {
			return nil, errors.New("ocorreu um erro desconhecido ao criar o pedido")
		}
	}

	gorm.DB.Where("id = ?", order.CustomerId).Find(&order.Customer)

	result := order.ToDomain()

	return &result, nil
}

func (r OrderRepository) List(pageSize int64, pageNumber int64, status string) ([]entities.Order, error) {

	var orderModel []models.Order

	if len(status) == 0 {
		gorm.DB.Preload("Products").Preload("Customer").Find(&orderModel)
	} else {
		gorm.DB.Preload("Products").Preload("Customer").Where("status = ?", status).Find(&orderModel)
	}

	var order []entities.Order

	for _, orderModel := range orderModel {
		order = append(order, orderModel.ToDomain())
	}

	return order, nil
}
