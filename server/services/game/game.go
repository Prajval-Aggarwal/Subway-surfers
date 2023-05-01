package game

import (
	"math"
	"subway/server/db"
	"subway/server/model"
	"subway/server/request"
	"subway/server/response"

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
