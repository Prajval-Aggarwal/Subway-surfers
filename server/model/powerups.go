package model

import (
	"time"

	"gorm.io/gorm"
)

type PowerUp struct {
	PowerUp_Id   string         `json:"powerUpId"`
	PowerUp_Name string         `json:"powerUpName"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}
