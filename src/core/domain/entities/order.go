package entities

type Order struct {
	ID        uint                 `json:"id"`
	Customer  Customer             `json:"customer"`
	Products  []ProductInsideOrder `json:"products"`
	Status    string               `json:"status"`
	CreatedAt string               `json:"createdAt"`
	UpdatedAt string               `json:"updatedAt"`
}

func (o *Order) GetProductIds() []uint {
	var productIds []uint
	for _, p := range o.Products {
		productIds = append(productIds, p.Product.ID)
	}
	return productIds
}

func (o *Order) GetProductInsideOrderById(id uint) ProductInsideOrder {
	for _, p := range o.Products {
		if id == p.Product.ID {
			return p
		}
	}
	return ProductInsideOrder{Quantity: 1}
}
