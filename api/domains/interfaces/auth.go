package interfaces

import (
	"context"
	"net/http"

	authEntity "github.com/humamalamin/test-case-dating/api/domains/entities/auth"
)

type AuthHandler interface {
	Register() http.Handler
	Login() http.Handler
}

type AuthRepository interface {
	Register(ctx context.Context, req *authEntity.Auth) error
	Login(ctx context.Context, req *authEntity.Auth) (*authEntity.Auth, error)
}

type AuthService interface {
	Register(ctx context.Context, req *authEntity.Auth) error
	Login(ctx context.Context, req *authEntity.Auth) (*authEntity.AccessToken, error)
}
