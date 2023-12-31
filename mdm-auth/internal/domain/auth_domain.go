package domain

import "encoding/json"

const (
	InternalError        = `Internal Error`
	BadRequest           = `Invalid request body`
	AccessTokenRequired  = `Bearer access token is required`
	RefreshTokenRequired = `Bearer refresh token is required`
	BasicRequired        = `Basic authorization is required`
	NotValidCreds        = `Credentials are not valid`
	NotValidToken        = `Token is not valid`
	TokenValidateFail    = `Failed to validate token`
	UserIdRequired       = `user-id is required`
	ClientIdRequired     = `client-id is required`
	RoleNameRequired     = `role-name is required`
)

type AuthResponse struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int    `json:"expires_in"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
	RefreshToken     string `json:"refresh_token"`
	TokenType        string `json:"token_type"`
	NotBeforePolicy  int    `json:"not-before-policy"`
	SessionState     string `json:"session_state"`
	Scope            string `json:"scope"`
}

type ValidResponse struct {
	Exp               int         `json:"exp"`
	Iat               int         `json:"iat"`
	Jti               string      `json:"jti"`
	Iss               string      `json:"iss"`
	Aud               interface{} `json:"aud"`
	Sub               string      `json:"sub"`
	Typ               string      `json:"typ"`
	Azp               string      `json:"azp"`
	SessionState      string      `json:"session_state"`
	Name              string      `json:"name"`
	GivenName         string      `json:"given_name"`
	FamilyName        string      `json:"family_name"`
	PreferredUsername string      `json:"preferred_username"`
	Email             string      `json:"email"`
	EmailVerified     bool        `json:"email_verified"`
	Acr               string      `json:"acr"`
	RealmAccess       struct {
		Roles []string `json:"roles"`
	} `json:"realm_access"`
	ResourceAccess map[string]map[string][]string `json:"resource_access"`
	Scope          string                         `json:"scope"`
	Sid            string                         `json:"sid"`
	ClientId       string                         `json:"client_id"`
	Username       string                         `json:"username"`
	Active         bool                           `json:"active"`
}

type RefreshTokenInput struct {
	RefreshToken string `json:"refresh_token"`
}

type CreatePermissionInput struct {
	Title        *string  `json:"title" `
	Code         string   `json:"code" binding:"required"`
	Role         string   `json:"role" binding:"required"`
	ResourceCode string   `json:"resource_code" binding:"required"`
	ScopeCodes   []string `json:"scope_codes" binding:"required"`
}

type CheckPermissionInput struct {
	PermissionCode string `json:"permission_code" binding:"required"`
	Scope          string `json:"scope" binding:"required"`
}

type PermissionOut struct {
	Resource    string          `db:"resource_code" json:"resource_code"`
	Permissions json.RawMessage `db:"permissions" json:"permissions"`
}

type AutoGenerated struct {
	FirstName     string        `json:"firstName"`
	LastName      string        `json:"lastName"`
	Username      string        `json:"username"`
	Email         string        `json:"email"`
	EmailVerified string        `json:"emailVerified"`
	Enabled       string        `json:"enabled"`
	Credentials   []Credentials `json:"credentials"`
	Groups        []any         `json:"groups"`
	Attributes    Attributes    `json:"attributes"`
}
type Credentials struct {
	Type      string `json:"type"`
	Value     string `json:"value"`
	Temporary bool   `json:"temporary"`
}
type Attributes struct {
}

type UserCreateResponse struct {
	Status  int    `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	UserID  string `json:"user-id,omitempty"`
}

type RoleCrudResponse struct {
	Status  int    `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	RoleIds []Role `json:"role-id,omitempty"`
}

type Role struct {
	Id string `json:"id"`
}

type Roles struct {
	Id   *string `json:"id"`
	Name *string `json:"name"`
}

type ChangePassword struct {
	OldPassword *string `json:"old-password,omitempty"`
	NewPassword string  `json:"new-password"`
	IsTemporary *string `json:"is-temporary,omitempty"`
}

type JwtStruct struct {
	Claims struct {
		Acr               string `json:"acr"`
		Aud               string `json:"aud"`
		Azp               string `json:"azp"`
		Email             string `json:"email"`
		EmailVerified     bool   `json:"email_verified"`
		Exp               int    `json:"exp"`
		FamilyName        string `json:"family_name"`
		GivenName         string `json:"given_name"`
		Iat               int    `json:"iat"`
		Iss               string `json:"iss"`
		Jti               string `json:"jti"`
		Name              string `json:"name"`
		PreferredUsername string `json:"preferred_username"`
		RealmAccess       struct {
			Roles []string `json:"roles"`
		} `json:"realm_access"`
		ResourceAccess struct {
			RealmManagement struct {
				Roles []string `json:"roles"`
			} `json:"realm-management"`
		} `json:"resource_access"`
		Scope        string `json:"scope"`
		SessionState string `json:"session_state"`
		Sid          string `json:"sid"`
		Sub          string `json:"sub"`
		Typ          string `json:"typ"`
	} `json:"Claims"`
}

type Client struct {
	Id       *string `json:"id"`
	ClientId *string `json:"clientId"`
}
