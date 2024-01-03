package v1

import (
	V1Domains "change-it/internal/business/domains/v1"
	"time"
)

type Petitions struct {
	Id          string    `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	Likes       int64     `db:"likes"`
	Voices      int64     `db:"voices"`
	OwnerID     string    `db:"owner_id"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func FromPetitionsV1Domain(d *V1Domains.PetitionDomain) Petitions {
	return Petitions{
		Id:          d.ID,
		Title:       d.Title,
		Description: d.Description,
		OwnerID:     d.OwnerID,
		Likes:       d.Likes,
		Voices:      d.Voices,
		CreatedAt:   d.CreatedAt,
		UpdatedAt:   d.UpdatedAt,
	}
}

func (p *Petitions) ToV1Domain() V1Domains.PetitionDomain {
	return V1Domains.PetitionDomain{
		ID:          p.Id,
		Title:       p.Title,
		Description: p.Description,
		OwnerID:     p.OwnerID,
		Likes:       p.Likes,
		Voices:      p.Voices,
	}
}

func ToArrayOfPetitionsV1Domain(p *[]Petitions) []V1Domains.PetitionDomain {
	var result []V1Domains.PetitionDomain

	for _, val := range *p {
		result = append(result, val.ToV1Domain())
	}

	return result
}
