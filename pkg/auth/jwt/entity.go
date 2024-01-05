package jwtAuth

import "github.com/golang-jwt/jwt"

type JwtData struct {
	UserID string `json:"uid"`
	jwt.StandardClaims
}
