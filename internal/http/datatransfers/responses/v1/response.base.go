package v1

type PageResponse struct {
	Data        interface{} `json:"data"`
	CurrentPage int         `json:"current_page"`
	PageSize    int         `json:"page_size"`
	TotalPages  int         `json:"total_pages"`
}
