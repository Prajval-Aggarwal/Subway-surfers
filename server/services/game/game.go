package game

import (
	"fmt"
	"math"
	"subway/server/db"
	"subway/server/model"
	"subway/server/request"
	"subway/server/response"
	"subway/server/services/avatar"

	"github.com/gin-gonic/gin"
)

func EndGameService(ctx *gin.Context, playerId string, endGameRequest request.EndGameRequest) {

	var playerCoins model.PlayerCoins
	var playerDetails model.Player

	err := db.FindById(&playerCoins, playerId, "p_id")
	if err != nil {
		response.ErrorResponse(ctx, 400, err.Error())
		return
	}

	err = db.FindById(&playerDetails, playerId, "p_id")
	if err != nil {
		response.ErrorResponse(ctx, 400, err.Error())
		return
	}

	if endGameRequest.Distance > 100000 {
		fmt.Println("unlock avatar function called")
		avatar.UnlockAvtar(ctx, playerId, endGameRequest.Distance)
	}

	//upadate player coins
	playerCoins.Coins += endGameRequest.CoinsCollected

	err = db.UpdateRecord(&playerCoins, playerId, "p_id").Error
	if err != nil {
		response.ErrorResponse(ctx, 400, err.Error())
		return
	}

	//update player total distance and high score

	playerDetails.HighScore = int64(math.Max(float64(playerDetails.HighScore), float64(endGameRequest.Distance)))
	playerDetails.TotalDistance += endGameRequest.Distance

	err = db.UpdateRecord(&playerDetails, playerId, "p_id").Error
	if err != nil {
		response.ErrorResponse(ctx, 400, err.Error())
		return
	}

	response.ShowResponse("Success", 200, "Game ended successfully", nil, ctx)
}
