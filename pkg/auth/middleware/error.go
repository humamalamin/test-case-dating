package middleware

import "errors"

var (
	ErrorAuthHeaderEmpty   = errors.New("Invalida Access Token")
	ErrorAuthNotHaveBearer = errors.New("Authorization header doesn't have bearer format")
	ErrorAuthNotHaveToken  = errors.New("Authorization header doesn't have access token value")
	ErrorUserFromContext   = errors.New("failed getting user info from context")
)
