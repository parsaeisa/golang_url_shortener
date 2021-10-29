package store

import (
	"github.com/go-redis/redis"
	"github.com/go-redis/redis/v8"
)

type Cache struct {
	redisClient *redis.Client
}
