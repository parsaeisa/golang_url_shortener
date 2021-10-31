package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/parsaeisa/technical_test_golang/shortener"
	"github.com/parsaeisa/technical_test_golang/store"
)

type request struct {
	URL string
}

const base_url = "http://localhost:8080/"

func CreateShortUrl(c *gin.Context) {

	var newRequest request

	if err := c.BindJSON(&newRequest); err != nil {
		return
	}

	shortened_url := shortener.UrlShortener(newRequest.URL, "null")
	store.AddEncodedURL(shortened_url, newRequest.URL, "0")

	c.IndentedJSON(http.StatusOK, gin.H{"shortened_url": base_url + shortened_url})
}

func NavigateToLink(c *gin.Context) {
	shortened_url := c.Param("shortened_url")

	original_url := store.GetDecodedURL(shortened_url)

	c.Redirect(302, original_url)
}
