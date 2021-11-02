package store

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

type redisCLI struct {
	redisClient *redis.Client
}

var storeService = &redisCLI{}

func ConnectToRedis() *redisCLI {

	// create redis connection
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	//check redis connection
	_, err := redisClient.Ping().Result()
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

// we store shortened urls and its original
// in redis caching system . the records will
// be deleted after 1 day .
func AddEncodedURL(short_url, original_url, userId string) {
	err := storeService.redisClient.Set(short_url, original_url, TimeLimit).Err()

	if err != nil {
		fmt.Sprintf("Failed while saving , the error : %v", err)
	}

}

// retrive original url from redis .
func GetDecodedURL(short_url string) string {
	result, err := storeService.redisClient.Get(short_url).Result()
	if err != nil {
		fmt.Sprintf("Failed while retrieving , the error : %v", err)
		return "Failed while retrieving , the error :" + err.Error()
	}

	return result
}
