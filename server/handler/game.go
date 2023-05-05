package handler

import (
	"subway/server/request"
	"subway/server/response"
	"subway/server/services/game"
	"subway/server/utils"
	"subway/server/validation"

	"github.com/gin-gonic/gin"
)

// @Description	Ends the game
// @Accept			json
// @Produce		json
// @Success		200			{object}	response.Success
// @Failure		400			{object}	response.Error
// @Param			gameDetails	body		request.EndGameRequest	true	"Players record after game end"
// @Tags			Game
// @Router			/end-game [post]
func EndGameHandler(ctx *gin.Context) {
	// playerID, exists := ctx.Get("playerId")
	// fmt.Println("player id is :", playerID)
	// if !exists {
	// 	response.ErrorResponse(ctx, 401, "Unauthorised")
	// 	return
	// }
	var endGameRequest request.EndGameRequest
	utils.RequestDecoding(ctx, &endGameRequest)

	err := validation.CheckValidation(&endGameRequest)
	if err != nil {
		response.ErrorResponse(ctx, 400, err.Error())
		return
	}

	game.EndGameService(ctx, "123", endGameRequest)
}
