package dtos

type CreateOrderDto struct {
	CustomerId uint                 `json:"customer_id" validate:"nonzero"`
	Products   []ProductInsideOrder `json:"products" validate:"nonzero"`
}

type ProductInsideOrder struct {
	Id          uint   `json:"id" validate:"nonzero"`
	Quantity    int    `json:"quantity" validate:"nonzero"`
	Observation string `json:"observation"`
}

func (dto *CreateOrderDto) GetProductIds() []uint {
	var productIds []uint
	for _, p := range dto.Products {
		productIds = append(productIds, p.Id)
	}
	return productIds
}
