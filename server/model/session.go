package model

type Session struct {
	S_ID  string `json:"sessionId" gorm:"default:uuid_generate_v4();primaryKey"`
	P_Id  string `json:"playerId"`
	Token string `json:"token"`
}

type ResetSession struct {
	P_ID       string `json:"playerEmail"`
	ResetToken string `json:"resetToken"`
}
