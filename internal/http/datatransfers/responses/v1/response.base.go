package v1

type PageResponse struct {
	Data        interface{} `json:"data"`
	CurrentPage int64       `json:"current_page"`
	PageSize    int64       `json:"page_size"`
	TotalPages  int64       `json:"total_pages"`
}
