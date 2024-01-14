package v1

import (
	V1Domains "change-it/internal/business/domains/v1"
	V1DomainErrors "change-it/internal/business/errors/v1"
	V1DatasourceErrors "change-it/internal/datasources/errors/v1"
	"change-it/pkg/helpers"
	"change-it/pkg/logger"
	"context"
	"errors"
)

type petitionUsecase struct {
	petitionRepository V1Domains.PetitionRepository
	userRepository     V1Domains.UserRepository
}

func (p *petitionUsecase) Create(ctx context.Context, domain *V1Domains.PetitionDomain) (err error) {
	return p.petitionRepository.Create(ctx, domain)
}

func (p *petitionUsecase) Update(ctx context.Context, domain *V1Domains.PetitionDomain) (err error) {
	_, err = p.petitionRepository.GetByID(ctx, domain.ID)

	if err != nil {
		return &V1DomainErrors.NotFoundError{
			Message:    "petition not found",
			StatusCode: 409,
		}
	}

	return p.petitionRepository.Update(ctx, domain)
}

func (p *petitionUsecase) Delete(ctx context.Context, id string, userId string, userRoles []string) (err error) {

	domain, err := p.petitionRepository.GetByID(ctx, id)

	// switch case style
	if err != nil {
		var nferr *V1DatasourceErrors.NotFoundError
		switch {
		case errors.As(err, &nferr):
			logger.Info("petition not found error", nil)
			return &V1DomainErrors.NotFoundError{
				Message:    "petition not found",
				StatusCode: 409,
			}
		default:
			logger.Info("unknown error", nil)
			return err
		}
	}

	if userId == domain.OwnerID {
		return p.petitionRepository.Delete(ctx, id)
	}

	if helpers.IsArrayContains(userRoles, "admin") {
		return p.petitionRepository.Delete(ctx, id)
	}

	return &V1DomainErrors.ForbiddenError{
		Message:    "User cannot delete this petition",
		StatusCode: 403,
	}
}

func (p *petitionUsecase) Like(ctx context.Context, id string, userId string) (err error) {

	_, err = p.petitionRepository.GetByID(ctx, id)

	// if style
	if err != nil {
		var nferr *V1DatasourceErrors.NotFoundError
		if errors.As(err, &nferr) {
			return &V1DomainErrors.NotFoundError{
				Message:    "petition not found",
				StatusCode: 409,
			}
		}
		return err
	}

	isUserLiked, err := p.userRepository.IsUserLikedPetition(ctx, userId, id)
	if err != nil {
		return err
	}

	if isUserLiked {
		return &V1DomainErrors.AlreadyLikedError{
			Message:    "petition already liked",
			StatusCode: 409,
		}
	}

	return p.petitionRepository.Like(ctx, id, userId)
}

func (p *petitionUsecase) Voice(ctx context.Context, id string, userId string) (err error) {
	_, err = p.petitionRepository.GetByID(ctx, id)

	if err != nil {
		var nferr *V1DatasourceErrors.NotFoundError
		if errors.As(err, &nferr) {
			return &V1DomainErrors.NotFoundError{
				Message:    "petition not found",
				StatusCode: 409,
			}
		}
		return err
	}

	isUserVoiced, err := p.userRepository.IsUserVoicedPetition(ctx, userId, id)
	if err != nil {
		return err
	}

	if isUserVoiced {
		return &V1DomainErrors.AlreadyVoicedError{
			Message:    "petition already voiced",
			StatusCode: 409,
		}
	}

	return p.petitionRepository.Voice(ctx, id, userId)
}

func (p *petitionUsecase) GetAll(ctx context.Context) ([]*V1Domains.PetitionDomain, error) {
	return p.petitionRepository.GetAll(ctx)
}

func NewPetitionUsecase(petitionRepository V1Domains.PetitionRepository, userRepository V1Domains.UserRepository) V1Domains.PetitionUseсase {
	return &petitionUsecase{
		petitionRepository: petitionRepository,
		userRepository:     userRepository,
	}
}
