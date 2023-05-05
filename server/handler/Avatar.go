package handler

import (
	"subway/server/request"
	"subway/server/response"
	"subway/server/services/avatar"
	"subway/server/utils"

	"github.com/gin-gonic/gin"
)

func ShowAvatarHandler(ctx *gin.Context) {
	avatar.ShowAvatarService(ctx)
}

func UpdateAvatarHandler(ctx *gin.Context) {
	var avatarRequest request.UpdateAvatarRequest
	// playerID, exists := ctx.Get("playerId")
	// fmt.Println("player id is :", playerID)
	// if !exists {
	// 	response.ErrorResponse(ctx, 401, "Unauthorised")
	// 	return
	// }
	err := utils.RequestDecoding(ctx, &avatarRequest)
	if err != nil {
		response.ErrorResponse(ctx, 400, err.Error())
		return
	}
	err = avatarRequest.Validate()
	if err != nil {
		response.ErrorResponse(ctx, 400, err.Error())
		return
	}
	avatar.UpdateAvatarService(ctx, "123", avatarRequest)
}
