package authentication

import (
	"errors"
	"fmt"
	"subway/server/db"
	"subway/server/model"
	"subway/server/provider"
	"subway/server/request"
	"subway/server/response"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	gomail "gopkg.in/mail.v2"
	"gorm.io/gorm"
)

func RegisterService(ctx *gin.Context, registerRequest request.RegisterRequest) {
	var player model.Player
	var playerCoins model.PlayerCoins
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

	//select random avatar for the user
	var avatar struct {
		AvatarId string
	}
	query := "SELECT avatar_id FROM avatars WHERE status='Unlocked' ORDER BY RANDOM() LIMIT 1"
	db.RawQuery(query, avatar)
	player.CurrAvatar = avatar.AvatarId

	player.Password = password
	err := db.CreateRecord(&player)
	if err != nil {
		response.ErrorResponse(ctx, 500, err.Error())
		return
	}
	playerCoins.P_ID = player.P_ID

	err = db.CreateRecord(&playerCoins)
	if err != nil {
		response.ErrorResponse(ctx, 500, err.Error())
		return
	}

	response.ShowResponse("Success", 201, "Player registered successfully", &player, ctx)

}

func LoginService(ctx *gin.Context, loginRequest request.LoginRequest) {
	var playerDetails model.Player
	var tokenClaims model.Claims
	if !db.RecordExist("players", "email", loginRequest.Email) {
		response.ErrorResponse(ctx, 404, "Player needs to register to proceed with login..")
		return
	}
	db.FindById(&playerDetails, loginRequest.Email, "email")

	//comapring password using bcrypt

	// err := bcrypt.CompareHashAndPassword([]byte(playerDetails.Password), []byte(loginRequest.Password))
	// if err != nil {
	// 	response.ErrorResponse(ctx, 401, "Unauthorised")
	// 	return
	// }

	if playerDetails.Password != loginRequest.Password {
		response.ErrorResponse(ctx, 401, "Invalid credentials")
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

	//if record exists then update the session token else ccreate the session token
	if db.RecordExist("sessions", "p_id", playerDetails.P_ID) {
		err := db.UpdateRecord(session, playerDetails.P_ID, "p_id").Error
		if err != nil {
			response.ErrorResponse(ctx, 500, err.Error())
			return
		}
	} else {
		err := db.CreateRecord(&session)
		if err != nil {
			response.ErrorResponse(ctx, 500, err.Error())
			return
		}
	}

	response.ShowResponse("Sucess", 200, "Login sucessfull", tokenString, ctx)
	//creating login record

}

func LogoutService(ctx *gin.Context, playerId string) {
	var sessionDetails model.Session
	if !db.RecordExist("sessions", "p_id", playerId) {
		response.ErrorResponse(ctx, 404, "Session for current user has already been ended")
		return
	}
	err := db.DeleteRecord(&sessionDetails, playerId, "p_id")
	if err != nil {
		response.ErrorResponse(ctx, 500, err.Error())
		return
	}

}

func UpdatePasswordService(ctx *gin.Context, password request.UpdatePasswordRequest, playerID string) {
	var playerDetails model.Player
	err := db.FindById(&playerDetails, playerID, "p_id")
	if err != nil {
		response.ErrorResponse(ctx, 400, err.Error())
		return
	}
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

	err = db.UpdateRecord(&playerDetails, playerID, "p_id").Error
	if err != nil {
		response.ErrorResponse(ctx, 400, err.Error())
		return
	}
	response.ShowResponse("Success", 200, "Password updated successfully", nil, ctx)

}

func UpdateNameService(ctx *gin.Context, playerName request.UpdateNameRequest, playerID string) {
	var playerDetails model.Player
	err := db.FindById(&playerDetails, playerID, "p_id")
	if err != nil {
		response.ErrorResponse(ctx, 400, err.Error())
		return
	}
	playerDetails.P_Name = playerName.P_Name

	err = db.UpdateRecord(&playerDetails, playerID, "p_id").Error
	if err != nil {
		response.ErrorResponse(ctx, 400, err.Error())
		return
	}
	response.ShowResponse("Success", 200, "Player name updated successfully", nil, ctx)

}

func ForgotPassService(ctx *gin.Context, forgotPassRequest request.ForgotPassRequest) {
	expirationTime := time.Now().Add(time.Minute * 5)
	var playerDetails model.Player
	err := db.FindById(&playerDetails, forgotPassRequest.Email, "email")
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.ErrorResponse(ctx, 404, "Email is not registered")
		return
	} else if err != nil {
		response.ErrorResponse(ctx, 400, err.Error())
		return
	}

	resetClaims := model.Claims{
		P_Id: playerDetails.P_ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	tokenString := provider.GenerateToken(resetClaims, ctx)

	resetSession := model.ResetSession{
		P_ID:       playerDetails.P_ID,
		ResetToken: tokenString,
	}
	err = db.CreateRecord(&resetSession)

	if err != nil {
		response.ErrorResponse(ctx, 500, err.Error())
		return
	}

	//send mail
	m := gomail.NewMessage()
	m.SetHeader("From", "prajwal1711@gmail.com")

	m.SetHeader("To", "namanagg59@gmail.com")
	m.SetHeader("Subject", "Reset Password link")
	link := "http://localhost:3000/reset-password?token=" + tokenString
	body := fmt.Sprintf("<a href=\"%s\">Click here to reset your password</a>", link)

	m.SetBody("text/html", body)

	d := gomail.NewDialer("smtp.gmail.com", 587, "prajwal1711@gmail.com", "agovqanwcgnewxmt")

	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}

}

func ResetPasswordService(ctx *gin.Context, tokenString string, password request.UpdatePasswordRequest) {

	//decode the token and get the playerid
	claims, err := provider.DecodeToken(tokenString)
	if err != nil {
		response.ErrorResponse(ctx, 400, err.Error())
		return
	}
	var resetSession model.ResetSession
	db.FindById(&resetSession, claims.P_Id, "p_id")

	if resetSession.ResetToken != tokenString {
		response.ErrorResponse(ctx, 403, "Forbidden request")
		return
	}
	//update password
	UpdatePasswordService(ctx, password, claims.P_Id)

	//delete the record from reset session table

	err = db.DeleteRecord(&resetSession, claims.P_Id, "p_id")
	if err != nil {
		response.ErrorResponse(ctx, 400, err.Error())
		return
	}
}
