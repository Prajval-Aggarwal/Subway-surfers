package model

type DailyReward struct {
	P_Id      string `json:"playerId"`
	Date      string `json:"date"`
	Quantity  int64  `json:"quantity"`
	PowerUpId string `json:"powerUp"`
	Status    string `json:"status"`
}
