package v1

import (
	"errors"
	"mdm-auth/internal/domain"
	"net/http"
	"strings"

	"github.com/Nerzal/gocloak/v13"

	"github.com/gin-gonic/gin"
)

const (
	BEARER = "Bearer "
)

func (h *Handler) InitAuthRoutes(apiGroup *gin.RouterGroup) {
	authGroup := apiGroup.Group("/auth")
	{

		authGroup.POST("/login", h.LoginUser)
		authGroup.POST("/refresh", h.RefreshUserToken)
		authGroup.POST("/validate", h.ValidateToken)
		authGroup.POST("/logout", h.LogoutUser)
		authGroup.POST("/change-password", h.SetPassword)

		permissionGroup := authGroup.Group("/permission")
		{
			permissionGroup.GET("/", h.GetUserAccess)
			permissionGroup.POST("/", h.CreatePermission)
			permissionGroup.POST("/check", h.CheckPermission)
		}

	}
	userGroup := apiGroup.Group("/user")
	{
		userGroup.POST("/create", h.CreatreUser)
		userGroup.POST("/update", h.UpdateUser)
		userGroup.GET("/get-list", h.GetAllUsers)
		userGroup.GET("/get-user", h.GetUserInfo)
		userGroup.DELETE("/delete", h.DeleteUser)

	}
	roleGroup := apiGroup.Group("/role")
	{
		roleGroup.GET("/get-user-roles", h.GetAvailableRoles)
		roleGroup.POST("/add-user-roles", h.AddRolesToUser)
		roleGroup.DELETE("/delete-user-roles", h.DeleteClientRoleFromUser)

		roleGroup.GET("/get-list", h.GetRealmRoles)
		roleGroup.POST("/create", h.CreateRealmRole)
		roleGroup.DELETE("/delete", h.DeleteRealmRole)

		roleGroup.GET("/get-clients", h.GetClients)
		roleGroup.GET("/get-client-roles", h.GetClientRoles)
		roleGroup.PUT("/create-client-role", h.CreateClientRole)
		roleGroup.DELETE("/delete-client-role", h.DeleteClientRole)
	}
}

func (h *Handler) LoginUser(c *gin.Context) {
	username, password, ok := c.Request.BasicAuth()
	if !ok {
		UnauthorizedResp(c, "Basic authorization is required")
		return
	}
	authResponse, err := h.services.AuthService.Login(c.Request.Context(), username, password)
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, &authResponse)
}

func (h *Handler) RefreshUserToken(c *gin.Context) {
	var inp domain.RefreshTokenInput
	if err := c.ShouldBindJSON(&inp); err != nil {
		UnauthorizedResp(c, err.Error())
		return
	}
	refreshResponse, err := h.services.AuthService.Refresh(c.Request.Context(), inp.RefreshToken)
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, &refreshResponse)
}

func (h *Handler) ValidateToken(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		UnauthorizedResp(c, "Bearer accessToken is required")
		return
	}
	accessToken, errBear := ParseBearToken(auth)
	if errBear != nil {
		UnauthorizedResp(c, errBear.Error())
		return
	}
	validResponse, err := h.services.AuthService.ValidateToken(c.Request.Context(), accessToken)
	if err != nil {
		ForbiddenResp(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, validResponse)
}

func (h *Handler) CheckPermission(c *gin.Context) {
	var inp domain.CheckPermissionInput
	if err := c.ShouldBindJSON(&inp); err != nil {
		UnauthorizedResp(c, err.Error())
		return
	}

	auth := c.GetHeader("Authorization")
	if auth == "" {
		UnauthorizedResp(c, "Bearer accessToken is required")
		return
	}
	accessToken, errBear := ParseBearToken(auth)
	if errBear != nil {
		UnauthorizedResp(c, errBear.Error())
		return
	}
	isAcceptable, err := h.services.AuthService.CheckPermission(c.Request.Context(), accessToken, inp.PermissionCode, inp.Scope)
	if err != nil {
		ErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"has_permission": isAcceptable})
}

func (h *Handler) CreatePermission(c *gin.Context) {
	var inp domain.CreatePermissionInput
	if err := c.ShouldBindJSON(&inp); err != nil {
		UnauthorizedResp(c, err.Error())
		return
	}

	auth := c.GetHeader("Authorization")
	if auth == "" {
		UnauthorizedResp(c, "Bearer accessToken is required")
		return
	}
	accessToken, errBear := ParseBearToken(auth)
	if errBear != nil {
		UnauthorizedResp(c, errBear.Error())
		return
	}
	err := h.services.AuthService.CreatePermission(c.Request.Context(), accessToken,
		inp.Title, inp.Code, inp.Role, inp.ResourceCode, inp.ScopeCodes)
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Permission successfully saved",
	})
}

func (h *Handler) GetUserAccess(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		UnauthorizedResp(c, "Bearer accessToken is required")
		return
	}
	accessToken, errBear := ParseBearToken(auth)
	if errBear != nil {
		UnauthorizedResp(c, errBear.Error())
		return
	}
	newPermission, err := h.services.AuthService.GetPermissions(c.Request.Context(), accessToken)
	if err != nil {
		ErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}
	c.JSON(http.StatusOK, newPermission)
}

func (h *Handler) LogoutUser(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		UnauthorizedResp(c, "Bearer refreshToken is required")
		return
	}

	refreshToken, errBear := ParseBearToken(auth)
	if errBear != nil {
		UnauthorizedResp(c, errBear.Error())
		return
	}
	if err := h.services.AuthService.Logout(c.Request.Context(), refreshToken); err != nil {
		ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Logout successfully",
	})

}

func ParseBearToken(header string) (accessToken string, err error) {
	const prefix = "Bearer "
	if len(header) < len(prefix) || !strings.EqualFold(header[:len(prefix)], prefix) {
		return "", errors.New("Wrong auth length")
	}
	return header[len(prefix):], nil
}

func (h *Handler) CreatreUser(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		UnauthorizedResp(c, "Bearer accessToken is required")
		return
	}
	accessToken, errBear := ParseBearToken(auth)
	if errBear != nil {
		UnauthorizedResp(c, errBear.Error())
		return
	}
	var user gocloak.User

	if err := c.ShouldBindJSON(&user); err != nil {
		BadRequestResp(c, err.Error())
		return
	}

	userID, err := h.services.AuthService.CreateNewUser(c.Request.Context(), user, accessToken)
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	response := domain.UserCreateResponse{Status: http.StatusOK, Message: "User successfully created", UserID: userID}
	c.JSON(http.StatusOK, &response)

}

func (h *Handler) UpdateUser(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		UnauthorizedResp(c, "Bearer accessToken is required")
		return
	}
	accessToken, errBear := ParseBearToken(auth)
	if errBear != nil {
		UnauthorizedResp(c, errBear.Error())
		return
	}
	var user gocloak.User

	if err := c.ShouldBindJSON(&user); err != nil {
		BadRequestResp(c, err.Error())
		return
	}

	err := h.services.AuthService.UpdateUserInfo(c.Request.Context(), user, accessToken)
	if err != nil {
		ErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	response := domain.UserCreateResponse{Status: http.StatusOK, Message: "User successfully updated", UserID: *user.ID}
	c.JSON(http.StatusOK, &response)

}

func (h *Handler) GetAvailableRoles(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		UnauthorizedResp(c, "Bearer accessToken is required")
		return
	}
	accessToken, errBear := ParseBearToken(auth)
	if errBear != nil {
		UnauthorizedResp(c, errBear.Error())
		return
	}
	userId := c.GetHeader("user-id")
	if userId == "" {
		BadRequestResp(c, "user-id is required")
		return
	}

	roles, err := h.services.AuthService.GetUserRoles(c.Request.Context(), userId, accessToken)
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	var rolesList []domain.Roles
	var r domain.Roles
	for _, cMapping := range roles.ClientMappings {
		for _, role := range *cMapping.Mappings {
			r = domain.Roles{Id: role.ID, Name: role.Name}
			rolesList = append(rolesList, r)
		}

	}
	c.JSON(http.StatusOK, rolesList)

}

func (h *Handler) GetAllUsers(c *gin.Context) {
	params := gocloak.GetUsersParams{}
	auth := c.GetHeader("Authorization")
	if auth == "" {
		UnauthorizedResp(c, "Bearer accessToken is required")
		return
	}
	accessToken, errBear := ParseBearToken(auth)
	if errBear != nil {
		UnauthorizedResp(c, errBear.Error())
		return
	}
	userList, err := h.services.AuthService.GetUsersList(c.Request.Context(), params, accessToken)
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, &userList)

}

func (h *Handler) GetUserInfo(c *gin.Context) {

	auth := c.GetHeader("Authorization")
	if auth == "" {
		UnauthorizedResp(c, "Bearer accessToken is required")
		return
	}
	accessToken, errBear := ParseBearToken(auth)
	if errBear != nil {
		UnauthorizedResp(c, errBear.Error())
		return
	}

	userId := c.GetHeader("user-id")
	if userId == "" {
		BadRequestResp(c, "user-id is required")
		return
	}
	userInfo, err := h.services.AuthService.GetUser(c.Request.Context(), userId, accessToken)
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, &userInfo)

}

func (h *Handler) DeleteUser(c *gin.Context) {

	auth := c.GetHeader("Authorization")
	if auth == "" {
		UnauthorizedResp(c, "Bearer accessToken is required")
		return
	}
	accessToken, errBear := ParseBearToken(auth)
	if errBear != nil {
		UnauthorizedResp(c, errBear.Error())
		return
	}

	userId := c.GetHeader("user-id")
	if userId == "" {
		BadRequestResp(c, "user-id is required")
		return
	}
	err := h.services.AuthService.DeleteUserByID(c.Request.Context(), userId, accessToken)
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	response := domain.UserCreateResponse{Status: http.StatusOK, Message: "User successfully deleted", UserID: userId}
	c.JSON(http.StatusOK, response)

}

func (h *Handler) AddRolesToUser(c *gin.Context) {

	auth := c.GetHeader("Authorization")
	if auth == "" {
		UnauthorizedResp(c, domain.AccessTokenRequired)
		return
	}
	accessToken, errBear := ParseBearToken(auth)
	if errBear != nil {
		UnauthorizedResp(c, errBear.Error())
		return
	}

	userId := c.GetHeader("user-id")

	if userId == "" {
		BadRequestResp(c, domain.UserIdRequired)
		return
	}
	idOfClient := c.GetHeader("client-id")
	if idOfClient == "" {
		BadRequestResp(c, domain.ClientIdRequired)
		return
	}
	var roles []gocloak.Role

	if err := c.ShouldBindJSON(&roles); err != nil {
		BadRequestResp(c, err.Error())
		return
	}
	err := h.services.AuthService.AddClientRolesToUser(c.Request.Context(), idOfClient, userId, accessToken, roles)
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	response := domain.RoleCrudResponse{Status: http.StatusOK, Message: "Roles successfully added to user"}
	c.JSON(http.StatusOK, response)

}

func (h *Handler) DeleteRealmRole(c *gin.Context) {

	auth := c.GetHeader("Authorization")
	if auth == "" {
		UnauthorizedResp(c, "Bearer accessToken is required")
		return
	}
	accessToken, errBear := ParseBearToken(auth)
	if errBear != nil {
		UnauthorizedResp(c, errBear.Error())
		return
	}

	roleName := c.GetHeader("role-name")
	if roleName == "" {
		BadRequestResp(c, "role-name is required")
		return
	}

	err := h.services.AuthService.DeleteRealmRole(c.Request.Context(), roleName, accessToken)
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	response := domain.RoleCrudResponse{Status: http.StatusOK, Message: "Role successfully deleted"}
	c.JSON(http.StatusOK, response)

}

func (h *Handler) DeleteRealmRoleFromUser(c *gin.Context) {

	auth := c.GetHeader("Authorization")
	if auth == "" {
		UnauthorizedResp(c, "Bearer accessToken is required")
		return
	}
	accessToken, errBear := ParseBearToken(auth)
	if errBear != nil {
		UnauthorizedResp(c, errBear.Error())
		return
	}

	userId := c.GetHeader("user-id")
	if userId == "" {
		BadRequestResp(c, "user-id is required")
		return
	}

	var roles []gocloak.Role

	if err := c.ShouldBindJSON(&roles); err != nil {
		BadRequestResp(c, err.Error())
		return
	}
	err := h.services.AuthService.DeleteRealmRoleFromUser(c.Request.Context(), userId, accessToken, roles)
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	response := domain.RoleCrudResponse{Status: http.StatusOK, Message: "User roles successfully deleted"}
	c.JSON(http.StatusOK, response)

}

func (h *Handler) CreateRealmRole(c *gin.Context) {

	auth := c.GetHeader("Authorization")
	if auth == "" {
		UnauthorizedResp(c, "Bearer accessToken is required")
		return
	}
	accessToken, errBear := ParseBearToken(auth)
	if errBear != nil {
		UnauthorizedResp(c, errBear.Error())
		return
	}

	var role gocloak.Role

	if err := c.ShouldBindJSON(&role); err != nil {
		BadRequestResp(c, err.Error())
		return
	}
	_, err := h.services.AuthService.CreateRole(c.Request.Context(), accessToken, role)
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	response := domain.RoleCrudResponse{Status: http.StatusOK, Message: "Role successfully created"}
	c.JSON(http.StatusOK, response)

}

func (h *Handler) GetRealmRoles(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		UnauthorizedResp(c, "Bearer accessToken is required")
		return
	}
	accessToken, errBear := ParseBearToken(auth)
	if errBear != nil {
		UnauthorizedResp(c, errBear.Error())
		return
	}
	roles, err := h.services.AuthService.GetRealmAllRoles(c.Request.Context(), accessToken)
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	var rolesList []domain.Roles
	var r domain.Roles
	for _, role := range roles {
		if *role.Name != "default-roles-mdm" {
			r = domain.Roles{Id: role.ID, Name: role.Name}
			rolesList = append(rolesList, r)
		}
	}
	c.JSON(http.StatusOK, &rolesList)

}

func (h *Handler) SetPassword(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		UnauthorizedResp(c, "Bearer accessToken is required")
		return
	}
	accessToken, errBear := ParseBearToken(auth)
	if errBear != nil {
		UnauthorizedResp(c, errBear.Error())
		return
	}
	var changePassword domain.ChangePassword

	if err := c.ShouldBindJSON(&changePassword); err != nil {
		BadRequestResp(c, err.Error())
		return
	}

	err := h.services.AuthService.SetPassword(c.Request.Context(), accessToken, changePassword, false)
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	response := domain.UserCreateResponse{Status: http.StatusOK, Message: "Password successfully updated"}
	c.JSON(http.StatusOK, &response)

}

func (h *Handler) GetClientRoles(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		UnauthorizedResp(c, domain.AccessTokenRequired)
		return
	}
	accessToken, errBear := ParseBearToken(auth)
	if errBear != nil {
		UnauthorizedResp(c, errBear.Error())
		return
	}
	idOfClient := c.GetHeader("client-id")
	if idOfClient == "" {
		BadRequestResp(c, domain.ClientIdRequired)
		return
	}
	roles, err := h.services.AuthService.GetClientRoles(c.Request.Context(), idOfClient, accessToken)
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	var rolesList []domain.Roles
	var r domain.Roles
	for _, role := range roles {
		r = domain.Roles{Id: role.ID, Name: role.Name}
		rolesList = append(rolesList, r)
	}
	c.JSON(http.StatusOK, &rolesList)

}

func (h *Handler) GetClients(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		UnauthorizedResp(c, domain.AccessTokenRequired)
		return
	}
	accessToken, errBear := ParseBearToken(auth)
	if errBear != nil {
		UnauthorizedResp(c, errBear.Error())
		return
	}
	clients, err := h.services.AuthService.GetClients(c.Request.Context(), accessToken)
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	var clientsList []domain.Client
	var r domain.Client
	for _, client := range clients {
		r = domain.Client{Id: client.ID, ClientId: client.ClientID}
		clientsList = append(clientsList, r)
	}
	c.JSON(http.StatusOK, &clientsList)
}

func (h *Handler) CreateClientRole(c *gin.Context) {

	auth := c.GetHeader("Authorization")
	if auth == "" {
		UnauthorizedResp(c, domain.AccessTokenRequired)
		return
	}
	accessToken, errBear := ParseBearToken(auth)
	if errBear != nil {
		UnauthorizedResp(c, errBear.Error())
		return
	}
	idOfClient := c.GetHeader("client-id")
	if idOfClient == "" {
		BadRequestResp(c, domain.ClientIdRequired)
		return
	}

	var role gocloak.Role

	if err := c.ShouldBindJSON(&role); err != nil {
		BadRequestResp(c, err.Error())
		return
	}
	_, err := h.services.AuthService.CreateClientRole(c.Request.Context(), accessToken, idOfClient, role)
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	response := domain.RoleCrudResponse{Status: http.StatusOK, Message: "Role successfully created"}
	c.JSON(http.StatusOK, response)

}

func (h *Handler) DeleteClientRole(c *gin.Context) {

	auth := c.GetHeader("Authorization")
	if auth == "" {
		UnauthorizedResp(c, domain.AccessTokenRequired)
		return
	}
	accessToken, errBear := ParseBearToken(auth)
	if errBear != nil {
		UnauthorizedResp(c, errBear.Error())
		return
	}

	roleName := c.GetHeader("role-name")
	if roleName == "" {
		BadRequestResp(c, domain.RoleNameRequired)
		return
	}

	idOfClient := c.GetHeader("client-id")
	if idOfClient == "" {
		BadRequestResp(c, domain.ClientIdRequired)
		return
	}

	err := h.services.AuthService.DeleteClientRole(c.Request.Context(), accessToken, idOfClient, roleName)
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	response := domain.RoleCrudResponse{Status: http.StatusOK, Message: "Role successfully deleted"}
	c.JSON(http.StatusOK, response)

}

func (h *Handler) DeleteClientRoleFromUser(c *gin.Context) {

	auth := c.GetHeader("Authorization")
	if auth == "" {
		UnauthorizedResp(c, domain.AccessTokenRequired)
		return
	}
	accessToken, errBear := ParseBearToken(auth)
	if errBear != nil {
		UnauthorizedResp(c, errBear.Error())
		return
	}

	userId := c.GetHeader("user-id")
	if userId == "" {
		BadRequestResp(c, domain.UserIdRequired)
	}

	idOfClient := c.GetHeader("client-id")
	if idOfClient == "" {
		BadRequestResp(c, domain.ClientIdRequired)
		return
	}

	var roles []gocloak.Role

	if err := c.ShouldBindJSON(&roles); err != nil {
		BadRequestResp(c, err.Error())
		return
	}
	err := h.services.AuthService.DeleteClientRolesFromUser(c.Request.Context(), userId, idOfClient, accessToken, roles)
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	response := domain.RoleCrudResponse{Status: http.StatusOK, Message: "User roles successfully deleted"}
	c.JSON(http.StatusOK, response)

}
