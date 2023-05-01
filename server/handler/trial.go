package handler

import (
	"fmt"
	"subway/server/response"

	"github.com/gin-gonic/gin"
)

func WelcomeHandler(ctx *gin.Context) {
	playerID, exists := ctx.Get("playerId")
	fmt.Println("player id is :", playerID)
	if !exists {
		response.ErrorResponse(ctx, 401, "Unauthorised")
		return
	}
	ctx.JSON(200, gin.H{
		"Text": "Welcome to the game",
	})

}
