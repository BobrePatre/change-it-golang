package v1

import (
	V1Domains "change-it/internal/business/domains/v1"
	V1DatasourceErrors "change-it/internal/datasources/errors/v1"
	V1Records "change-it/internal/datasources/records/v1"
	"context"
	"github.com/jmoiron/sqlx"
)

type postgrePetitonRepository struct {
	conn *sqlx.DB
}

func (p *postgrePetitonRepository) Create(ctx context.Context, domain *V1Domains.PetitionDomain) (err error) {
	_, err = p.conn.NamedQueryContext(ctx, "INSERT INTO petitions (id, title, description, owner_id, created_at, updated_at) VALUES (uuid_generate_v4(), :title, :description, :owner_id,  current_timestamp, current_timestamp)", V1Records.FromPetitionsV1Domain(domain))
	if err != nil {
		return err
	}
	return nil
}

func (p *postgrePetitonRepository) Update(ctx context.Context, domain *V1Domains.PetitionDomain) (err error) {

	_, err = p.conn.NamedQueryContext(ctx, "UPDATE petitions SET title = :title, description = :description, updated_at = current_timestamp WHERE id = :id", V1Records.FromPetitionsV1Domain(domain))

	if err != nil {
		return err
	}
	return nil
}

func (p *postgrePetitonRepository) Delete(ctx context.Context, id string) (err error) {

	_, err = p.conn.QueryContext(ctx, "DELETE FROM petitions WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

func (p *postgrePetitonRepository) Like(ctx context.Context, id string, userId string) (err error) {

	like := V1Records.Likes{
		UserId:     userId,
		PetitionId: id,
	}

	_, err = p.conn.NamedQueryContext(ctx, "INSERT INTO likes (user_id, petition_id) VALUES (:user_id, :petition_id)", like)
	if err != nil {
		return err
	}
	return nil
}

func (p *postgrePetitonRepository) Voice(ctx context.Context, id string, userId string) (err error) {

	_, err = p.conn.QueryContext(ctx, "INSERT INTO voices (user_id, petition_id) VALUES ($1, $2)", userId, id)
	if err != nil {
		return err
	}
	return nil
}

func (p *postgrePetitonRepository) GetByID(ctx context.Context, id string) (outDomain *V1Domains.PetitionDomain, err error) {
	var petitionRecord V1Records.Petitions

	query := `
       SELECT petitions.*, 
              (SELECT COUNT(*) FROM likes where petition_id = petitions.id) AS likes_count, 
              (SELECT COUNT(*) FROM voices where petition_id = petitions.id) AS voices_count
       FROM petitions 
       WHERE petitions.id = $1 GROUP BY petitions.id
   `
	err = p.conn.GetContext(ctx, &petitionRecord, query, id)
	if err != nil {
		return nil, &V1DatasourceErrors.NotFoundError{Message: err.Error()}
	}

	outDomain = petitionRecord.ToV1Domain()
	return outDomain, nil
}
func (p *postgrePetitonRepository) GetAll(ctx context.Context, pageNumber int, pageSize int) (outDomains []*V1Domains.PetitionDomain, total int, err error) {

	offset := (pageNumber - 1) * pageSize

	query := `
        SELECT 
            petitions.*, 
            (SELECT COUNT(*) FROM likes WHERE petition_id = petitions.id) AS likes_count, 
            (SELECT COUNT(*) FROM voices WHERE petition_id = petitions.id) AS voices_count
        FROM petitions
        GROUP BY petitions.id
        ORDER BY MAX(created_at) DESC
        OFFSET $1
        LIMIT $2
        
    `

	var outRecords []V1Records.Petitions
	err = p.conn.SelectContext(ctx, &outRecords, query, offset, pageSize)

	query = `SELECT COUNT(*) FROM petitions`

	var totalRecords int
	err = p.conn.GetContext(ctx, &totalRecords, query)

	if totalRecords%pageSize == 0 {
		total = totalRecords / pageSize
	} else {
		total = totalRecords/pageSize + 1
	}

	outDomains = V1Records.ToArrayOfPetitionsV1Domain(&outRecords)
	if err != nil {
		return outDomains, total, err
	}

	return outDomains, total, nil
}

func NewPetitionRepository(conn *sqlx.DB) V1Domains.PetitionRepository {
	return &postgrePetitonRepository{
		conn: conn,
	}
}
