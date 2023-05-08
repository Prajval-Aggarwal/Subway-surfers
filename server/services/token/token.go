package token

import (
	"fmt"
	"os"
	"subway/server/response"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// Token Payload
type Claims struct {
	P_Id string `json:"playerId"`
	jwt.RegisteredClaims
}

// Generate JWT Token
func GenerateToken(claims Claims, context *gin.Context) string {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("JWTKEY")))

	if err != nil {
		response.ErrorResponse(context, 401, "Error signing token")
	}
	return tokenString
}

// Decode Token function an checking whether the token is valid or not
func DecodeToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

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
