package dtos

type CreateOrderDto struct {
	CustomerId int                  `json:"customer_id" validate:"nonzero"`
	Products   []ProductInsideOrder `json:"products" validate:"nonzero"`
}

type ProductInsideOrder struct {
	Id          int    `json:"id" validate:"nonzero"`
	Quantity    int    `json:"quantity" validate:"nonzero"`
	Observation string `json:"observation"`
}

func (dto *CreateOrderDto) GetProductIds() []int {
	var productIds []int
	for _, p := range dto.Products {
		productIds = append(productIds, p.Id)
	}
	return productIds
}

func (dto *CreateOrderDto) GetProduct(id int) ProductInsideOrder {
	for _, p := range dto.Products {
		if id == p.Id {
			return p
		}
	}
	return ProductInsideOrder{Quantity: 1}
}
