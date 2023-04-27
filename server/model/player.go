package model

import (
	"time"

	"gorm.io/gorm"
)

type Player struct {
	P_ID          string         `json:"playerId" gorm:"default:uuid_generate_v4()"`
	P_Name        string         `json:"playerName"`
	Email         string         `json:"email" gorm:"unique"`
	Password      string         `json:"password" `
	HighScore     int64          `json:"highScore"`
	TotalDistance int64          `json:"totalDistance"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

type PlayerCoins struct {
	P_ID  string `json:"playerId"`
	Coins int64  `json:"coins"`
}

type PlayerPowerUps struct {
	P_ID       string `json:"playerId"`
	PowerUp_Id string `json:"powerUpId"`
	Quantity   int64  `json:"quantity"`
}
