package dto

import "time"

type Order struct {
	ID          int         `json:"id"`
	OrderNumber string      `json:"order_number"`
	SKU         string      `json:"sku"`
	Quantity    int         `json:"quantity"`
	Status      OrderStatus `json:"status"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

type OrderStatus string

const (
	StatusPending   OrderStatus = "PENDING"
	StatusConfirmed OrderStatus = "CONFIRMED"
	StatusCancelled OrderStatus = "CANCELLED"
	StatusFailed    OrderStatus = "FAILED"
)
