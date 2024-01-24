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

func (r OrderRepository) Create(order *entities.Order) (*entities.Order, error) {
	var model models.Order

	customerExists := gorm.DB.First(&model.Customer, order.Customer.ID)

	if customerExists.Error != nil {
		return nil, errors.New("o cliente informado não existe!")
	}

	productIDs := order.GetProductIds()

	var productModel models.Product
	productsExists := gorm.DB.Find(&productModel, productIDs)

	//TODO: criar na mão order.Products
	var productsOrderModel []models.OrderProduct

	for _, p := range order.Products {
		productsOrderModel = append(productsOrderModel, models.OrderProduct{ProductID: p.Product.ID, Quantity: p.Quantity, Observation: p.Observation})
	}

	model.Products = productsOrderModel

	if productsExists.Error != nil {
		return nil, errors.New("ocorreu um erro ao encontrar os produtos!")
	}

	//if productsExists.RowsAffected != int64(len(productIDs)) {
	//	return nil, errors.New("alguns dos produtos não foram encontrados!")
	//}

	model.Status = order.Status

	if err := gorm.DB.Create(&model).Error; err != nil {
		return &entities.Order{}, errors.New("ocorreu um erro desconhecido ao criar o pedido")
	}

	result := model.ToDomain()
	return &(result), nil
}
