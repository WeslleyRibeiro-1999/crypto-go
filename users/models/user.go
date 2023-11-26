package models

import "time"

type User struct {
	ID          int64     `json:"id" gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
