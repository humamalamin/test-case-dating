package middleware

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/humamalamin/test-case-dating/pkg/config"
	"github.com/rs/zerolog/log"

	jsonHelper "github.com/humamalamin/test-case-dating/helpers/json"
	jwtAuth "github.com/humamalamin/test-case-dating/pkg/auth/jwt"
)

type Middleware interface {
	InitLog(next http.Handler) http.Handler
	CheckToken(next http.Handler) http.Handler
	GetUserInformationContext(ctx context.Context) (*UserData, error)
}

type Options struct {
	jwt jwtAuth.Jwt
}

func (o *Options) getTokenHeader(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", ErrorAuthHeaderEmpty
	}

	accessToken := strings.Split(authHeader, " ")
	if accessToken[0] != "Bearer" {
		return "", ErrorAuthNotHaveBearer
	}

	if len(accessToken) == 1 {
		return "", ErrorAuthNotHaveToken
	}

	return accessToken[1], nil
}

// CheckToken implements Middleware
func (o *Options) CheckToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessToken, err := o.getTokenHeader(r)
		if err != nil {
			code := "[Middleware] CheckToken - 1"
			log.Error().Err(err).Msg(code)
			err = ErrorAuthHeaderEmpty
			jsonHelper.ErrorResponse(w, "4001", http.StatusUnauthorized, nil, err.Error())
			return
		}

		jwtData, err := o.jwt.VerifyAccessToken(accessToken)
		if err != nil {
			code := "[Middleware] CheckToken - 2"
			log.Error().Err(err).Msg(code)
			err = ErrorAuthHeaderEmpty
			jsonHelper.ErrorResponse(w, "4001", http.StatusUnauthorized, nil, err.Error())
			return
		}

		if jwtData == nil {
			code := "[Middleware] CheckToken - 3"
			log.Error().Err(err).Msg(code)
			err = ErrorAuthHeaderEmpty
			jsonHelper.ErrorResponse(w, "4001", http.StatusUnauthorized, nil, err.Error())
			return
		}

		ctx := context.WithValue(r.Context(), config.ContextKey("userInfo"), jwtData)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetUserInformationContext implements Middleware
func (o *Options) GetUserInformationContext(ctx context.Context) (*UserData, error) {
	userInfo, ok := ctx.Value(config.ContextKey("userInfo")).(*jwtAuth.JwtData)
	if !ok || userInfo == nil {
		log.Error().Err(ErrorUserFromContext)
		return nil, ErrorUserFromContext
	}

	userData := &UserData{
		UserID: userInfo.UserID,
	}

	return userData, nil
}

// InitLog implements Middleware
func (o *Options) InitLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bodyBytes, _ := ioutil.ReadAll(r.Body)
		r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		now := time.Now()

		ctx := context.WithValue(r.Context(), config.ContextKey("body"), bodyBytes)
		ctx = context.WithValue(ctx, config.ContextKey("startTime"), now)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func NewMiddleware(cfg *config.Config) Middleware {
	opt := new(Options)
	opt.jwt = jwtAuth.NewJwt(cfg)

	return opt
}
