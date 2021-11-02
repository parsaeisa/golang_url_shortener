package store

import (
	"log"
	"testing"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
)

func TestRedisConnection(t *testing.T) {

	c := ConnectToRedis()

	assert.True(t, c.redisClient != nil)
}

func TestStoreToRedis(t *testing.T) {
	original := "https://club.snapp.ir/radio-snapp/newseason/"
	short := "http://localhost:8080/s8yQsd"

	// make connections to mocked redis
	mr, err := miniredis.Run()

	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	storeService.redisClient = redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	// calling store methods
	AddEncodedURL(short, original)

	decodedUrl := GetDecodedURL(short)

	assert.Equal(t, original, decodedUrl)

}
