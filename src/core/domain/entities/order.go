package entities

type Order struct {
	ID        uint           `json:"id"`
	CreatedAt string         `json:"createdAt"`
	UpdatedAt string         `json:"updatedAt"`
	Status    string         `json:"status"`
	Customer  Customer       `json:"customer"`
	Products  []OrderProduct `json:"products"`
}

func (order *Order) GetProductIds() []uint {
	var productIds []uint

	for _, p := range order.Products {
		productIds = append(productIds, p.ProductID)
	}

	return productIds
}
