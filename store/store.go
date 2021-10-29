package store

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type redisCLI struct {
	redisClient *redis.Client
}

var storeService = &redisCLI{}
var c = context.Background()

func ConnectToRedis() *redisCLI {

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

const TimeLimit = 24 * time.Hour

// for increasing the pace of program
// we want to store encoded urls and
// later retrieve the decoded url
func AddEncodedURL(short, original, userId string) {
	err := storeService.redisClient.Set(c, short, original, TimeLimit)

	if err != nil {
		fmt.Sprintf("Failed while saving , the error : %v", err)
	}

}

func GetDecodedURL(short string) string {
	result, err := storeService.redisClient.Get(c, short).Result()
	if err != nil {
		fmt.Sprintf("Failed while saving , the error : %v", err)
		return err.Error()
	}

	return result
}
