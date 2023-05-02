package handler

import (
	"fmt"
	"subway/server/response"
	"subway/server/services/dailyreward"

	"github.com/gin-gonic/gin"
)

//	@Description	Collet reward handler
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	response.Success
//	@Failure		400	{object}	response.Error
//	@Tags			Daily Reward
//	@Router			/collect-reward [get]
func RewardCollectedHandler(ctx *gin.Context) {
	playerID, exists := ctx.Get("playerId")
	fmt.Println("player id is :", playerID)
	if !exists {
		response.ErrorResponse(ctx, 401, "Unauthorised")
		return
	}
	dailyreward.RewardCollectedService(ctx, playerID.(string))
}
