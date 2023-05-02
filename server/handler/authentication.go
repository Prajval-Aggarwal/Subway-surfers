package handler

import (
	"fmt"
	"subway/server/request"
	"subway/server/response"
	"subway/server/services/authentication"
	"subway/server/utils"
	"subway/server/validation"

	"github.com/gin-gonic/gin"
)

//	@Description	Register a player
//	@Accept			json
//	@Produce		json
//	@Success		200				{object}	response.Success
//	@Failure		400				{object}	response.Error
//	@Param			playerDetails	body		request.RegisterRequest	true	"Details of the player"
//	@Tags			Authentication
//	@Router			/register-player [post]
func RegisterHandler(ctx *gin.Context) {
	var registerRequest request.RegisterRequest

	utils.RequestDecoding(ctx, &registerRequest)

	err := validation.CheckValidation(&registerRequest)
	if err != nil {
		response.ErrorResponse(ctx, 400, err.Error())
		return
	}
	authentication.RegisterService(ctx, registerRequest)

}

//	@Description	Log in a player
//	@Accept			json
//	@Produce		json
//	@Success		200				{object}	response.Success
//	@Failure		400				{object}	response.Error
//	@Param			playerDetails	body		request.LoginRequest	true	"Details of the player"
//	@Tags			Authentication
//	@Router			/login [post]
func LoginHandler(ctx *gin.Context) {
	var loginRequest request.LoginRequest

	utils.RequestDecoding(ctx, &loginRequest)
	err := validation.CheckValidation(&loginRequest)
	if err != nil {
		response.ErrorResponse(ctx, 400, err.Error())
		return
	}
	authentication.LoginService(ctx, loginRequest)
}

//	@Description	Logs out a player
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	response.Success
//	@Failure		400	{object}	response.Error
//	@Tags			Authentication
//	@Router			/logout [delete]
func LogoutHandler(ctx *gin.Context) {

	playerID, exists := ctx.Get("playerId")
	fmt.Println("player id is :", playerID)
	if !exists {
		response.ErrorResponse(ctx, 401, "Unauthorised")
		return
	}

	authentication.LogoutService(ctx, playerID.(string))
}

//	@Description	Updates the password of the player
//	@Accept			json
//	@Produce		json
//	@Success		200			{object}	response.Success
//	@Failure		400			{object}	response.Error
//	@Param			newPassword	body		request.UpdatePasswordRequest	true	"New password of the player"
//	@Tags			Authentication
//	@Router			/update-pass [patch]
func UpdatePasswordHandler(ctx *gin.Context) {

	//get player id from the context that is passed from middleware
	playerID, exists := ctx.Get("playerId")
	fmt.Println("player id is :", playerID)
	if !exists {
		response.ErrorResponse(ctx, 401, "Unauthorised")
		return
	}
	var password request.UpdatePasswordRequest
	utils.RequestDecoding(ctx, &password)

	err := validation.CheckValidation(&password)
	if err != nil {
		response.ErrorResponse(ctx, 400, err.Error())
		return
	}
	authentication.UpdatePasswordService(ctx, password, playerID.(string))
}

//	@Description	Updates the player name of the player
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	response.Success
//	@Failure		400		{object}	response.Error
//	@Param			newName	body		request.UpdateNameRequest	true	"New name of the player"
//	@Tags			Authentication
//	@Router			/update-name [patch]
func UpdateNameHandler(ctx *gin.Context) {

	playerID, exists := ctx.Get("playerId")
	fmt.Println("player id is :", playerID)
	if !exists {
		response.ErrorResponse(ctx, 401, "Unauthorised")
		return
	}
	var playerName request.UpdateNameRequest
	utils.RequestDecoding(ctx, &playerName)

	err := validation.CheckValidation(&playerName)
	if err != nil {
		response.ErrorResponse(ctx, 400, err.Error())
		return
	}

	authentication.UpdateNameService(ctx, playerName, playerID.(string))
}

//	@Description	Forgot password
//	@Accept			json
//	@Produce		json
//	@Success		200			{object}	response.Success
//	@Failure		400			{object}	response.Error
//	@Param			playerEmail	body		request.ForgotPassRequest	true	"Players registers email"
//	@Tags			Authentication
//	@Router			/forgot-password [post]
func ForgotPasswordHandler(ctx *gin.Context) {
	var forgotRequest request.ForgotPassRequest
	utils.RequestDecoding(ctx, &forgotRequest)
	fmt.Println("forgot", forgotRequest)
	err := validation.CheckValidation(&forgotRequest)
	if err != nil {
		response.ErrorResponse(ctx, 400, err.Error())
		return
	}

	authentication.ForgotPassService(ctx, forgotRequest)
}

//	@Description	Reset password
//	@Accept			json
//	@Produce		json
//	@Success		200			{object}	response.Success
//	@Failure		400			{object}	response.Error
//	@Param			NewPassword	body		request.UpdatePasswordRequest	true	"Players new password"
//	@Tags			Authentication
//	@Router			/reset-password [post]
func ResetPasswordHandler(ctx *gin.Context) {
	tokenString := ctx.Request.URL.Query().Get("token")
	var password request.UpdatePasswordRequest

	utils.RequestDecoding(ctx, &password)

	err := validation.CheckValidation(&password)
	if err != nil {
		response.ErrorResponse(ctx, 400, err.Error())
		return
	}

	authentication.ResetPasswordService(ctx, tokenString, password)

}
