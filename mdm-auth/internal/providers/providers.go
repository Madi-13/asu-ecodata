package providers

import (
	"mdm-auth/config"
)

type Providers struct {
	KeycloakProvider KeycloakProvider
}

func NewProviders(cfg config.Keycloak) (*Providers, error) {
	keycloakProvider := NewKeycloakDeps(cfg.Host, cfg.Realm, cfg.Client, cfg.Secret, cfg.PublicKey)

	return &Providers{
		KeycloakProvider: keycloakProvider,
	}, nil
}
