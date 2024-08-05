package dto

import "github.com/Sotatek-HungNgo3/be-practical-order/enum"

type OrderDto struct {
	ProductName   string             `json:"productName" binding:"required"`
	UnitPrice     float64            `json:"unitPrice" binding:"required"`
	Quantity      int64              `json:"quantity" binding:"required"`
	PaymentMethod enum.PaymentMethod `json:"paymentMethod" binding:"required,paymentMethod"`
}
