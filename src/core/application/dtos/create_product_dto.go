package dtos

//TODO: acho que isso ficaria melhor em driver/http/dtos

type CreateProductDto struct {
	Name        string  `json:"name" validate:"nonzero"`
	Price       float64 `json:"price" validate:"nonzero"`
	Description string  `json:"description" validate:"nonzero"`
	CategoryId  int     `json:"category_id" validate:"nonzero"`
}
