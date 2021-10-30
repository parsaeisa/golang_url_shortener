package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/parsaeisa/technical_test_golang/shortener"
	"github.com/parsaeisa/technical_test_golang/store"
)

type request struct {
	url string
}

const base_url = "http://localhost:8080/"

func CreateShortUrl(c *gin.Context) {

	var newRequest request

	if err := c.BindJSON(&newRequest); err != nil {
		return
	}

	shortened_url := shortener.UrlShortener(newRequest.url, "null")
	store.AddEncodedURL(shortened_url, newRequest.url, "0")

	c.IndentedJSON(http.StatusNotFound, gin.H{"shortened_url": base_url + shortened_url})
}
