package handler

import (
	"fmt"
	"subway/server/request"
	"subway/server/response"
	"subway/server/services/avatar"
	"subway/server/utils"

	"github.com/gin-gonic/gin"
)

//	@Description	Show the list of avatars
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	response.Success
//	@Failure		400	{object}	response.Error
//	@Tags			Avatar
//	@Router			/show-avatars [get]
func ShowAvatarHandler(ctx *gin.Context) {
	avatar.ShowAvatarService(ctx)
}

//	@Description	Updates the avatar for the player
//	@Accept			json
//	@Produce		json
//	@Success		200				{object}	response.Success
//	@Failure		400				{object}	response.Error
//	@Failure		401				{object}	response.Error
//	@Param			newAvatarName	body		request.UpdateAvatarRequest	true	"Id of the new avatar"
//	@Tags			Avatar
//	@Router			/update-avatar [patch]
func UpdateAvatarHandler(ctx *gin.Context) {
	var avatarRequest request.UpdateAvatarRequest
	playerID, exists := ctx.Get("playerId")
	fmt.Println("player id is :", playerID)
	if !exists {
		response.ErrorResponse(ctx, utils.UNAUTHORIZED, "Unauthorised")
		return
	}
	err := utils.RequestDecoding(ctx, &avatarRequest)
	if err != nil {
		response.ErrorResponse(ctx, utils.BAD_REQUEST, err.Error())
		return
	}
	err = avatarRequest.Validate()
	if err != nil {
		response.ErrorResponse(ctx, utils.BAD_REQUEST, err.Error())
		return
	}
	avatar.UpdateAvatarService(ctx, playerID.(string), avatarRequest)
}
