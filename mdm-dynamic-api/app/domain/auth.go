package domain

type Auth struct {
	Acr               string      `json:"acr"`
	Aud               interface{} `json:"aud"`
	Azp               string      `json:"azp"`
	EmailVerified     bool        `json:"email_verified"`
	Exp               int         `json:"exp"`
	FamilyName        string      `json:"family_name"`
	GivenName         string      `json:"given_name"`
	Iat               int         `json:"iat"`
	Iss               string      `json:"iss"`
	Jti               string      `json:"jti"`
	Name              string      `json:"name"`
	PreferredUsername string      `json:"preferred_username"`
	RealmAccess       struct {
		Roles []string `json:"roles"`
	} `json:"realm_access"`
	ResourceAccess struct {
		Account struct {
			Roles []string `json:"roles"`
		} `json:"account"`
		MdmCli struct {
			Roles []string `json:"roles"`
		} `json:"mdm-cli"`
	} `json:"resource_access"`
	Scope        string `json:"scope"`
	SessionState string `json:"session_state"`
	Sid          string `json:"sid"`
	Sub          string `json:"sub"`
	Typ          string `json:"typ"`
}

type Permission struct {
	PermissionCode string `json:"permission_code"`
	Scope          string `json:"scope"`
}

type CreatePermission struct {
	Code         string   `json:"code"`
	Title        string   `json:"title"`
	Role         string   `json:"role"`
	ResourceCode string   `json:"resource_code"`
	ScopeCodes   []string `json:"scope_codes"`
}

type TokenStruct struct {
	AccessToken      string `json:"access_token"`
	IDToken          string `json:"id_token"`
	ExpiresIn        int    `json:"expires_in"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
	RefreshToken     string `json:"refresh_token"`
	TokenType        string `json:"token_type"`
	NotBeforePolicy  int    `json:"not-before-policy"`
	SessionState     string `json:"session_state"`
	Scope            string `json:"scope"`
}

type Users struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
type DMLUsers struct {
	OperationType string `json:"operation_type"`
	Users         Users  `json:"user"`
}
