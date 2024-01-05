package services

import (
	"context"
	"errors"

	"github.com/golang-jwt/jwt"
	"github.com/humamalamin/test-case-dating/api/domains/entities/auth"
	"github.com/humamalamin/test-case-dating/api/domains/interfaces"
	"github.com/humamalamin/test-case-dating/api/repositories"
	"github.com/humamalamin/test-case-dating/helpers/conv"
	jwtAuth "github.com/humamalamin/test-case-dating/pkg/auth/jwt"
	"github.com/humamalamin/test-case-dating/pkg/config"
	"github.com/humamalamin/test-case-dating/pkg/manager"
	"github.com/rs/zerolog/log"
)

type Auth struct {
	repo interfaces.AuthRepository
	cfg  *config.Config
	jwt  jwtAuth.Jwt
}

var (
	code string
)

// Login implements interfaces.AuthService.
func (service *Auth) Login(ctx context.Context, req *auth.Auth) (*auth.AccessToken, error) {
	result, err := service.repo.Login(ctx, req)
	if err != nil {
		code := "[Service] Login - 1"
		log.Error().Err(err).Msg(code)
		return nil, err
	}

	if result == nil {
		code := "[Service] Login - 2"
		log.Error().Err(err).Msg(code)
		err = errors.New("data empty")
		return nil, err
	}

	password := conv.HashShaPassword(req.Password, service.cfg.SaltPassword)
	hash, _ := conv.HasPassword(password)
	match := conv.CheckPasswordHash(result.Password, hash)
	if !match {
		code := "[Service] Login - 3"
		log.Error().Err(err).Msg(code)
		err = errors.New("password / email invalid")
		return nil, err
	}

	response, err := service.generateToken(ctx, result)
	if err != nil {
		code := "[Service] Login - 4"
		log.Error().Err(err).Msg(code)
		err = errors.New("failed generate token")
		return nil, err
	}

	return response, nil
}

// Register implements interfaces.AuthService.
func (service *Auth) Register(ctx context.Context, req *auth.Auth) error {
	password := conv.HashShaPassword(req.Password, service.cfg.SaltPassword)
	req.Password = password

	err := service.repo.Register(ctx, req)
	if err != nil {
		code = "[Service] Register - 1"
		log.Error().Err(err).Msg(code)
		return err
	}

	return nil
}

func (service *Auth) generateToken(ctx context.Context, req *auth.Auth) (*auth.AccessToken, error) {
	jwtData := &jwtAuth.JwtData{
		UserID: req.ID,
		StandardClaims: jwt.StandardClaims{
			Id:        req.ID,
			NotBefore: jwt.TimeFunc().Local().Unix(),
		},
	}

	accessToken, refreshToken, expiredAt, err := service.jwt.GenerateToken(jwtData)
	if err != nil {
		code := "[Service] GenerateToken - 1"
		log.Error().Err(err).Msg(code)
		return nil, err
	}

	token := auth.AccessToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiredAt:    expiredAt,
	}

	return &token, nil
}

func NewAuthService(mng manager.Manager) interfaces.AuthService {
	service := new(Auth)
	service.jwt = mng.GetJwt()
	service.cfg = mng.GetConfig()
	service.repo = repositories.NewAuthRepository(mng.GetDatabase())

	return service
}
