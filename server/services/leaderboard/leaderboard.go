package leaderboard

import (
	"subway/server/db"
	"subway/server/response"

	"github.com/gin-gonic/gin"
)

func ShowLeaderBoardService(ctx *gin.Context) {
	query := "SELECT p_name,high_score FROM players ORDER BY high_score;"
	var leaderboard []response.Leaderboard
	err := db.RawQuery(query, leaderboard)
	if err != nil {
		response.ErrorResponse(ctx, 400, "Failed to fetch leaderboard information")
		return
	}

	response.ShowResponse("Sucess", 200, "Leaderboard fetched sucessfully", leaderboard, ctx)
}
