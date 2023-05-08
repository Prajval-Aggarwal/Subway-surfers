package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"subway/server/request"
	"subway/server/response"
	"subway/server/services/powerups"
	"subway/server/utils"
)

// @Description	Show Power ups
// @Accept			json
// @Produce		json
// @Success		200	{object}	response.Success
// @Failure		400	{object}	response.Error
// @Tags			Powerups
// @Router			/show-powerups [post]
func ShowPowerUpsHandler(ctx *gin.Context) {
	powerups.ShowPowerUpsService(ctx)
}

// @Description	Details of the power up used
// @Accept			json
// @Produce		json
// @Success		200		{object}	response.Success
// @Failure		400		{object}	response.Error
// @Failure		401	{object}	response.Error
// @Param			Details	body		request.PowerUpRequest	true	"Power Up used details"
// @Tags			Powerups
// @Router			/use-powerup [post]
func UsePowerUpHandler(ctx *gin.Context) {
	playerID, exists := ctx.Get("playerId")
	fmt.Println("player id is :", playerID)
	if !exists {
		response.ErrorResponse(ctx, utils.UNAUTHORIZED, "Unauthorised")
		return
	}
	var powerupRequest request.PowerUpRequest
	err := utils.RequestDecoding(ctx, &powerupRequest)
	if err != nil {
		response.ErrorResponse(ctx, utils.BAD_REQUEST, err.Error())
		return
	}
	err = powerupRequest.Validate()
	if err != nil {
		response.ErrorResponse(ctx, utils.BAD_REQUEST, err.Error())
		return
	}

	powerups.UsePowerUpService(ctx, playerID.(string), powerupRequest)
}

// @Description	Details of the power up bought
// @Accept			json
// @Produce		json
// @Success		200		{object}	response.Success
// @Failure		400		{object}	response.Error
// @Failure		401	{object}	response.Error
// @Param			Details	body		request.BuyRequest	true	"Power Up bought details"
// @Tags			Powerups
// @Router			/buy-powerup [post]
func BuyPowerupHandler(ctx *gin.Context) {
	playerID, exists := ctx.Get("playerId")
	fmt.Println("player id is :", playerID)
	if !exists {
		response.ErrorResponse(ctx, utils.UNAUTHORIZED, "Unauthorised")
		return
	}

	var buyRequest request.BuyRequest
	err := utils.RequestDecoding(ctx, &buyRequest)
	if err != nil {
		response.ErrorResponse(ctx, utils.BAD_REQUEST, err.Error())
		return
	}

	err = buyRequest.Validate()
	if err != nil {
		response.ErrorResponse(ctx, utils.BAD_REQUEST, err.Error())
		return
	}
	powerups.BuyPowerupService(ctx, playerID.(string), buyRequest)
}
