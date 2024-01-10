package dtos

type ListCustomerDto struct {
	CPF string `form:"cpf" json:"cpf" validate:"min=0,max=11,regexp=^[0-9]*$"`
}
