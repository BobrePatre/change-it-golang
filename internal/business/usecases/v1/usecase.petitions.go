package v1

import (
	V1Domains "change-it/internal/business/domains/v1"
	V1DomainErrors "change-it/internal/business/errors/v1"
	V1DatasourceErrors "change-it/internal/datasources/errors/v1"
	"context"
	"errors"
	"github.com/snykk/go-rest-boilerplate/pkg/logger"
)

type petitionUsecase struct {
	petitionRepository V1Domains.PetitionRepository
	userRepository     V1Domains.UserRepository
}

func (p *petitionUsecase) Save(ctx context.Context, domain *V1Domains.PetitionDomain) (err error) {
	_, err = p.petitionRepository.GetByID(ctx, domain.ID)

	if err != nil {
		return p.petitionRepository.Create(ctx, domain)
	}

	return p.petitionRepository.Update(ctx, domain)
}

func (p *petitionUsecase) Delete(ctx context.Context, id string, userId string) (err error) {

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

	if userId != domain.OwnerID {
		return p.petitionRepository.Delete(ctx, id)
	}

	return p.petitionRepository.Delete(ctx, id)
}

func (p *petitionUsecase) Like(ctx context.Context, id string, userId string) (err error) {
	_, err = p.petitionRepository.GetByID(ctx, id)

	// if style
	if err != nil {
		var nferr *V1DomainErrors.NotFoundError
		if errors.As(err, &nferr) {
			return &V1DomainErrors.NotFoundError{
				Message:    "petition not found",
				StatusCode: 409,
			}
		}
		return err
	}

	likedPetitions, err := p.userRepository.GetLikedPetitions(ctx, userId)
	if err != nil {
		return err
	}

	for _, petition := range likedPetitions {
		if petition.ID == id {
			return &V1DomainErrors.AlreadyLikedError{
				Message:    "petition already liked",
				StatusCode: 409,
			}
		}
	}

	return p.petitionRepository.Like(ctx, id, userId)
}

func (p *petitionUsecase) Voice(ctx context.Context, id string, userId string) (err error) {
	_, err = p.petitionRepository.GetByID(ctx, id)

	if err != nil {
		var nferr *V1DomainErrors.NotFoundError
		if errors.As(err, &nferr) {
			return &V1DomainErrors.NotFoundError{
				Message:    "petition not found",
				StatusCode: 409,
			}
		}
		return err
	}

	voicedPetitions, err := p.userRepository.GetVoicedPetitions(ctx, userId)
	if err != nil {
		return err
	}

	for _, petition := range voicedPetitions {
		if petition.ID == id {
			return &V1DomainErrors.AlreadyVoicedError{
				Message:    "petition already voiced",
				StatusCode: 409,
			}
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
