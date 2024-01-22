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

func (r OrderRepository) Create(order *models.Order, productIds []uint) (*entities.Order, error) {
	gorm.DB.Where("id IN (?)", productIds).Find(&order.Products)

	if err := gorm.DB.Create(&order).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return nil, errors.New("pedido já existe no sistema")
		} else {
			return nil, errors.New("ocorreu um erro desconhecido ao criar o pedido")
		}
	}

	gorm.DB.Where("id = ?", order.CustomerID).Find(&order.Customer)

	result := order.ToDomain()

	return &result, nil
}

func (r OrderRepository) List(sortBy string, orderBy string, status string) ([]entities.Order, error) {

	var orderModel []models.Order

	if len(sortBy) == 0 {
		sortBy = "created_at"
	}

	if len(sortBy) == 0 {
		sortBy = "ASC"
	}

	if len(status) == 0 {
		gorm.DB.Preload("Products").Preload("Customer").Order(sortBy + " " + orderBy).Find(&orderModel)
	} else {
		gorm.DB.Preload("Products").Preload("Customer").Order(sortBy+" "+orderBy).Where("status = ?", status).Find(&orderModel)
	}

	var order []entities.Order

	for _, orderModel := range orderModel {
		order = append(order, orderModel.ToDomain())
	}

	return order, nil
}

func (r OrderRepository) FindyId(orderId uint) *entities.Order {
	var orderModel models.Order
	gorm.DB.First(&orderModel, orderId)

	result := orderModel.ToDomain()

	return &result
}

func (r OrderRepository) Update(order *entities.Order) {
	var orderModel models.Order

	gorm.DB.First(&orderModel, order.ID)

	gorm.DB.Model(&orderModel).Updates(models.Order{Status: order.Status})
}
