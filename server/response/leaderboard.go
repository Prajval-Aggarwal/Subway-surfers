package response

type Leaderboard struct {
	P_Name string `json:"playerName"`
	Score  int64  `json:"score"`
}
