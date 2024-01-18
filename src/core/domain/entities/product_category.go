package entities

type ProductCategory struct {
	ID          uint   `json:"id"`
	Description string `json:"description"`
}

func (pc ProductCategory) IsExistingProductCategory() bool {
	return pc.ID > 0
}
