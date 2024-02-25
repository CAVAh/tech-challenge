package dtos

type ChangeOrderStatusDto struct {
	OrderId        uint   `json:"orderId"`
	ChangeToStatus string `json:"changeToStatus"`
}
