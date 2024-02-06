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

func (p *postgreUserRepository) GetOwnedPetitions(ctx context.Context, userId string, pageNumber int, pageSize int) (outDomains []*V1Domains.PetitionDomain, total int, err error) {

	offset := (pageNumber - 1) * pageSize

	query := `
        SELECT petitions.*, 
               (SELECT COUNT(*) FROM likes WHERE likes.petition_id = petitions.id) AS likes_count,
               (SELECT COUNT(*) FROM voices WHERE voices.petition_id = petitions.id) AS voices_count
        FROM petitions
        WHERE petitions.owner_id = $1
        GROUP BY petitions.id 
        ORDER BY MAX(created_at) DESC
        OFFSET $2
        LIMIT $3
    `

	var outRecords []V1Records.Petitions
	err = p.conn.SelectContext(ctx, &outRecords, query, userId, offset, pageSize)

	query = `SELECT COUNT(*) FROM petitions WHERE owner_id = $1`

	var totalRecords int
	err = p.conn.GetContext(ctx, &totalRecords, query, userId)

	if totalRecords%pageSize == 0 {
		total = totalRecords / pageSize
	} else {
		total = totalRecords/pageSize + 1
	}

	outDomains = V1Records.ToArrayOfPetitionsV1Domain(&outRecords)
	if err != nil {
		return outDomains, 0, err
	}

	return outDomains, total, nil
}

func (p *postgreUserRepository) GetLikedPetitions(ctx context.Context, userId string, pageNumber int, pageSize int) (outDomains []*V1Domains.PetitionDomain, total int, err error) {

	offset := (pageNumber - 1) * pageSize

	query := `
        SELECT petitions.*, 
               (SELECT COUNT(*) FROM likes WHERE likes.petition_id = petitions.id) AS likes_count,
               (SELECT COUNT(*) FROM voices WHERE voices.petition_id = petitions.id) AS voices_count
        FROM petitions
        JOIN likes ON petitions.id = likes.petition_id
        WHERE likes.user_id = $1
        GROUP BY petitions.id 
        ORDER BY MAX(created_at) DESC
        OFFSET $2
        LIMIT $3
    `

	var outRecords []V1Records.Petitions
	err = p.conn.SelectContext(ctx, &outRecords, query, userId, offset, pageSize)

	query = `        
		SELECT COUNT(petitions.*)
        FROM petitions
        JOIN likes ON petitions.id = likes.petition_id
        WHERE likes.user_id = $1
        `

	var totalRecords int
	err = p.conn.GetContext(ctx, &totalRecords, query, userId)

	if totalRecords%pageSize == 0 {
		total = totalRecords / pageSize
	} else {
		total = totalRecords/pageSize + 1
	}

	outDomains = V1Records.ToArrayOfPetitionsV1Domain(&outRecords)
	if err != nil {
		return outDomains, 0, err
	}

	return outDomains, total, nil
}

func (p *postgreUserRepository) GetVoicedPetitions(ctx context.Context, userId string, pageNumber int, pageSize int) (outDomains []*V1Domains.PetitionDomain, total int, err error) {

	offset := (pageNumber - 1) * pageSize

	query := `
        SELECT petitions.*, 
               (SELECT COUNT(*) FROM likes WHERE likes.petition_id = petitions.id) AS likes_count,
               (SELECT COUNT(*) FROM voices WHERE voices.petition_id = petitions.id) AS voices_count
        FROM petitions
        JOIN voices ON petitions.id = voices.petition_id
        WHERE voices.user_id = $1
        GROUP BY petitions.id 
        ORDER BY MAX(created_at) DESC
        OFFSET $2
        LIMIT $3
    `

	var outRecords []V1Records.Petitions
	err = p.conn.SelectContext(ctx, &outRecords, query, userId, offset, pageSize)

	query = `SELECT COUNT(petitions)
        FROM petitions
        JOIN voices ON petitions.id = voices.petition_id
        WHERE voices.user_id = $1`

	var totalRecords int
	err = p.conn.GetContext(ctx, &totalRecords, query, userId)

	if totalRecords%pageSize == 0 {
		total = totalRecords / pageSize
	} else {
		total = totalRecords/pageSize + 1
	}

	outDomains = V1Records.ToArrayOfPetitionsV1Domain(&outRecords)
	if err != nil {
		return outDomains, 0, err
	}

	return outDomains, total, nil
}

func (p *postgreUserRepository) IsUserLikedPetition(ctx context.Context, userId string, petitionId string) (res bool, err error) {
	query := `SELECT COUNT(*) FROM likes WHERE user_id = $1 AND petition_id = $2`

	var result int64
	err = p.conn.GetContext(ctx, &result, query, userId, petitionId)
	if err != nil {
		return false, err
	}

	return result > 0, nil

}

func (p *postgreUserRepository) IsUserVoicedPetition(ctx context.Context, userId string, petitionId string) (res bool, err error) {
	query := `SELECT COUNT(*) FROM voices WHERE user_id = $1 AND petition_id = $2`

	var result int64
	err = p.conn.GetContext(ctx, &result, query, userId, petitionId)
	if err != nil {
		return false, err
	}

	return result > 0, nil

}

func NewUserRepository(conn *sqlx.DB) V1Domains.UserRepository {
	return &postgreUserRepository{
		conn: conn,
	}
}
