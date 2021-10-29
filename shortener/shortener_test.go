package shortener

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncoder(t *testing.T) {
	userId := "e0dba740-fc4b-4977-872c-d360239e6b1a"
	original_url := "https://www.eddywm.com/lets-build-a-url-shortener-in-go-with-redis-part-2-storage-layer/"
	short_url := UrlShortener(original_url, userId)

	assert.Equal(t, short_url, "d66yfx7NNw7")

}
