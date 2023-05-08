package leaderboard

import (
	"subway/server/db"
	"subway/server/response"
	"subway/server/utils"

	"github.com/gin-gonic/gin"
)

// Showing the world position of the player accoring to high score
func ShowLeaderBoardService(ctx *gin.Context) {
	query := "SELECT p_name,high_score FROM players ORDER BY high_score;"
	var leaderboard []response.Leaderboard
	err := db.RawQuery(query, leaderboard)
	if err != nil {
		response.ErrorResponse(ctx, utils.BAD_REQUEST, "Failed to fetch leaderboard information")
		return
	}

	response.ShowResponse("Sucess", utils.SUCCESS, "Leaderboard fetched sucessfully", leaderboard, ctx)
}
