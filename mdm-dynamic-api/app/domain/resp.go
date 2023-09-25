package domain

type ResponsibilityRequest struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type ResponsibilityDeleteRequest struct {
	Code string `json:"code"`
}

type ResponseWithPagination struct {
	TotalPage   int         `json:"total_page"`
	TotalCount  int         `json:"total_count"`
	CurrentPage int         `json:"current_page"`
	Limit       int         `json:"limit"`
	Data        interface{} `json:"data"`
}
