package repositories

import (
	"errors"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
	"github.com/CAVAh/api-tech-challenge/src/db/gorm"
	"github.com/CAVAh/api-tech-challenge/src/db/models"
)

type OrderRepository struct {
}

func SetDefaultValues(sortBy string, orderBy string, status string) (string, string, string) {
	//TODO: sortBy, orderBy and status needs to be ENUMs, otherwise, it pops syntax error on log

	if sortBy == "" {
		sortBy = "created_at"
	}

	if orderBy == "" {
		orderBy = "ASC"
	}

	return sortBy, orderBy, status
}

func (r OrderRepository) List(sortBy string, orderBy string, status string) ([]entities.Order, error) {
	var orderModel []models.Order

	sortBy, orderBy, status = SetDefaultValues(sortBy, orderBy, status)

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

func (r OrderRepository) FindById(orderId uint) *entities.Order {
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
	model.PaymentStatus = order.PaymentStatus

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

	return order.ToDomain(products)
}

func (r OrderRepository) ExistsOrderProduct(productId uint) bool {
	var orderModel models.OrderProduct

	gorm.DB.First(&orderModel, productId)

	return orderModel.OrderID > 0
}
