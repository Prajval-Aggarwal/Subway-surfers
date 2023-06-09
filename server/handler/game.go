package handler

import (
	"fmt"
	"subway/server/request"
	"subway/server/response"
	"subway/server/services/game"
	"subway/server/utils"

	"github.com/gin-gonic/gin"
)

//	@Description	Ends the game
//	@Accept			json
//	@Produce		json
//	@Success		200			{object}	response.Success
//	@Failure		400			{object}	response.Error
//	@Failure		401			{object}	response.Error
//	@Param			gameDetails	body		request.EndGameRequest	true	"Players record after game end"
//	@Tags			Game
//	@Router			/end-game [post]
func EndGameHandler(ctx *gin.Context) {
	playerID, exists := ctx.Get("playerId")
	fmt.Println("player id is :", playerID)
	if !exists {
		response.ErrorResponse(ctx, utils.UNAUTHORIZED, "Unauthorised")
		return
	}
	var endGameRequest request.EndGameRequest
	err := utils.RequestDecoding(ctx, &endGameRequest)
	if err != nil {
		response.ErrorResponse(ctx, utils.BAD_REQUEST, err.Error())
		return
	}
	err = endGameRequest.Validate()
	if err != nil {
		response.ErrorResponse(ctx, utils.BAD_REQUEST, err.Error())
		return
	}

	game.EndGameService(ctx, playerID.(string), endGameRequest)
}
