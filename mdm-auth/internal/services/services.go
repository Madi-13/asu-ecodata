package services

import (
	"mdm-auth/internal/providers"
	"mdm-auth/internal/repositories"
)

type Deps struct {
	Providers    *providers.Providers
	Repositories *repositories.Repositories
}

type Services struct {
	AuthService AuthService
}

func NewServices(deps *Deps) *Services {
	auth := NewAuthService(deps.Providers.KeycloakProvider, deps.Repositories.Auth)
	return &Services{
		AuthService: auth,
	}
}
