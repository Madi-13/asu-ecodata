package domain

type RoleDeleteRequest struct {
	Code string `json:"code"`
}

type RoleRequest struct {
	Code    string `json:"code"`
	Title   string `json:"title"`
	TitleKz string `json:"title_kk"`
}

type RoleResponse struct {
	RoleRequest
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}
