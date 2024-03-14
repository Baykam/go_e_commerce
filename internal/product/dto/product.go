package dto

import (
	"golang_testing_grpc/pkg/paging"
	"time"
)

type Product struct {
	ID          string    `json:"id"`
	Code        string    `json:"code"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Active      bool      `json:"active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ListProductReq struct {
	Name      string `json:"name,omitempty" form:"name"`
	Code      string `json:"code,omitempty" form:"code"`
	Page      int64  `json:"-" form:"page"`
	Limit     int64  `json:"-" form:"limit"`
	OrderBy   string `json:"-" form:"order_by"`
	OrderDesc bool   `json:"-" form:"order_desc"`
}

type ListProductRes struct {
	Products   []*Product        `json:"products"`
	Pagination paging.Pagination `json:"pagination"`
}
