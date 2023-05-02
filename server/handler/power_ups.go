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

func ShowPowerUpsHandler(ctx *gin.Context) {
	powerups.ShowPowerUpsService(ctx)
}

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
