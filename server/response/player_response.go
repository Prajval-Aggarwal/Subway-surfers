package response

type PlayerDetails struct {
	P_ID          string `json:"playerId"`
	P_Name        string
	Email         string
	HighScore     int64
	TotalDistance int64
	Coins         int64
	PowerUps      []PowerUp
}

type PowerUp struct {
	PowerUp_Name string
	Quantity     int64
}
