package v1

import (
	V1Domains "change-it/internal/business/domains/v1"
	V1Records "change-it/internal/datasources/records/v1"
	"context"
	"github.com/jmoiron/sqlx"
)

type postgreUserRepository struct {
	conn *sqlx.DB
}

func (p postgreUserRepository) GetLikedPetitions(ctx context.Context, userId string) (outDomains []*V1Domains.PetitionDomain, err error) {

	query := `
        SELECT petitions.*, 
               (SELECT COUNT(*) FROM likes WHERE likes.petition_id = petitions.id) AS likes_count,
               (SELECT COUNT(*) FROM voices WHERE voices.petition_id = petitions.id) AS voices_count
        FROM petitions
        JOIN likes ON petitions.id = likes.petition_id
        WHERE likes.user_id = $1
    `

	var outRecords []V1Records.Petitions
	err = p.conn.SelectContext(ctx, &outRecords, query, userId)

	outDomains = V1Records.ToArrayOfPetitionsV1Domain(&outRecords)
	if err != nil {
		return outDomains, err
	}

	return outDomains, nil
}

func (p postgreUserRepository) GetVoicedPetitions(ctx context.Context, userId string) (outDomains []*V1Domains.PetitionDomain, err error) {
	query := `
        SELECT petitions.*, 
               (SELECT COUNT(*) FROM likes WHERE likes.petition_id = petitions.id) AS likes_count,
               (SELECT COUNT(*) FROM voices WHERE voices.petition_id = petitions.id) AS voices_count
        FROM petitions
        JOIN voices ON petitions.id = voices.petition_id
        WHERE voices.user_id = $1
    `

	var outRecords []V1Records.Petitions
	err = p.conn.SelectContext(ctx, &outRecords, query, userId)

	outDomains = V1Records.ToArrayOfPetitionsV1Domain(&outRecords)
	if err != nil {
		return outDomains, err
	}

	return outDomains, nil
}

func NewUserRepository(conn *sqlx.DB) V1Domains.UserRepository {
	return &postgreUserRepository{
		conn: conn,
	}
}
