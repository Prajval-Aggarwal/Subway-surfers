package handler

import (
	"subway/server/request"
	"subway/server/services/avatar"
	"subway/server/utils"

	"github.com/gin-gonic/gin"
)

func ShowAvatarHandler(ctx *gin.Context) {
	avatar.ShowAvatarService(ctx)
}

func UpdateAvatarHandler(ctx *gin.Context) {
	var avatarRequest request.AvatarRequest
	// playerID, exists := ctx.Get("playerId")
	// fmt.Println("player id is :", playerID)
	// if !exists {
	// 	response.ErrorResponse(ctx, 401, "Unauthorised")
	// 	return
	// }
	utils.RequestDecoding(ctx, &avatarRequest)
	avatar.UpdateAvatarService(ctx, "123", avatarRequest)
}
