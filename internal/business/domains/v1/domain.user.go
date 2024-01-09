package v1

import "context"

type UserDetails struct {
	Roles    []string
	UserId   string
	Email    string
	Username string
}

type UserRepository interface {
	GetLikedPetitions(ctx context.Context, userId string) (outDomains []*PetitionDomain, err error)
	GetVoicedPetitions(ctx context.Context, userId string) (outDomains []*PetitionDomain, err error)
	IsUserLikedPetition(ctx context.Context, userId string, petitionId string) (res bool, err error)
	IsUserVoicedPetition(ctx context.Context, userId string, petitionId string) (res bool, err error)
}
