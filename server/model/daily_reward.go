package model

import "time"

type DailyReward struct {
	P_Id      string `json:"playerId" gorm:"primaryKey"`
	Quantity  int64  `json:"quantity"`
	PowerUpId string `json:"powerUp"`
	Status    string `json:"status"`
	CreatedAt time.Time
}
