package graph

import (
	"auth-system/internal/auth"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	AuthService *auth.AuthService
}

func NewResolver() *Resolver {
	return &Resolver{
		AuthService: auth.NewAuthService(),
	}
}
