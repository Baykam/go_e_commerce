package model

import "time"

type OrderStatus string

const (
	OrderStatusNew        OrderStatus = "new"
	OrderStatusInProgress OrderStatus = "in-progress"
	OrderStatusDone       OrderStatus = "done"
	OrderStatusCancelled  OrderStatus = "cancelled"
)

type Order struct {
	ID         string     `json:"id"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
	Code       string     `json:"code"`
	UserID     string     `json:"user_id"`
	User       *User
	Lines      []*OrderLine `json:"lines"`
	TotalPrice float64      `json:"total_price"`
	Status     OrderStatus  `json:"status"`
}
