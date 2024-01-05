package auth

import (
	"github.com/gorilla/mux"
	"github.com/humamalamin/test-case-dating/pkg/manager"
)

func NewRoutes(r *mux.Router, mgr manager.Manager) {
	authHandler := NewAuthHandler(mgr)
	apiAuth := r.PathPrefix("/v1/auth").Subrouter()

	apiAuth.Handle("/register", authHandler.Register()).Methods("POST")
	apiAuth.Handle("/login", authHandler.Login()).Methods("POST")
}
