package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type MemoryCache struct {
	Host string `json:"host" yaml:"host"`

	// Data info.
	Username string `json:"user" yaml:"user"`
	Password string `json:"password" yaml:"password"`
}

type RedisHelper struct {
	Memory redisResources
}

type redisResources interface {
	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
}
