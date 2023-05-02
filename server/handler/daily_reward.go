package handler

import (
	"fmt"
	"subway/server/response"
	"subway/server/services/dailyreward"

	"github.com/gin-gonic/gin"
)

func RewardCollectedHandler(ctx *gin.Context) {
	playerID, exists := ctx.Get("playerId")
	fmt.Println("player id is :", playerID)
	if !exists {
		response.ErrorResponse(ctx, 401, "Unauthorised")
		return
	}
	dailyreward.RewardCollectedService(ctx, playerID.(string))
}
