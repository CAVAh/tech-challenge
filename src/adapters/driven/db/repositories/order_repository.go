package repositories

import (
	"errors"
	"github.com/CAVAh/api-tech-challenge/src/adapters/driven/db/gorm"
	"github.com/CAVAh/api-tech-challenge/src/adapters/driven/db/models"
	"github.com/CAVAh/api-tech-challenge/src/core/application/dtos"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
)

type OrderRepository struct {
}

func (r OrderRepository) Create(dto dtos.CreateOrderDto) (*entities.Order, error) {
	var order models.Order

	gorm.DB.Where("id = ?", dto.CustomerId).Find(&order.Customer)
	gorm.DB.Where("id IN (?)", dto.GetProductIds()).Find(&order.Products)

	if err := gorm.DB.Create(&order).Error; err != nil {
		return nil, errors.New("ocorreu um erro desconhecido ao criar o pedido")
	}

	for _, p := range order.Products {
		var op models.OrderProduct
		var product = dto.GetProduct(int(p.ID))
		gorm.DB.Where("order_id = ? and product_id = ?", order.ID, p.ID).Find(&op)
		gorm.DB.Model(&op).
			Update("Quantity", product.Quantity).
			Update("Observation", product.Observation)
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
