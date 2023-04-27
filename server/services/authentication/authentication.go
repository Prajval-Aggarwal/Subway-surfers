package authentication

import (
	"subway/server/db"
	"subway/server/model"
	"subway/server/provider"
	"subway/server/request"
	"subway/server/response"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RegisterService(ctx *gin.Context, registerRequest request.RegisterRequest) {
	var player model.Player
	player.P_Name = registerRequest.P_Name
	player.Email = registerRequest.Email

	password := registerRequest.Password

	//using bcrypt
	// bs, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	// if err != nil {
	// 	response.ErrorResponse(ctx, 500, "Unable to hash the password")
	// 	return
	// }
	// player.Password=string(bs)

	player.Password = password
	err := db.CreateRecord(&player)
	if err != nil {
		response.ErrorResponse(ctx, 500, err.Error())
		return
	}

	response.ShowResponse("Success", 201, "Player registered successfully", &player, ctx)

}

func LoginService(ctx *gin.Context, loginRequest request.LoginRequest) {
	var playerDetails model.Player
	var tokenClaims model.Claims
	db.FindById(&playerDetails, loginRequest.Email, "emial")

	//comapring password using bcrypt

	// err := bcrypt.CompareHashAndPassword([]byte(playerDetails.Password), []byte(loginRequest.Password))
	// if err != nil {
	// 	response.ErrorResponse(ctx, 401, "Unauthorised")
	// 	return
	// }

	if playerDetails.Password != loginRequest.Password {
		response.ErrorResponse(ctx, 401, "Unauthorised")
		return
	}
	expirationTime := time.Now().Add(time.Minute * 10)
	tokenClaims.P_Id = playerDetails.P_ID
	tokenClaims.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(expirationTime)

	//genertaing token
	tokenString := provider.GenerateToken(tokenClaims, ctx)

	//creating session record

	session := model.Session{
		P_Id:  playerDetails.P_ID,
		Token: tokenString,
	}
	err := db.CreateRecord(&session)
	if err != nil {

		response.ErrorResponse(ctx, 500, err.Error())
		return
	}
	//creating login record

}

func LogoutService(ctx *gin.Context, logoutRequest request.LogoutRequest) {
	var sessionDetails model.Session
	if !db.RecordExist("sessions", "p_id", logoutRequest.P_Id) {
		response.ErrorResponse(ctx, 404, "Session for current user has already been ended")
		return
	}
	err := db.DeleteRecord(&sessionDetails, logoutRequest.P_Id, "p_id")
	if err != nil {
		response.ErrorResponse(ctx, 500, err.Error())
		return
	}

}

func UpdatePasswordService(ctx *gin.Context, password request.UpdatePasswordRequest, playerID string) {
	var playerDetails model.Player
	db.FindById(&playerDetails, playerID, "p_id")
	if playerDetails.Password == password.Password {
		response.ErrorResponse(ctx, 400, "Password should be differnt from previous password")
		return
	}
	//using bcrypt
	// bs, err := bcrypt.GenerateFromPassword([]byte(password.Password), 14)
	// if err != nil {
	// 	response.ErrorResponse(ctx, 500, "Unable to hash the password")
	// 	return
	// }
	// playerDetails.Password = string(bs)
	playerDetails.Password = password.Password

	err := db.UpdateRecord(&playerDetails, playerID, "p_id").Error
	if err != nil {
		response.ErrorResponse(ctx, 400, err.Error())
		return
	}
	response.ShowResponse("Success", 200, "Password updated successfully", nil, ctx)

}

func UpdateNameService(ctx *gin.Context, playerName request.UpdateNameRequest, playerID string) {
	var playerDetails model.Player
	db.FindById(&playerDetails, playerID, "p_id")
	playerDetails.P_Name = playerName.P_Name

	err := db.UpdateRecord(&playerDetails, playerID, "p_id").Error
	if err != nil {
		response.ErrorResponse(ctx, 400, err.Error())
		return
	}
	response.ShowResponse("Success", 200, "Player name updated successfully", nil, ctx)

}
