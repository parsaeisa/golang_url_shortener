package store

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type Cache struct {
	redisClient *redis.Client
}

var storeService = &Cache{}
var c = context.Background()

const TimeLimit = 24 * time.Hour

func ConnectToRedis() *Cache {

	// create redis connection
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:8000",
		Password: "",
		DB:       0,
	})

	//check redis connection
	_, err := redisClient.Ping(c).Result()
	if err != nil {
		fmt.Sprintf("Redis failed : %v", err)
	}

	fmt.Printf("Redis successed")
	storeService.redisClient = redisClient
	return storeService
}
