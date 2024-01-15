package v1

type Id struct {
	ID string `uri:"id" json:"id" query:"id" validate:"required,uuid4"`
}

type PageRequest struct {
	PageNumber int64 `query:"page"`
	PageSize   int64 `query:"page-size"`
}
