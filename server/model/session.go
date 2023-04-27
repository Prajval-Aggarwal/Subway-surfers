package model

type Session struct {
	S_ID  string `json:"sessionId"`
	P_Id  string `json:"playerId"`
	Token string `json:"token"`
}
