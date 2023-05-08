package handler

import (
	"fmt"
	"subway/server/response"
	"subway/server/services/player"
	"subway/server/utils"

	"github.com/gin-gonic/gin"
)

// @Description	Show player details
// @Accept			json
// @Produce		json
// @Success		200	{object}	response.Success
// @Failure		400	{object}	response.Error
// @Failure		401	{object}	response.Error
// @Tags			Player
// @Router			/show-player [get]
func ShowPlayerDetailsHandler(ctx *gin.Context) {
	playerID, exists := ctx.Get("playerId")
	fmt.Println("player id is :", playerID)
	if !exists {
		response.ErrorResponse(ctx, utils.UNAUTHORIZED, "Unauthorised")
		return
	}
	player.ShowPlayerDetailsService(ctx, playerID.(string))
}
