package bootstrap

import (
	"github.com/veaquer/go_backend_template/internal/cache"
	"github.com/veaquer/go_backend_template/internal/config"
	"context"

	"github.com/redis/go-redis/v9"
)

func ProvideRedis(cfg *config.Config) *cache.RedisCache {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		panic("Failed to connect to Redis: " + err.Error())
	}

	return cache.NewRedisCache(rdb)
}
