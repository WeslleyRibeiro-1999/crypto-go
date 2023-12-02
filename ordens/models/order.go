package models

import "time"

type OrderType string

const (
	MarketOrder OrderType = "market"
	LimitOrder  OrderType = "limit"
)

type Order struct {
	ID        int64     `json:"id" gorm:"column:id;primary_key;AUTO_INCREMENT"`
	UserID    int64     `json:"user_id"`
	Pair      string    `json:"pair"`
	Amount    float64   `json:"amount" gorm:"column:amount;type:numeric(10,2)"`
	Direction string    `json:"direction"`
	Type      OrderType `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateOrderRequest struct {
	UserID    int64     `json:"user_id"`
	Pair      string    `json:"pair"`
	Amount    float64   `json:"amount" gorm:"column:amount;type:numeric(10,2)"`
	Direction string    `json:"direction"`
	Type      OrderType `json:"type"`
}

type CreateOrderResponse struct {
	Pair      string  `json:"pair"`
	Amount    float64 `json:"amount"`
	Direction string  `json:"direction"`
}
