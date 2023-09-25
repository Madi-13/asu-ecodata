package domain

const (
	LocaleRu       = `ru`
	LocaleKz       = `kz`
	UsernameHeader = "Username-Login"
)

type UserRequest struct {
	Username string `json:"username"`
	RoleCode string `json:"role_code"`
	Limit    int    `json:"limit"`
	Page     int    `json:"page"`
}

type UserResponse struct {
	Username  string  `json:"username"`
	RoleTitle *string `json:"role_title"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Email     string  `json:"email"`
	IIN       string  `json:"iin"`
	BIN       string  `json:"bin"`
	Mobile    string  `json:"mobile"`
}

type UserDeleteRequest struct {
	Username string `json:"username"`
}
