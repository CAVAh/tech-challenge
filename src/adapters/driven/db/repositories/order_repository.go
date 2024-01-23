package repositories

import (
	"errors"
	"github.com/CAVAh/api-tech-challenge/src/adapters/driven/db/gorm"
	"github.com/CAVAh/api-tech-challenge/src/adapters/driven/db/models"
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

func (r OrderRepository) Create(order entities.Order) (*entities.Order, error) {
	var model models.Order

	gorm.DB.Find(&model.Customer, order.Customer.ID)
	gorm.DB.Find(&model.Products, order.GetProductIds())

	if err := gorm.DB.Create(&model).Error; err != nil {
		return &entities.Order{}, errors.New("ocorreu um erro desconhecido ao criar o pedido")
	}

	for _, p := range model.Products {
		var op models.OrderProduct
		var product = order.GetProductInsideOrderById(p.ID)

		gorm.DB.Where("order_id = ? and product_id = ?", model.ID, p.ID).Find(&op)
		gorm.DB.Model(&op).
			Update("Quantity", product.Quantity).
			Update("Observation", product.Observation)
	}

	result := model.ToDomain()
	return &(result), nil
}
