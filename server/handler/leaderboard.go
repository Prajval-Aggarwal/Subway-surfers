package handler

import (
	"subway/server/services/leaderboard"

	"github.com/gin-gonic/gin"
)

func ShowLeaderBoardHandler(ctx *gin.Context) {
	leaderboard.ShowLeaderBoardService(ctx)
}
