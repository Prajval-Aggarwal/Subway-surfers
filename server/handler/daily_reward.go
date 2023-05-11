package handler

import (
	"fmt"
	"subway/server/response"
	"subway/server/services/dailyreward"
	"subway/server/utils"

	"github.com/gin-gonic/gin"
)

//	@Description	Collet reward handler
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	response.Success
//	@Failure		400	{object}	response.Error
//	@Failure		401	{object}	response.Error
//	@Tags			Daily Reward
//	@Router			/collect-reward [get]
func RewardCollectedHandler(ctx *gin.Context) {
	playerID, exists := ctx.Get("playerId")
	fmt.Println("player id is :", playerID)
	if !exists {
		response.ErrorResponse(ctx, utils.UNAUTHORIZED, "Unauthorised")
		return
	}
	dailyreward.RewardCollectedService(ctx, playerID.(string))
}

//	@Description	Shows the reward or the day of the player
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	response.Success
//	@Failure		400	{object}	response.Error
//	@Failure		401	{object}	response.Error
//	@Tags			Daily Reward
//	@Router			/show-reward [get]
func ShowPlayerRewardHandler(ctx *gin.Context) {
	playerID, exists := ctx.Get("playerId")
	fmt.Println("player id is :", playerID)
	if !exists {
		response.ErrorResponse(ctx, utils.UNAUTHORIZED, "Unauthorised")
		return
	}
	dailyreward.ShowPlayerRewardService(ctx, playerID.(string))

}
