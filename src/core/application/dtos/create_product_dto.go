package dtos

type CreateProductDto struct {
	Name        string  `json:"name" validate:"nonzero"`
	Price       float64 `json:"price" validate:"nonzero"`
	Description string  `json:"description" validate:"nonzero"`
	CategoryId  uint    `json:"categoryId" validate:"nonzero"`
}
