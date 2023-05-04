package model

type PowerUp struct {
	PowerUp_Id   string `json:"powerUpId" gorm:"default:uuid_generate_v4()"`
	PowerUp_Name string `json:"powerUpName"`
	Price        int64  `json:"price"`
}
