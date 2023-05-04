package model

import "time"

type Cart struct {
	CartId    string `json:"cartId" gorm:"default:uuid_generate_v4();primaryKey"`
	CreatedAt time.Time
}
type CartItem struct {
	CartId   string
	ItemId   string
	Quantity int64
}
