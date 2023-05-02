package handler

import (
	"subway/server/services/leaderboard"

	"github.com/gin-gonic/gin"
)

//	@Description	Shows the leaderboard
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	response.Success
//	@Failure		400	{object}	response.Error
//	@Tags			Leaderboard
//	@Router			/show-leaderboard [get]
func ShowLeaderBoardHandler(ctx *gin.Context) {
	leaderboard.ShowLeaderBoardService(ctx)
}
