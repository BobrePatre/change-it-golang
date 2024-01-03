package v1

import (
	V1Domains "change-it/internal/business/domains/v1"
	"time"
)

type PetitionResponse struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	OwnerID     string    `json:"owner_id"`
	Likes       int64     `json:"likes"`
	Voices      int64     `json:"voices"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func FromV1Domain(domain *V1Domains.PetitionDomain) *PetitionResponse {
	return &PetitionResponse{
		Id:          domain.ID,
		Title:       domain.Title,
		Description: domain.Description,
		OwnerID:     domain.OwnerID,
		Likes:       domain.Likes,
		Voices:      domain.Voices,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}

func ArrayFromV1Domains(domains []V1Domains.PetitionDomain) []*PetitionResponse {
	var out []*PetitionResponse
	for _, domain := range domains {
		out = append(out, FromV1Domain(&domain))
	}
	return out
}
