package drivers

import (
	"github.com/redis/go-redis/v9"
)

type RedisConfig struct {
	Address  string
	Password string
	DB       int
}

func (config *RedisConfig) InitializeRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Address,
		Password: config.Password,
		DB:       config.DB,
	})
	return client
}
