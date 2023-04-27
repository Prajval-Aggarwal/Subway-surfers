package model

import "github.com/golang-jwt/jwt/v4"

type Claims struct {
	P_Id string `json:"playerId"`
	jwt.RegisteredClaims
}
