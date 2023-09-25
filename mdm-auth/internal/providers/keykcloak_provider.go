package providers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mdm-auth/internal/domain"

	"github.com/Nerzal/gocloak/v13"
	"github.com/golang-jwt/jwt/v4"
)

type KeycloakProvider interface {
	Authorize(ctx context.Context, username, password string) (*gocloak.JWT, error)
	RefreshToken(ctx context.Context, refreshToken string) (*gocloak.JWT, error)
	ValidAuthToken(ctx context.Context, accessToken string) (*jwt.MapClaims, error)
	Logout(ctx context.Context, refreshToken string) error
	CreateNewUser(ctx context.Context, user gocloak.User, token string) (string, error)
	UpdateUser(ctx context.Context, user gocloak.User, token string) error
	GetUserByID(ctx context.Context, userId string, token string) (*gocloak.User, error)
	DeleteUser(ctx context.Context, userId string, token string) error
	AddRealmRoleToUser(ctx context.Context, userId string, token string, roles []gocloak.Role) error
	DeleteRealmRole(ctx context.Context, token, roleName string) error
	DeleteRealmRoleFromUser(ctx context.Context, userId string, token string, roles []gocloak.Role) error
	CreateRealmRole(ctx context.Context, token string, role gocloak.Role) (string, error)
	GetUsers(ctx context.Context, params gocloak.GetUsersParams, token string) ([]*gocloak.User, error)
	GetAvailableRoles(ctx context.Context, userID, accessToken string) (*gocloak.MappingsRepresentation, error)
	GetRealmAllRoles(ctx context.Context, accessToken string) ([]*gocloak.Role, error)
	SetPassword(ctx context.Context, token string, changePassword domain.ChangePassword, temporary bool) error
	GetClients(ctx context.Context, token string) ([]*gocloak.Client, error)
	GetClientRoles(ctx context.Context, idOfClient, token string) ([]*gocloak.Role, error)
	CreateClientRole(ctx context.Context, token, idOfClient string, role gocloak.Role) (string, error)
	DeleteClientRole(ctx context.Context, token, idOfClient string, roleName string) error
	DeleteClientRolesFromUser(ctx context.Context, idOfClient, userId, token string, roles []gocloak.Role) error
	AddClientRolesToUser(ctx context.Context, idOfClient, userId, token string, roles []gocloak.Role) error
}

type KeycloakDeps struct {
	Host, ClientId, ClientSecret, Realm string
	KeycloakClient                      *gocloak.GoCloak
	PublicKey                           string
}

func NewKeycloakDeps(host, realm, clientId, clientSecret, publicKey string) *KeycloakDeps {
	return &KeycloakDeps{
		Host:           host,
		Realm:          realm,
		ClientId:       clientId,
		ClientSecret:   clientSecret,
		KeycloakClient: gocloak.NewClient(host),
		PublicKey:      "-----BEGIN CERTIFICATE-----\n" + publicKey + "\n-----END CERTIFICATE-----",
	}
}

func (deps *KeycloakDeps) AddClientRolesToUser(ctx context.Context, idOfClient, userId, token string, roles []gocloak.Role) error {
	client := deps.KeycloakClient
	err := client.AddClientRolesToUser(ctx, token, deps.Realm, idOfClient, userId, roles)
	if err != nil {
		return err
	}
	return nil
}

func (deps *KeycloakDeps) Authorize(ctx context.Context, username, password string) (*gocloak.JWT, error) {
	jwt, err := deps.KeycloakClient.Login(ctx, deps.ClientId, deps.ClientSecret, deps.Realm, username, password)
	if err != nil {
		return nil, err
	}
	return jwt, nil
}

func (deps *KeycloakDeps) RefreshToken(ctx context.Context, refreshToken string) (*gocloak.JWT, error) {
	jwt, err := deps.KeycloakClient.RefreshToken(ctx, refreshToken, deps.ClientId, deps.ClientSecret, deps.Realm)
	if err != nil {
		return nil, err
	}
	return jwt, nil
}

func (deps *KeycloakDeps) ValidAuthToken(ctx context.Context, accessToken string) (*jwt.MapClaims, error) {
	key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(deps.PublicKey))
	if err != nil {
		return nil, fmt.Errorf("validate: parse key: %w", err)
	}

	tok, err := jwt.Parse(accessToken, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
		}
		return key, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := tok.Claims.(jwt.MapClaims)

	if !ok || !tok.Valid {
		return nil, fmt.Errorf("validate: invalid")
	}

	return &claims, nil
}

func (deps *KeycloakDeps) Logout(ctx context.Context, refreshToken string) error {
	err := deps.KeycloakClient.Logout(ctx, deps.ClientId, deps.ClientSecret, deps.Realm, refreshToken)
	if err != nil {
		return err
	}
	return nil
}

func (deps *KeycloakDeps) CreateNewUser(ctx context.Context, user gocloak.User, token string) (string, error) {
	client := deps.KeycloakClient
	userID, err := client.CreateUser(ctx, token, deps.Realm, user)
	if err != nil {
		println("Oh no!, failed to create user :(")
		return "", err
	}
	return userID, nil
}

func (deps *KeycloakDeps) UpdateUser(ctx context.Context, user gocloak.User, token string) error {
	client := deps.KeycloakClient
	err := client.UpdateUser(ctx, token, deps.Realm, user)
	if err != nil {
		println("Oh no!, failed to update user :(")
		return err
	}
	return nil
}

func (deps *KeycloakDeps) GetUserByID(ctx context.Context, userId string, token string) (*gocloak.User, error) {
	client := deps.KeycloakClient
	user, err := client.GetUserByID(ctx, token, deps.Realm, userId)
	if err != nil {
		println("Oh no!, failed to get user :(")
		return nil, err
	}
	return user, nil
}

func (deps *KeycloakDeps) AddRealmRoleToUser(ctx context.Context, userId string, token string, roles []gocloak.Role) error {
	client := deps.KeycloakClient
	err := client.AddRealmRoleToUser(ctx, token, deps.Realm, userId, roles)
	if err != nil {
		println("Oh no!, failed to get user :(")
		return err
	}
	return nil
}

func (deps *KeycloakDeps) DeleteRealmRole(ctx context.Context, token, roleName string) error {
	client := deps.KeycloakClient
	err := client.DeleteRealmRole(ctx, token, deps.Realm, roleName)
	if err != nil {
		println("Oh no!, failed to delete role :(")
		return err
	}
	return nil
}

func (deps *KeycloakDeps) DeleteRealmRoleFromUser(ctx context.Context, userId string, token string, roles []gocloak.Role) error {
	client := deps.KeycloakClient
	err := client.DeleteRealmRoleFromUser(ctx, token, deps.Realm, userId, roles)
	if err != nil {
		println("Oh no!, failed to get user :(")
		return err
	}
	return nil
}

func (deps *KeycloakDeps) CreateRealmRole(ctx context.Context, token string, role gocloak.Role) (string, error) {
	client := deps.KeycloakClient
	roleId, err := client.CreateRealmRole(ctx, token, deps.Realm, role)
	if err != nil {
		println("Oh no!, failed to get user :(")
		return "", err
	}
	return roleId, nil
}

func (deps *KeycloakDeps) DeleteUser(ctx context.Context, userId string, token string) error {
	client := deps.KeycloakClient
	err := client.DeleteUser(ctx, token, deps.Realm, userId)
	if err != nil {
		println("Oh no!, failed to delete user :(")
		return err
	}
	return nil
}

func (deps *KeycloakDeps) GetUsers(ctx context.Context, params gocloak.GetUsersParams, token string) ([]*gocloak.User, error) {
	client := deps.KeycloakClient
	users, err := client.GetUsers(ctx, token, deps.Realm, params)
	if err != nil {
		println("Oh no!, failed to get users list :(")
		return nil, err
	}
	return users, nil
}

func (deps *KeycloakDeps) GetRealmAllRoles(ctx context.Context, token string) ([]*gocloak.Role, error) {
	params := gocloak.GetRoleParams{}
	rolesList, err := deps.KeycloakClient.GetRealmRoles(ctx, token, deps.Realm, params)
	if err != nil {
		println("Oh no!, failed to get realm roles :(")
		return nil, err
	}
	return rolesList, nil
}

func (deps *KeycloakDeps) GetAvailableRoles(ctx context.Context, userID, accessToken string) (*gocloak.MappingsRepresentation, error) {
	client := deps.KeycloakClient
	id, err := client.GetRoleMappingByUserID(ctx, accessToken, deps.Realm, userID)
	if err != nil {
		println("Oh no!, failed to get user roles :(")
		return nil, err
	}
	return id, nil
}

func (deps *KeycloakDeps) SetPassword(ctx context.Context, token string, changePassword domain.ChangePassword, temporary bool) error {
	client := deps.KeycloakClient
	decodedJwt := parseJwtToken(token)
	_, err := client.Login(ctx, deps.ClientId, deps.ClientSecret, deps.Realm, decodedJwt.Claims.PreferredUsername, *changePassword.OldPassword)
	if err != nil {
		return err
	}

	err = client.SetPassword(ctx, token, decodedJwt.Claims.Sub, deps.Realm, changePassword.NewPassword, temporary)
	if err != nil {
		println("Oh no!, failed to get user :(")
		return err
	}
	return nil
}

func parseJwtToken(token string) domain.JwtStruct {
	tkn, _, err := jwt.NewParser().ParseUnverified(token, jwt.MapClaims{})
	if err != nil {
		log.Fatal(err)
	}

	byteTkn, err := json.Marshal(tkn)
	if err != nil {
		log.Fatal(err)
	}
	var decodedJwt domain.JwtStruct
	json.Unmarshal(byteTkn, &decodedJwt)

	return decodedJwt
}

func (deps *KeycloakDeps) GetClients(ctx context.Context, token string) ([]*gocloak.Client, error) {
	params := gocloak.GetClientsParams{}
	clients, err := deps.KeycloakClient.GetClients(ctx, token, deps.Realm, params)
	if err != nil {
		return nil, err
	}
	return clients, nil
}

func (deps *KeycloakDeps) GetClientRoles(ctx context.Context, idOfClient, token string) ([]*gocloak.Role, error) {
	params := gocloak.GetRoleParams{BriefRepresentation: ToBoolPointer(false)}
	rolesList, err := deps.KeycloakClient.GetClientRoles(ctx, token, deps.Realm, idOfClient, params)
	if err != nil {
		return nil, err
	}
	return rolesList, nil
}

func (deps *KeycloakDeps) CreateClientRole(ctx context.Context, token, idOfClient string, role gocloak.Role) (string, error) {
	client := deps.KeycloakClient
	roleId, err := client.CreateClientRole(ctx, token, deps.Realm, idOfClient, role)
	if err != nil {
		return "", err
	}
	return roleId, nil
}

func ToBoolPointer(v bool) *bool {
	return &v
}

func (deps *KeycloakDeps) DeleteClientRolesFromUser(ctx context.Context, idOfClient, userId, token string, roles []gocloak.Role) error {
	client := deps.KeycloakClient
	err := client.DeleteClientRolesFromUser(ctx, token, deps.Realm, idOfClient, userId, roles)
	if err != nil {
		return err
	}
	return nil
}

func (deps *KeycloakDeps) DeleteClientRole(ctx context.Context, token, idOfClient string, roleName string) error {
	client := deps.KeycloakClient
	err := client.DeleteClientRole(ctx, token, deps.Realm, idOfClient, roleName)
	if err != nil {
		return err
	}
	return nil
}
