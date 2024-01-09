package v1

import (
	"context"
	"time"
)

type PetitionDomain struct {
	ID          string
	Title       string
	Description string
	OwnerID     string
	Likes       int64
	Voices      int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type PetitionUseсase interface {
	Save(ctx context.Context, domain *PetitionDomain) (err error)
	Delete(ctx context.Context, id string, userId string, userRoles []string) (err error)
	Like(ctx context.Context, id string, userId string) (err error)
	Voice(ctx context.Context, id string, userId string) (err error)
	GetAll(ctx context.Context) ([]*PetitionDomain, error)
}

type PetitionRepository interface {
	Create(ctx context.Context, domain *PetitionDomain) (err error)
	Update(ctx context.Context, domain *PetitionDomain) (err error)
	Delete(ctx context.Context, id string) (err error)
	Like(ctx context.Context, id string, userId string) (err error)
	Voice(ctx context.Context, id string, userId string) (err error)
	GetAll(ctx context.Context) (outDomains []*PetitionDomain, err error)
	GetByID(ctx context.Context, id string) (outDomain *PetitionDomain, err error)
}
