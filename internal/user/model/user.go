package model

import "time"

type UserRole string

const (
	Admin    UserRole = "admin"
	Customer UserRole = "customer"
	Guest    UserRole = "guest"
)

type User struct {
	ID        string     `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	Role      UserRole   `json:"role"`
}
