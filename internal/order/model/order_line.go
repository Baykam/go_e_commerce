package model

import "time"

type OrderLine struct {
	ID        string     `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	OrderID   string     `json:"order_id"`
	ProductID string     `json:"product_id"`
	Product   *Product
	Quantity  uint    `json:"quantity"`
	Price     float64 `json:"price"`
}
