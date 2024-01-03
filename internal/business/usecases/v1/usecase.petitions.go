package v1

import (
	V1Domains "change-it/internal/business/domains/v1"
	"context"
	"github.com/google/uuid"
)

type petitionUsecase struct {
	petitionRepository V1Domains.PetitionRepository
}

func (p *petitionUsecase) Create(ctx context.Context, domain *V1Domains.PetitionDomain) (err error) {

	domain.OwnerID = uuid.New().String()

	err = p.petitionRepository.Create(ctx, domain)
	if err != nil {
		return err
	}
	return nil
}

func (p *petitionUsecase) Delete(ctx context.Context, id string, userId string) (err error) {
	//TODO implement me
	panic("implement me")
}

func (p *petitionUsecase) Like(ctx context.Context, id string, userId string) (err error) {
	//TODO implement me
	panic("implement me")
}

func (p *petitionUsecase) Voice(ctx context.Context, id string, userId string) (err error) {
	//TODO implement me
	panic("implement me")
}

func (p *petitionUsecase) GetAll(ctx context.Context) ([]V1Domains.PetitionDomain, error) {
	//TODO implement me
	panic("implement me")
}

func NewPetitionUsecase(petitionRepository V1Domains.PetitionRepository) V1Domains.PetitionUseсase {
	return &petitionUsecase{
		petitionRepository: petitionRepository,
	}
}
