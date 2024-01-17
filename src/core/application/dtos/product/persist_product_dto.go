package product

type PersistProductDto struct {
	ID          uint
	Name        string  `json:"name" validate:"nonzero"`
	Price       float64 `json:"price" validate:"nonzero"`
	Description string  `json:"description" validate:"nonzero"`
	CategoryID  uint    `json:"categoryId" validate:"nonzero"`
}
