package game

import (
	"fmt"
	"math"
	"subway/server/db"
	"subway/server/model"
	"subway/server/request"
	"subway/server/response"
	"subway/server/services/avatar"
	"subway/server/utils"

	"github.com/gin-gonic/gin"
)

//EndGameService engs the game for th player
func EndGameService(ctx *gin.Context, playerId string, endGameRequest request.EndGameRequest) {

	var playerCoins model.PlayerCoins
	var playerDetails model.Player

	err := db.FindById(&playerCoins, playerId, "p_id")
	if err != nil {
		response.ErrorResponse(ctx, utils.BAD_REQUEST, err.Error())
		return
	}

	err = db.FindById(&playerDetails, playerId, "p_id")
	if err != nil {
		response.ErrorResponse(ctx, utils.BAD_REQUEST, err.Error())
		return
	}

	//If the distance covered by the player is greater that 100000 then only unlock avatar
	if endGameRequest.Distance > 100000 {
		fmt.Println("unlock avatar function called")
		avatar.UnlockAvtar(ctx, playerId, endGameRequest.Distance)
	}

	//Upadting player's coins
	playerCoins.Coins += endGameRequest.CoinsCollected

	err = db.UpdateRecord(&playerCoins, playerId, "p_id").Error
	if err != nil {
		response.ErrorResponse(ctx, utils.BAD_REQUEST, err.Error())
		return
	}

	//Updating player total distance coverred and high score till now

	playerDetails.HighScore = int64(math.Max(float64(playerDetails.HighScore), float64(endGameRequest.Distance)))
	playerDetails.TotalDistance += endGameRequest.Distance

	err = db.UpdateRecord(&playerDetails, playerId, "p_id").Error
	if err != nil {
		response.ErrorResponse(ctx, utils.BAD_REQUEST, err.Error())
		return
	}

	response.ShowResponse("Success", utils.SUCCESS, "Game ended successfully", nil, ctx)
}
