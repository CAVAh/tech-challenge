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
		order = append(order, OrderModelToOrderEntity(&orderModel))
	}

	return order, nil
}

func (r OrderRepository) FindyId(orderId uint) *entities.Order {
	var orderModel models.Order
	gorm.DB.First(&orderModel, orderId)

	result := OrderModelToOrderEntity(&orderModel)

	return &result
}

func (r OrderRepository) Update(order *entities.Order) {
	var orderModel models.Order

	gorm.DB.First(&orderModel, order.ID)

	gorm.DB.Model(&orderModel).Updates(models.Order{Status: order.Status})
}

func (r OrderRepository) Create(order *entities.Order) (*entities.Order, error) {
	var model models.Order

	gorm.DB.First(&model.Customer, order.Customer.ID)
	gorm.DB.Find(&model.Products, order.GetProductIds())

	var productsOrderModel []models.OrderProduct
	for _, p := range order.Products {
		productsOrderModel = append(productsOrderModel, models.OrderProduct{
			ProductID:   p.Product.ID,
			Quantity:    p.Quantity,
			Observation: p.Observation,
		})
	}

	model.Products = productsOrderModel
	model.Status = order.Status

	if err := gorm.DB.Create(&model).Error; err != nil {
		return &entities.Order{}, errors.New("ocorreu um erro desconhecido ao criar o pedido")
	}

	result := OrderModelToOrderEntity(&model)
	return &result, nil
}

func OrderModelToOrderEntity(order *models.Order) entities.Order {
	gorm.DB.Preload("Products").Preload("Customer").Where("id = ?", order.ID).Find(&order)

	var orderProducts []models.OrderProduct
	gorm.DB.Preload("Product").Where("order_id = ?", order.ID).Find(&orderProducts)

	var products []entities.ProductInsideOrder
	for _, p := range orderProducts {
		products = append(products, entities.ProductInsideOrder{
			Product:     p.Product.ToDomain(),
			Quantity:    p.Quantity,
			Observation: p.Observation,
		})
	}

	return entities.Order{
		ID:        order.ID,
		CreatedAt: order.CreatedAt.Format("2006-01-02 15:04:05"),
		Customer:  order.Customer.ToDomain(),
		Status:    order.Status,
		Products:  products,
	}
}
