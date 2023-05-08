package gateway

import (
	"subway/server/db"
	"subway/server/model"
	"subway/server/response"
	"subway/server/services/token"
	"subway/server/utils"

	"github.com/gin-gonic/gin"
)

func PlayerAuthentication(ctx *gin.Context) {
	tokenString, err := utils.GetTokenFromAuthHeader(ctx)
	var playerSession model.Session
	if err != nil {
		response.ErrorResponse(ctx, 404, err.Error())
		ctx.Abort()
		return
	}

	claims, err := token.DecodeToken(tokenString)
	if err != nil {
		response.ErrorResponse(ctx, 401, err.Error())
		ctx.Abort()
		return
	}

	if !db.RecordExist("sessions", "p_id", claims.P_Id) {
		response.ErrorResponse(ctx, 401, "Unauthorized")
		ctx.Abort()
		return
	}

	db.FindById(&playerSession, claims.P_Id, "p_id")

	if tokenString != playerSession.Token {
		response.ErrorResponse(ctx, 401, "Unauthorized")
		ctx.Abort()
		return
	}
	ctx.Set("playerId", claims.P_Id)
	ctx.Next()
}
