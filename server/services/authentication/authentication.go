package authentication

import (
	"errors"
	"fmt"
	"subway/server/db"
	"subway/server/model"
	"subway/server/request"
	"subway/server/response"
	"subway/server/services/token"
	"subway/server/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	gomail "gopkg.in/mail.v2"
	"gorm.io/gorm"
)

// Register Service registers the player and stores it into db
func RegisterService(ctx *gin.Context, registerRequest request.RegisterRequest) {
	var player model.Player
	var playerCoins model.PlayerCoins
	var avatar struct {
		AvatarId string
	}
	player.P_Name = registerRequest.P_Name
	player.Email = registerRequest.Email

	password := registerRequest.Password

	//Checking whether the passwrod meets the requirements
	err := utils.IsPassValid(password)
	if err != nil {
		response.ErrorResponse(ctx, 400, err.Error())
		return
	}

	// Hashing the password
	// bs, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	// if err != nil {
	// 	response.ErrorResponse(ctx, 500, "Unable to hash the password")
	// 	return
	// }
	// player.Password=string(bs)

	// selecting random avatar for the player that are unlocked
	query := "SELECT avatar_id FROM avatars WHERE status='Unlocked' ORDER BY RANDOM() LIMIT 1"
	db.RawQuery(query, avatar)
	player.CurrAvatar = avatar.AvatarId

	player.Password = password

	// Creating db record
	err = db.CreateRecord(&player)
	if err != nil {
		response.ErrorResponse(ctx, 500, err.Error())
		return
	}

	// Giving some coins to the at the time of registeration and storing into db
	playerCoins.P_ID = player.P_ID
	playerCoins.Coins = 1000
	err = db.CreateRecord(&playerCoins)
	if err != nil {
		response.ErrorResponse(ctx, utils.INTERNAL_SERVER_ERROR, err.Error())
		return
	}

	response.ShowResponse("Success", utils.CREATED, "Player registered successfully", &player, ctx)

}

// LoginService logs in the player
func LoginService(ctx *gin.Context, loginRequest request.LoginRequest) {
	expirationTime := time.Now().Add(time.Minute * 10)
	var playerDetails model.Player
	var tokenClaims token.Claims
	if !db.RecordExist("players", "email", loginRequest.Email) {
		response.ErrorResponse(ctx, utils.NOT_FOUND, "Player needs to register to proceed with login..")
		return
	}
	db.FindById(&playerDetails, loginRequest.Email, "email")

	// Comparing the password stored in db and entered password

	// err := bcrypt.CompareHashAndPassword([]byte(playerDetails.Password), []byte(loginRequest.Password))
	// if err != nil {
	// 	response.ErrorResponse(ctx, 401, "Unauthorised")
	// 	return
	// }

	if playerDetails.Password != loginRequest.Password {
		response.ErrorResponse(ctx, utils.UNAUTHORIZED, "Invalid credentials")
		return
	}

	//Creating token payload
	tokenClaims.P_Id = playerDetails.P_ID
	tokenClaims.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(expirationTime)

	// Gnerating JWT token for the player
	tokenString := token.GenerateToken(tokenClaims, ctx)

	//Creating session record
	session := model.Session{
		P_Id:  playerDetails.P_ID,
		Token: tokenString,
	}

	// If record exists inn session table then update the session token else create the session token
	if db.RecordExist("sessions", "p_id", playerDetails.P_ID) {
		err := db.UpdateRecord(session, playerDetails.P_ID, "p_id").Error
		if err != nil {
			response.ErrorResponse(ctx, utils.INTERNAL_SERVER_ERROR, err.Error())
			return
		}
	} else {
		err := db.CreateRecord(&session)
		if err != nil {
			response.ErrorResponse(ctx, utils.INTERNAL_SERVER_ERROR, err.Error())
			return
		}
	}

	response.ShowResponse("Sucess", utils.SUCCESS, "Login sucessfull", tokenString, ctx)

}

// LogoutService logs out
func LogoutService(ctx *gin.Context, playerId string) {

	// Finding the session of the player in session table and deleting it.
	var sessionDetails model.Session
	if !db.RecordExist("sessions", "p_id", playerId) {
		response.ErrorResponse(ctx, utils.NOT_FOUND, "Session for current user has already been ended")
		return
	}
	err := db.DeleteRecord(&sessionDetails, playerId, "p_id")
	if err != nil {
		response.ErrorResponse(ctx, utils.INTERNAL_SERVER_ERROR, err.Error())
		return
	}

}

// UpdatePasworService updates the password of the player
func UpdatePasswordService(ctx *gin.Context, password request.UpdatePasswordRequest, playerID string) {

	var playerDetails model.Player
	//Finding the player
	err := db.FindById(&playerDetails, playerID, "p_id")
	if err != nil {
		response.ErrorResponse(ctx, utils.BAD_REQUEST, err.Error())
		return
	}
	// Check if the password entered is same as previous oe
	if playerDetails.Password == password.Password {
		response.ErrorResponse(ctx, utils.BAD_REQUEST, "Password should be differnt from previous password")
		return
	}

	// Password validity check
	err = utils.IsPassValid(password.Password)
	if err != nil {
		response.ErrorResponse(ctx, utils.BAD_REQUEST, err.Error())
		return
	}

	//Hashing the new password
	// bs, err := bcrypt.GenerateFromPassword([]byte(password.Password), 14)
	// if err != nil {
	// 	response.ErrorResponse(ctx, 500, "Unable to hash the password")
	// 	return
	// }
	// playerDetails.Password = string(bs)
	playerDetails.Password = password.Password

	//Updating players new password
	err = db.UpdateRecord(&playerDetails, playerID, "p_id").Error
	if err != nil {
		response.ErrorResponse(ctx, utils.BAD_REQUEST, err.Error())
		return
	}
	response.ShowResponse("Success", utils.SUCCESS, "Password updated successfully", nil, ctx)

}

// UpdateNameService updates the player name
func UpdateNameService(ctx *gin.Context, playerName request.UpdateNameRequest, playerID string) {
	var playerDetails model.Player
	err := db.FindById(&playerDetails, playerID, "p_id")
	if err != nil {
		response.ErrorResponse(ctx, utils.BAD_REQUEST, err.Error())
		return
	}
	playerDetails.P_Name = playerName.P_Name

	err = db.UpdateRecord(&playerDetails, playerID, "p_id").Error
	if err != nil {
		response.ErrorResponse(ctx, utils.BAD_REQUEST, err.Error())
		return
	}
	response.ShowResponse("Success", utils.SUCCESS, "Player name updated successfully", nil, ctx)

}

// ForgotPassService sends an email to the user with reset link
func ForgotPassService(ctx *gin.Context, forgotPassRequest request.ForgotPassRequest) {
	expirationTime := time.Now().Add(time.Minute * 5)
	var playerDetails model.Player

	// finding the player email
	err := db.FindById(&playerDetails, forgotPassRequest.Email, "email")
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.ErrorResponse(ctx, utils.NOT_FOUND, "Email is not registered")
		return
	} else if err != nil {
		response.ErrorResponse(ctx, utils.BAD_REQUEST, err.Error())
		return
	}

	//Creating reset token payload and generating token form it
	resetClaims := token.Claims{
		P_Id: playerDetails.P_ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	tokenString := token.GenerateToken(resetClaims, ctx)

	// Creating reset session for reseting the password
	resetSession := model.ResetSession{
		P_ID:       playerDetails.P_ID,
		ResetToken: tokenString,
	}
	err = db.CreateRecord(&resetSession)

	if err != nil {
		response.ErrorResponse(ctx, utils.INTERNAL_SERVER_ERROR, err.Error())
		return
	}

	//Sending mail on players email address
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

// ResetPasswordService reset the password of the user
func ResetPasswordService(ctx *gin.Context, tokenString string, password request.UpdatePasswordRequest) {
	var resetSession model.ResetSession
	//Decoding the token and extracting require data
	claims, err := token.DecodeToken(tokenString)
	if err != nil {
		response.ErrorResponse(ctx, utils.BAD_REQUEST, err.Error())
		return
	}

	db.FindById(&resetSession, claims.P_Id, "p_id")

	if resetSession.ResetToken != tokenString {
		response.ErrorResponse(ctx, utils.FORBIDDEN, "Forbidden request")
		return
	}
	// Reusing he updatepasswordservice
	UpdatePasswordService(ctx, password, claims.P_Id)

	//After sucessfully reseting the password deleteing reset session of the player
	err = db.DeleteRecord(&resetSession, claims.P_Id, "p_id")
	if err != nil {
		response.ErrorResponse(ctx, utils.BAD_REQUEST, err.Error())
		return
	}
}
