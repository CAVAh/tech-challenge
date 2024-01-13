package dtos

//TODO: acho que isso ficaria melhor em driver/http/dtos

type CreateOrderDto struct {
	CustomerId int   `json:"customer_id" validate:"nonzero"`
	ProductIds []int `json:"product_ids" validate:"nonzero"`
}
