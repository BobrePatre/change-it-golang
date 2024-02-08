package utils

import (
	"change-it/internal/config"
	"change-it/internal/datasources/drivers"
	"github.com/redis/go-redis/v9"
)

func SetupRedisConn() (rclient *redis.Client) {
	rconfig := drivers.RedisConfig{
		Address:  config.AppConfig.RedisAddress,
		Password: config.AppConfig.RedisPassword,
		DB:       config.AppConfig.RedisDB,
	}

	conn := rconfig.InitializeRedis()
	return conn
}
