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

func LogoutHandler(ctx *gin.Context) {

	var logoutRequest request.LogoutRequest

	utils.RequestDecoding(ctx, &logoutRequest)

	err := validation.CheckValidation(&logoutRequest)
	if err != nil {
		response.ErrorResponse(ctx, 400, err.Error())
		return
	}
	authentication.LogoutService(ctx, logoutRequest)
}

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

//Add forgot password handler
