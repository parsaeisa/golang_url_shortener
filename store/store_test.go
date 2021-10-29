package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRedisConnection(t *testing.T) {

	c := ConnectToRedis()

	assert.True(t, c.redisClient != nil)
}

func TestStoreToRedis(t *testing.T) {
	original := "https://club.snapp.ir/radio-snapp/newseason/"
	short := "http://localhost:8080/s8yQsd"
	userId := "parsa"

	AddEncodedURL(short, original, userId)

	decodedUrl := GetDecodedURL(short)

	assert.Equal(t, original, decodedUrl)

}
