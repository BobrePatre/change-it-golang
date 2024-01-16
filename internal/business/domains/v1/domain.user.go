package v1

import "context"

type UserDetails struct {
	Roles      []string
	UserId     string
	Email      string
	Username   string
	Name       string
	FamilyName string
}

type UserUsecase interface {
	GetLikedPetitions(ctx context.Context, userId string, pageNumber int, pageSize int) (outPetitionDomains []*PetitionDomain, total int, err error)
	GetVoicedPetitions(ctx context.Context, userId string, pageNumber int, pageSize int) (outPetitionDomains []*PetitionDomain, total int, err error)
	GetOwnedPetitions(ctx context.Context, userId string, pageNumber int, pageSize int) (outPetitionDomains []*PetitionDomain, total int, err error)
}

type UserRepository interface {
	GetLikedPetitions(ctx context.Context, userId string, pageNumber int, pageSize int) (outDomains []*PetitionDomain, total int, err error)
	GetVoicedPetitions(ctx context.Context, userId string, pageNumber int, pageSize int) (outDomains []*PetitionDomain, total int, err error)
	GetOwnedPetitions(ctx context.Context, userId string, pageNumber int, pageSize int) (outDomains []*PetitionDomain, total int, err error)
	IsUserLikedPetition(ctx context.Context, userId string, petitionId string) (res bool, err error)
	IsUserVoicedPetition(ctx context.Context, userId string, petitionId string) (res bool, err error)
}
