package v1

type Id struct {
	ID string `uri:"id" json:"id" query:"id" validate:"required,uuid4"`
}
