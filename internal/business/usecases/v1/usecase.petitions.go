package v1

import (
	V1Domains "change-it/internal/business/domains/v1"
	V1DomainErrors "change-it/internal/business/errors/v1"
	"context"
	"errors"
	"github.com/google/uuid"
)

type petitionUsecase struct {
	petitionRepository V1Domains.PetitionRepository
}

func (p *petitionUsecase) Save(ctx context.Context, domain *V1Domains.PetitionDomain) (err error) {
	_, err = p.petitionRepository.GetByID(ctx, domain.ID)

	domain.OwnerID = uuid.New().String()
	if err != nil {
		return p.petitionRepository.Create(ctx, domain)
	}

	return p.petitionRepository.Update(ctx, domain)
}

func (p *petitionUsecase) Delete(ctx context.Context, id string, userId string) (err error) {

	domain, err := p.petitionRepository.GetByID(ctx, id)

	// switch case style
	if err != nil {
		var nferr *V1DomainErrors.NotFoundError
		switch {
		case errors.As(err, &nferr):
			return &V1DomainErrors.NotFoundError{
				Message:    "petition not found",
				StatusCode: 409,
			}
		default:
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

	return p.petitionRepository.Like(ctx, id, uuid.New().String())
}

func (p *petitionUsecase) Voice(ctx context.Context, id string, userId string) (err error) {
	_, err = p.petitionRepository.GetByID(ctx, id)

	userId = uuid.New().String()
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

	return p.petitionRepository.Voice(ctx, id, userId)
}

func (p *petitionUsecase) GetAll(ctx context.Context) ([]*V1Domains.PetitionDomain, error) {
	return p.petitionRepository.GetAll(ctx)
}

func NewPetitionUsecase(petitionRepository V1Domains.PetitionRepository) V1Domains.PetitionUseсase {
	return &petitionUsecase{
		petitionRepository: petitionRepository,
	}
}
