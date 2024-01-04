package v1

import (
	V1Domains "change-it/internal/business/domains/v1"
	V1DatasourceErrors "change-it/internal/datasources/errors/v1"
	V1Records "change-it/internal/datasources/records/v1"
	"change-it/pkg/logger"
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
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

	_, err = p.conn.NamedQueryContext(ctx, "DELETE FROM petitions WHERE id = :id", id)
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
		logger.Error(err.Error(), logrus.Fields{"userId": userId, "petitionId": id})
		return err
	}
	return nil
}

func (p *postgrePetitonRepository) Voice(ctx context.Context, id string, userId string) (err error) {

	voice := V1Records.Voices{
		UserId:     userId,
		PetitionId: id,
	}

	_, err = p.conn.NamedQueryContext(ctx, "INSERT INTO voices (user_id, petition_id) VALUES (:user_id, :petition_id)", voice)
	if err != nil {
		return err
	}
	return nil
}

func (p *postgrePetitonRepository) GetByID(ctx context.Context, id string) (outDomain *V1Domains.PetitionDomain, err error) {
	var petitionRecord V1Records.Petitions

	id = "4d8efe0b-e592-494e-bd4f-43d36d9944ee"
	query := `
       SELECT p.*, COUNT(l.*) AS likes_count, COUNT(v.*) AS voices_count
       FROM petitions AS p
       LEFT JOIN likes AS l ON p.id = l.petition_id
       LEFT JOIN voices AS v ON p.id = v.petition_id
       WHERE p.id = '4d8efe0b-e592-494e-bd4f-43d36d9944ee' GROUP BY p.id
   `
	logger.Info("Executing query with id", logrus.Fields{"id": id})
	err = p.conn.GetContext(ctx, &petitionRecord, query)
	if err != nil {
		logger.Error("Failed to execute query", logrus.Fields{"error": err.Error()})
		return nil, &V1DatasourceErrors.NotFoundError{Message: err.Error()}
	}

	outDomain = petitionRecord.ToV1Domain()
	return outDomain, nil
}
func (p *postgrePetitonRepository) GetAll(ctx context.Context) (outDomains []*V1Domains.PetitionDomain, err error) {
	query := `
        SELECT 
            petitions.*,
            (SELECT COUNT(*) FROM likes WHERE petition_id = petitions.id) AS likes_count,
            (SELECT COUNT(*) FROM voices WHERE petition_id = petitions.id) AS voices_count
        FROM petitions
        OFFSET 0
        LIMIT 5
    `

	var outRecords []V1Records.Petitions
	err = p.conn.SelectContext(ctx, &outRecords, query)
	if err != nil {
		return nil, err
	}

	return V1Records.ToArrayOfPetitionsV1Domain(&outRecords), nil
}

func NewPetitionRepository(conn *sqlx.DB) V1Domains.PetitionRepository {
	return &postgrePetitonRepository{
		conn: conn,
	}
}
