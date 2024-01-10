package dtos

type CreateCustomerDto struct {
	Name  string `json:"name" validate:"nonzero"`
	CPF   string `json:"cpf" validate:"len=11, regexp=^[0-9]*$"`
	Email string `json:"email" validate:"nonzero, regexp=^[a-z0-9._-]+@[a-z0-9.-]+\\.[a-z]*$"`
}
