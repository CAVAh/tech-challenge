package entities

type Order struct {
	ID            uint                 `json:"id"`
	Customer      Customer             `json:"customer"`
	Products      []ProductInsideOrder `json:"products"`
	Status        string               `json:"status"`
	PaymentStatus string               `json:"paymentStatus"`
	CreatedAt     string               `json:"createdAt"`
	UpdatedAt     string               `json:"updatedAt"`
}

func (o *Order) GetProductIds() []uint {
	var productIds []uint
	for _, p := range o.Products {
		productIds = append(productIds, p.Product.ID)
	}
	return productIds
}
