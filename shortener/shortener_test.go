package shortener

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncoder(t *testing.T) {
	original_url := "https://club.snapp.ir/radio-snapp/newseason/"
	short_url := UrlShortener(original_url)

	assert.Equal(t, short_url, "d66yfx7NNw7")

}
