package v1

import V1Domains "change-it/internal/business/domains/v1"

type PetitionRequest struct {
	Title       string `json:"title" validate:"required,min=5"`
	Description string `json:"description"`
}

func (r *PetitionRequest) ToV1Domain() *V1Domains.PetitionDomain {
	return &V1Domains.PetitionDomain{
		Title:       r.Title,
		Description: r.Description,
	}
}
