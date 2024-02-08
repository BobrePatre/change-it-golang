package v1

import (
	"change-it/internal/constants"
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type postgreAuthRepository struct {
	conn *redis.Client
}

func (p *postgreAuthRepository) SaveJwk(ctx context.Context, jwk string) {
	p.conn.Set(ctx, constants.JwkKey, jwk, 120*time.Minute)
}

func (p *postgreAuthRepository) GetJwk(ctx context.Context) (jwk string, err error) {
	jwk, err = p.conn.Get(ctx, constants.JwkKey).Result()
	if err != nil {
		return "", err
	}
	return jwk, nil
}
