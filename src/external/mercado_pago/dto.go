package mercado_pago

import (
	"fmt"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
	"time"
)

type Item struct {
	SkuNumber   string  `json:"sku_number"`
	Category    string  `json:"category"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	UnitPrice   float64 `json:"unit_price"`
	Quantity    int     `json:"quantity"`
	UnitMeasure string  `json:"unit_measure"`
	TotalAmount float64 `json:"total_amount"`
}

type MarcadoPagoRequest struct {
	Description       string  `json:"description"`
	ExpirationDate    string  `json:"expiration_date"`
	ExternalReference string  `json:"external_reference"`
	Items             []Item  `json:"items"`
	NotificationUrl   string  `json:"notification_url"`
	Title             string  `json:"title"`
	TotalAmount       float64 `json:"total_amount"`
}

type QrCreatedResponse struct {
	InStoreOrderId string `json:"in_store_order_id"`
	QrData         string `json:"qr_data"`
}

func ToItem(productInsideOrder entities.ProductInsideOrder) Item {
	product := productInsideOrder.Product
	item := Item{
		SkuNumber:   fmt.Sprint(product.ID),
		Category:    fmt.Sprint(product.CategoryId),
		Title:       product.Name,
		Description: product.Description,
		UnitPrice:   product.Price,
		Quantity:    productInsideOrder.Quantity,
		UnitMeasure: "unit",
		TotalAmount: float64(productInsideOrder.Quantity) * product.Price,
	}
	return item
}

func ToItems(products []entities.ProductInsideOrder) ([]Item, float64) {
	var items []Item
	var totalAmount float64

	for _, product := range products {
		item := ToItem(product)

		items = append(items, item)
		totalAmount = totalAmount + item.TotalAmount
	}
	return items, totalAmount
}

func MountMercadoPagoRequest(order entities.Order) MarcadoPagoRequest {
	items, totalAmount := ToItems(order.Products)

	t := time.Now().Local().Add(time.Minute * time.Duration(5))

	fiveMinExpiration := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d.%03d-03:00",
		t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond()%1000)

	return MarcadoPagoRequest{
		Title:             fmt.Sprintf("Pedido #%d", order.ID),
		Description:       fmt.Sprintf("Pedido #%d de %s", order.ID, order.Customer.Name),
		ExternalReference: fmt.Sprintf("%d", order.ID),
		Items:             items,
		TotalAmount:       totalAmount,
		ExpirationDate:    fiveMinExpiration,
		NotificationUrl:   GetAppNotificationUrl(),
	}
}
