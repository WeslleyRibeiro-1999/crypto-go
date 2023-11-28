package models

import "time"

type User struct {
	ID          int64     `json:"id" gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Name        string    `json:"name" validate:"required"`
	Email       string    `json:"email" gorm:"unique" validate:"required"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateUser struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

type CreatedUser struct {
	Status bool   `json:"status"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

type UserResponse struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

type UpdateUser struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}
