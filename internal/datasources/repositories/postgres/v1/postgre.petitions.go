package v1

import (
	V1Domains "change-it/internal/business/domains/v1"
	V1Records "change-it/internal/datasources/records/v1"
	"context"
	"github.com/jmoiron/sqlx"
)

type postgrePetitonRepository struct {
	conn *sqlx.DB
}

func NewPetitionRepository(conn *sqlx.DB) V1Domains.PetitionRepository {
	return &postgrePetitonRepository{
		conn: conn,
	}
}

func (p *postgrePetitonRepository) Create(ctx context.Context, domain *V1Domains.PetitionDomain) (err error) {
	_, err = p.conn.NamedQueryContext(ctx, "INSERT INTO petitions (id, title, description, owner_id, likes, voices, created_at, updated_at) VALUES (uuid_generate_v4(), :title, :description, :owner_id, :likes, :voices, current_timestamp, current_timestamp)", V1Records.FromPetitionsV1Domain(domain))
	if err != nil {
		return err
	}
	return nil
}

func (p *postgrePetitonRepository) Delete(ctx context.Context, id string) (err error) {

	_, err = p.conn.QueryContext(ctx, "DELETE FROM petitions WHERE id = :id", id)
	if err != nil {
		return err
	}

	return nil
}

func (p *postgrePetitonRepository) Like(ctx context.Context, id string, userId string) (err error) {
	tx, err := p.conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	like := V1Records.Likes{
		UserId:     userId,
		PetitionId: id,
	}

	_, err = tx.QueryContext(ctx, "INSERT INTO likes (user_id, petition_id) VALUES (:user_id, :petition_id)", like)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	_, err = tx.QueryContext(ctx, "UPDATE petitions SET likes = likes + 1 WHERE id = :id", id)

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	return nil
}

func (p *postgrePetitonRepository) Voice(ctx context.Context, id string, userId string) (err error) {
	tx, err := p.conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	like := V1Records.Likes{
		UserId:     userId,
		PetitionId: id,
	}

	_, err = tx.QueryContext(ctx, "INSERT INTO voices (user_id, petition_id) VALUES (:user_id, :petition_id)", like)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	_, err = tx.QueryContext(ctx, "UPDATE petitions SET voices = voices + 1 WHERE id = :id", id)

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	return nil
}

func (p *postgrePetitonRepository) GetAll(ctx context.Context) (outDomains []V1Domains.PetitionDomain, err error) {
	var petitionsRecords []V1Records.Petitions
	err = p.conn.SelectContext(ctx, &petitionsRecords, "SELECT * FROM petitions OFFSET 0 LIMIT 5")
	if err != nil {
		return nil, err
	}
	return V1Records.ToArrayOfPetitionsV1Domain(&petitionsRecords), nil
}
