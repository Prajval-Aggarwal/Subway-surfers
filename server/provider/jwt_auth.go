package provider

import (
	"fmt"
	"os"
	"subway/server/model"
	"subway/server/response"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// Generate JWT Token
func GenerateToken(claims model.Claims, context *gin.Context) string {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("JWTKEY")))

	if err != nil {
		response.ErrorResponse(context, 401, "Error signing token")
	}
	return tokenString
}

// Decode Token function
func DecodeToken(tokenString string) (*model.Claims, error) {
	claims := &model.Claims{}

	parsedToken, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("error")
		}
		return []byte(os.Getenv("JWTKEY")), nil
	})

	if err != nil || !parsedToken.Valid {
		return nil, fmt.Errorf("invalid or expired token")
	}

	return claims, nil
}

// Set cookie handler
// func SetCookie(context *gin.Context, tokenString string) {

// 	context.SetCookie(
// 		"cookie",
// 		tokenString,
// 		7200,
// 		"/",
// 		"localhost",
// 		false,
// 		true,
// 	)

// 	response.ShowResponse(
// 		"Success",
// 		200,
// 		"Cookies saved successfully",
// 		"",
// 		context,
// 	)
// }

// // Delete cookie handler
// func DeleteCookie(context *gin.Context) {
// 	context.SetCookie(
// 		"cookie",
// 		"",
// 		-1,
// 		"",
// 		"",
// 		false,
// 		false,
// 	)

// 	response.ShowResponse(
// 		"Success",
// 		200,
// 		"Cookie deleted successfully",
// 		"",
// 		context,
// 	)
// }
