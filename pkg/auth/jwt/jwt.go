package jwtAuth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/humamalamin/test-case-dating/pkg/config"
)

type Jwt interface {
	GenerateToken(data *JwtData) (string, string, int64, error)
	VerifyAccessToken(token string) (*JwtData, error)
	VerifyRefreshToken(token string) (string, error)
}

type Options struct {
	signinKey            string
	issuer               string
	accessTokenDuration  int
	refreshTokenDuration int
}

// GenerateToken implements Jwt
func (o *Options) GenerateToken(data *JwtData) (string, string, int64, error) {
	data.StandardClaims.ExpiresAt = time.Now().Local().Add(time.Second * time.Duration(o.accessTokenDuration)).Unix()
	acToken := jwt.NewWithClaims(jwt.SigningMethodHS512, data)
	accessToken, err := acToken.SignedString([]byte(o.signinKey))
	if err != nil {
		return "", "", 0, err
	}

	data.StandardClaims.ExpiresAt = time.Now().Local().Add(time.Second * time.Duration(o.refreshTokenDuration)).Unix()
	rfToken := jwt.NewWithClaims(jwt.SigningMethodHS512, data.StandardClaims)
	refreshToken, err := rfToken.SignedString([]byte(o.signinKey))
	if err != nil {
		return "", "", 0, err
	}

	return accessToken, refreshToken, data.StandardClaims.ExpiresAt, nil
}

// VerifyAccessToken implements Jwt
func (o *Options) VerifyAccessToken(token string) (*JwtData, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method invalid")
		}

		return []byte(o.signinKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, err
	}

	jwtData := &JwtData{
		UserID: claims["uid"].(string),
	}

	return jwtData, nil
}

// VerifyREfreshToken implements Jwt
func (o *Options) VerifyRefreshToken(token string) (string, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method invalid")
		}

		return []byte(o.signinKey), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return "", err
	}

	return claims["jti"].(string), nil
}

func NewJwt(cfg *config.Config) Jwt {
	opt := new(Options)
	opt.signinKey = cfg.JwtSigningKey
	opt.issuer = cfg.JwtIssuer
	opt.accessTokenDuration = cfg.JwtAccessTokenDuration
	opt.refreshTokenDuration = cfg.JwtRefreshTokenDuration
	return opt
}
