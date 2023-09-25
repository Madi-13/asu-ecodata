package services

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"mdm-auth/internal/domain"
	"mdm-auth/internal/providers"
	"mdm-auth/internal/repositories"
	"mdm-auth/internal/utils/logging"
	"strings"

	"github.com/Nerzal/gocloak/v13"
)

type AuthService interface {
	Login(ctx context.Context, username, password string) (*gocloak.JWT, error)
	Refresh(ctx context.Context, refreshToken string) (*gocloak.JWT, error)
	ValidateToken(ctx context.Context, accessToken string) (interface{}, error)
	CreatePermission(ctx context.Context, accessToken string, title *string, code, role, resourceCode string, scopeCodes []string) error
	GetPermissions(ctx context.Context, roleName string) (interface{}, error)
	CheckPermission(ctx context.Context, accessToken, permissionCode, scope string) (*bool, error)
	Logout(ctx context.Context, refreshToken string) error
	CreateNewUser(ctx context.Context, user gocloak.User, token string) (string, error)
	UpdateUserInfo(ctx context.Context, user gocloak.User, token string) error
	GetUserRoles(ctx context.Context, userId, accessToken string) (*gocloak.MappingsRepresentation, error)
	GetUsersList(ctx context.Context, params gocloak.GetUsersParams, accessToken string) ([]*gocloak.User, error)
	GetUser(ctx context.Context, userId string, accessToken string) (*gocloak.User, error)
	DeleteUserByID(ctx context.Context, userId string, accessToken string) error
	AddRoleToUser(ctx context.Context, userId string, accessToken string, roles []gocloak.Role) error
	DeleteRealmRole(ctx context.Context, roleName, accessToken string) error
	DeleteRealmRoleFromUser(ctx context.Context, userId string, accessToken string, roles []gocloak.Role) error
	CreateRole(ctx context.Context, accessToken string, role gocloak.Role) (string, error)
	GetRealmAllRoles(ctx context.Context, accessToken string) ([]*gocloak.Role, error)
	SetPassword(ctx context.Context, token string, changePassword domain.ChangePassword, temporary bool) error
	GetClients(ctx context.Context, accessToken string) ([]*gocloak.Client, error)
	GetClientRoles(ctx context.Context, idOfClient, token string) ([]*gocloak.Role, error)
	CreateClientRole(ctx context.Context, token, idOfClient string, role gocloak.Role) (string, error)
	DeleteClientRole(ctx context.Context, token, idOfClient string, roleName string) error
	DeleteClientRolesFromUser(ctx context.Context, idOfClient, userId, token string, roles []gocloak.Role) error
	AddClientRolesToUser(ctx context.Context, idOfClient, userId, accessToken string, roles []gocloak.Role) error
}

type AuthServiceDeps struct {
	KeycloakProvider providers.KeycloakProvider
	AuthRepository   repositories.Auth
}

func (deps *AuthServiceDeps) AddClientRolesToUser(ctx context.Context, idOfClient, userId, token string, roles []gocloak.Role) error {
	err := deps.KeycloakProvider.AddClientRolesToUser(ctx, idOfClient, userId, token, roles)
	if err != nil {
		logging.Error(domain.InternalError, err.Error())
		return err
	}
	return nil
}

func NewAuthService(keycloakProvider providers.KeycloakProvider, authRepo repositories.Auth) *AuthServiceDeps {
	return &AuthServiceDeps{
		KeycloakProvider: keycloakProvider,
		AuthRepository:   authRepo,
	}
}

func (deps *AuthServiceDeps) Login(ctx context.Context, username, password string) (*gocloak.JWT, error) {
	resp, err := deps.KeycloakProvider.Authorize(ctx, username, password)
	if err != nil {
		logging.Error("Could not get tokens", err.Error())
		return nil, err
	}
	return resp, err
}

func (deps *AuthServiceDeps) Refresh(ctx context.Context, refreshToken string) (*gocloak.JWT, error) {
	resp, err := deps.KeycloakProvider.RefreshToken(ctx, refreshToken)
	if err != nil {
		logging.Error("Could not get tokens", err.Error())
		return nil, err
	}
	return resp, err
}

func (deps *AuthServiceDeps) ValidateToken(ctx context.Context, accessToken string) (interface{}, error) {
	resp, err := deps.KeycloakProvider.ValidAuthToken(ctx, accessToken)
	if err != nil {
		logging.Error("Could not get tokens", err.Error())
		return nil, err
	}
	return resp, err
}

func (deps *AuthServiceDeps) CreatePermission(ctx context.Context, accessToken string, title *string, code,
	role, resourceCode string, scopeCodes []string) error {
	err := deps.AuthRepository.CreatePermissions(ctx, title, code, role, resourceCode, scopeCodes)
	if err != nil {
		logging.Error("Could not create new permission", err.Error())
		return err
	}
	return nil
}

func (deps *AuthServiceDeps) CheckPermission(ctx context.Context, accessToken, permissionCode, scope string) (*bool, error) {
	_, err := deps.KeycloakProvider.ValidAuthToken(ctx, accessToken)
	if err != nil {
		logging.Error("Could not get tokens", err.Error())
		return nil, err
	}
	token, err := decodeFromBase64(accessToken)
	if err != nil {
		logging.Error("Could marshall token", err.Error())
		return nil, err
	}
	roles := getRolesFromToken(token.ResourceAccess)
	if len(roles) == 0 {
		return nil, nil
	}
	hasPermission, err := deps.AuthRepository.CheckPermission(ctx, roles, permissionCode, scope)
	if err != nil {
		logging.Error("Could not get tokens", err.Error())
		return nil, err
	}
	return hasPermission, nil
}

func (deps *AuthServiceDeps) GetPermissions(ctx context.Context, accessToken string) (interface{}, error) {
	_, err := deps.KeycloakProvider.ValidAuthToken(ctx, accessToken)
	if err != nil {
		logging.Error("Could not get tokens", err.Error())
		return nil, err
	}
	token, err := decodeFromBase64(accessToken)
	if err != nil {
		logging.Error("Could marshall token", err.Error())
		return nil, err
	}
	roles := getRolesFromToken(token.ResourceAccess)
	if len(roles) == 0 {
		return nil, nil
	}
	permissions, err := deps.AuthRepository.GetPermissions(ctx, roles)
	if err != nil {
		logging.Error("Could not get tokens", err.Error())
		return nil, err
	}
	return permissions, err
}

func (deps *AuthServiceDeps) Logout(ctx context.Context, refreshToken string) error {
	err := deps.KeycloakProvider.Logout(ctx, refreshToken)
	if err != nil {
		logging.Error("Could not get tokens", err.Error())
		return err
	}
	return nil
}

func decodeFromBase64(base64String string) (*domain.ValidResponse, error) {
	var tokenPayload domain.ValidResponse
	tokenBase64Payload := strings.Split(base64String, ".")[1]
	strBytes, err := base64.RawStdEncoding.DecodeString(tokenBase64Payload)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(strBytes, &tokenPayload); err != nil {
		return nil, err
	}
	return &tokenPayload, nil
}

func getRolesFromToken(resourceAccess map[string]map[string][]string) []string {
	clientMap, ok := resourceAccess["mdm-cli"]
	if !ok {
		return nil
	}
	roles, ok := clientMap["roles"]
	if !ok {
		return nil
	}
	return roles
}

func (deps *AuthServiceDeps) CreateNewUser(ctx context.Context, user gocloak.User, token string) (string, error) {
	userID, err := deps.KeycloakProvider.CreateNewUser(ctx, user, token)
	if err != nil {
		logging.Error("Could not get tokens", err.Error())
		return "", err
	}
	return userID, err
}

func (deps *AuthServiceDeps) UpdateUserInfo(ctx context.Context, user gocloak.User, token string) error {
	err := deps.KeycloakProvider.UpdateUser(ctx, user, token)
	if err != nil {
		logging.Error("Could not get tokens", err.Error())
		return err
	}
	return err
}

func (deps *AuthServiceDeps) GetUser(ctx context.Context, userId string, token string) (*gocloak.User, error) {
	user, err := deps.KeycloakProvider.GetUserByID(ctx, userId, token)
	if err != nil {
		logging.Error("Could not get tokens", err.Error())
		return nil, err
	}
	return user, err
}

func (deps *AuthServiceDeps) DeleteUserByID(ctx context.Context, userId string, token string) error {
	err := deps.KeycloakProvider.DeleteUser(ctx, userId, token)
	if err != nil {
		logging.Error("Could not get tokens", err.Error())
		return err
	}
	return err
}

func (deps *AuthServiceDeps) AddRoleToUser(ctx context.Context, userId string, token string, roles []gocloak.Role) error {
	err := deps.KeycloakProvider.AddRealmRoleToUser(ctx, userId, token, roles)
	if err != nil {
		logging.Error("Could not get tokens", err.Error())
		return err
	}
	return err
}

func (deps *AuthServiceDeps) DeleteRealmRole(ctx context.Context, roleName, token string) error {
	err := deps.KeycloakProvider.DeleteRealmRole(ctx, token, roleName)
	if err != nil {
		logging.Error("Could not get tokens", err.Error())
		return err
	}
	return err
}

func (deps *AuthServiceDeps) DeleteRealmRoleFromUser(ctx context.Context, userId string, token string, roles []gocloak.Role) error {
	err := deps.KeycloakProvider.DeleteRealmRoleFromUser(ctx, userId, token, roles)
	if err != nil {
		logging.Error("Could not get tokens", err.Error())
		return err
	}
	return err
}

func (deps *AuthServiceDeps) CreateRole(ctx context.Context, token string, role gocloak.Role) (string, error) {
	roleId, err := deps.KeycloakProvider.CreateRealmRole(ctx, token, role)
	if err != nil {
		logging.Error("Could not get tokens", err.Error())
		return roleId, err
	}
	return roleId, err
}

func (deps *AuthServiceDeps) GetUsersList(ctx context.Context, params gocloak.GetUsersParams, token string) ([]*gocloak.User, error) {
	usersList, err := deps.KeycloakProvider.GetUsers(ctx, params, token)
	if err != nil {
		logging.Error("Could not get tokens", err.Error())
		return nil, err
	}
	return usersList, err
}

func (deps *AuthServiceDeps) GetUserRoles(ctx context.Context, userID, accessToken string) (*gocloak.MappingsRepresentation, error) {
	roles, err := deps.KeycloakProvider.GetAvailableRoles(ctx, userID, accessToken)
	if err != nil {
		logging.Error("Could not get tokens", err.Error())
		return nil, err
	}
	return roles, err
}

func (deps *AuthServiceDeps) GetRealmAllRoles(ctx context.Context, accessToken string) ([]*gocloak.Role, error) {
	roles, err := deps.KeycloakProvider.GetRealmAllRoles(ctx, accessToken)
	if err != nil {
		logging.Error("Could not get tokens", err.Error())
		return nil, err
	}
	return roles, err
}

func (deps *AuthServiceDeps) SetPassword(ctx context.Context, token string, changePassword domain.ChangePassword, temporary bool) error {
	err := deps.KeycloakProvider.SetPassword(ctx, token, changePassword, temporary)
	if err != nil {
		logging.Error("Could not get tokens", err.Error())
		return err
	}
	return nil
}

func (deps *AuthServiceDeps) GetClients(ctx context.Context, accessToken string) ([]*gocloak.Client, error) {
	roles, err := deps.KeycloakProvider.GetClients(ctx, accessToken)
	if err != nil {
		logging.Error(domain.InternalError, err.Error())
		return nil, err
	}
	return roles, nil
}

func (deps *AuthServiceDeps) GetClientRoles(ctx context.Context, idOfClient, token string) ([]*gocloak.Role, error) {
	clientRoles, err := deps.KeycloakProvider.GetClientRoles(ctx, idOfClient, token)
	if err != nil {
		logging.Error(domain.InternalError, err.Error())
		return nil, err
	}
	return clientRoles, nil
}

func (deps *AuthServiceDeps) CreateClientRole(ctx context.Context, token, idOfClient string, role gocloak.Role) (string, error) {
	roleId, err := deps.KeycloakProvider.CreateClientRole(ctx, token, idOfClient, role)
	if err != nil {
		logging.Error(domain.InternalError, err.Error())
		return roleId, err
	}
	return roleId, nil
}

func (deps *AuthServiceDeps) DeleteClientRolesFromUser(ctx context.Context, userId, idOfClient, token string, roles []gocloak.Role) error {
	err := deps.KeycloakProvider.DeleteClientRolesFromUser(ctx, idOfClient, userId, token, roles)
	if err != nil {
		logging.Error(domain.InternalError, err.Error())
		return err
	}
	return nil
}

func (deps *AuthServiceDeps) DeleteClientRole(ctx context.Context, token, idOfClient string, roleName string) error {
	err := deps.KeycloakProvider.DeleteClientRole(ctx, token, idOfClient, roleName)
	if err != nil {
		logging.Error(domain.InternalError, err.Error())
		return err
	}
	return nil
}
