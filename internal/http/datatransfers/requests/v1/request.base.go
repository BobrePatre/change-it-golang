package v1

type Id struct {
	ID string `uri:"id" json:"id" query:"id" validate:"required,uuid4"`
}

type PageRequest struct {
	PageNumber int `query:"page"`
	PageSize   int `query:"page-size"`
}
