package handler

import (
	"fmt"
	"subway/server/response"
	"subway/server/services/player"

	"github.com/gin-gonic/gin"
)

func ShowPlayerDetailsHandler(ctx *gin.Context) {
	playerID, exists := ctx.Get("playerId")
	fmt.Println("player id is :", playerID)
	if !exists {
		response.ErrorResponse(ctx, 401, "Unauthorised")
		return
	}
	player.ShowPlayerDetailsService(ctx, playerID.(string))
}
