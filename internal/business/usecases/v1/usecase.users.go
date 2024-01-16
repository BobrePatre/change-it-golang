package v1

import (
	V1Domains "change-it/internal/business/domains/v1"
	"context"
)

type userUsecase struct {
	userRepository V1Domains.UserRepository
}

func (u *userUsecase) GetLikedPetitions(ctx context.Context, userId string, pageNumber int, pageSize int) (outPetitionDomains []*V1Domains.PetitionDomain, total int, err error) {
	return u.userRepository.GetLikedPetitions(ctx, userId, pageNumber, pageSize)
}

func (u *userUsecase) GetVoicedPetitions(ctx context.Context, userId string, pageNumber int, pageSize int) (outPetitionDomains []*V1Domains.PetitionDomain, total int, err error) {
	return u.userRepository.GetVoicedPetitions(ctx, userId, pageNumber, pageSize)
}

func (u *userUsecase) GetOwnedPetitions(ctx context.Context, userId string, pageNumber int, pageSize int) (outPetitionDomains []*V1Domains.PetitionDomain, total int, err error) {
	return u.userRepository.GetOwnedPetitions(ctx, userId, pageNumber, pageSize)
}

func NewUserUsecase(userRepository V1Domains.UserRepository) V1Domains.UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}
