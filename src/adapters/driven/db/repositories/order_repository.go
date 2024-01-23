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

func (r OrderRepository) Create(dto dtos.CreateOrderDto) (*entities.Order, error) {
	var order models.Order

	gorm.DB.Find(&order.Customer, dto.CustomerId)
	gorm.DB.Find(&order.Products, dto.GetProductIds())

	if err := gorm.DB.Create(&order).Error; err != nil {
		return &entities.Order{}, errors.New("ocorreu um erro desconhecido ao criar o pedido")
	}

	for _, p := range order.Products {
		var op models.OrderProduct
		var product = dto.GetProduct(p.ID)
		gorm.DB.Where("order_id = ? and product_id = ?", order.ID, p.ID).Find(&op)
		gorm.DB.Model(&op).
			Update("Quantity", product.Quantity).
			Update("Observation", product.Observation)
	}

	result := order.ToDomain()
	return &(result), nil
}
