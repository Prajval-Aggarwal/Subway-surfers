package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"subway/server/request"
	"subway/server/response"
	"subway/server/services/powerups"
	"subway/server/utils"
	"subway/server/validation"
)

//	@Description	Show Power ups
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	response.Success
//	@Failure		400	{object}	response.Error
//	@Tags			Powerups
//	@Router			/show-powerups [post]
func ShowPowerUpsHandler(ctx *gin.Context) {
	powerups.ShowPowerUpsService(ctx)
}

//	@Description	Details of the power up used
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	response.Success
//	@Failure		400		{object}	response.Error
//	@Param			Details	body		request.PowerUpRequest	true	"Power Up used details"
//	@Tags			Powerups
//	@Router			/use-powerup [post]
func UsePowerUpHandler(ctx *gin.Context) {
	playerID, exists := ctx.Get("playerId")
	fmt.Println("player id is :", playerID)
	if !exists {
		response.ErrorResponse(ctx, 401, "Unauthorised")
		return
	}
	var powerupRequest request.PowerUpRequest
	utils.RequestDecoding(ctx, &powerupRequest)
	err := validation.CheckValidation(&powerupRequest)
	if err != nil {
		response.ErrorResponse(ctx, 400, err.Error())
		return
	}

	powerups.UsePowerUpService(ctx, playerID.(string), powerupRequest)
}

//	@Description	Details of the power up bought
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	response.Success
//	@Failure		400		{object}	response.Error
//	@Param			Details	body		request.BuyRequest	true	"Power Up bought details"
//	@Tags			Powerups
//	@Router			/buy-powerup [post]
func BuyPowerupHandler(ctx *gin.Context) {
	playerID, exists := ctx.Get("playerId")
	fmt.Println("player id is :", playerID)
	if !exists {
		response.ErrorResponse(ctx, 401, "Unauthorised")
		return
	}

	var BuyRequest request.BuyRequest
	utils.RequestDecoding(ctx, &BuyRequest)
	//add validation

	err := validation.CheckValidation(&BuyRequest)
	if err != nil {
		response.ErrorResponse(ctx, 400, err.Error())
		return
	}

	powerups.BuyPowerupService(ctx, playerID.(string), BuyRequest)
}
