package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRedisConnection(t *testing.T) {

	c := ConnectToRedis()

	assert.True(t, c.redisClient != nil)
}
