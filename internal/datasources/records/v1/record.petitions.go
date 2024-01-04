package v1

import (
	V1Domains "change-it/internal/business/domains/v1"
	"time"
)

type Petitions struct {
	Id          string    `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	OwnerID     string    `db:"owner_id"`
	LikesCount  int64     `db:"likes_count"`
	VoicesCount int64     `db:"voices_count"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func FromPetitionsV1Domain(d *V1Domains.PetitionDomain) *Petitions {
	return &Petitions{
		Id:          d.ID,
		Title:       d.Title,
		Description: d.Description,
		OwnerID:     d.OwnerID,
		CreatedAt:   d.CreatedAt,
		UpdatedAt:   d.UpdatedAt,
	}
}

func (p *Petitions) ToV1Domain() *V1Domains.PetitionDomain {
	return &V1Domains.PetitionDomain{
		ID:          p.Id,
		Title:       p.Title,
		Description: p.Description,
		OwnerID:     p.OwnerID,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
		Likes:       p.LikesCount,
		Voices:      p.VoicesCount,
	}
}

func ToArrayOfPetitionsV1Domain(p *[]Petitions) []*V1Domains.PetitionDomain {
	var result []*V1Domains.PetitionDomain

	for _, val := range *p {
		result = append(result, val.ToV1Domain())
	}

	return result
}
