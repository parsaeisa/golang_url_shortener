package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/parsaeisa/technical_test_golang/shortener"
	"github.com/parsaeisa/technical_test_golang/store"
)

type Link struct {
	short_url string
	Url       string
}

const base_url = "http://localhost:8080/"

// handler for shortening url in request
// and return it in response
// I have binded request to Link structure
func CreateShortUrl(c *gin.Context) {

	var link Link

	if err := c.BindJSON(&link); err != nil {
		return
	}

	// call urlshortener method to apply encodings
	validated_url, err := ValidateUrl(link.Url)
	fmt.Println(validated_url)
	if err != nil {
		c.IndentedJSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(validated_url)
	link.short_url = shortener.UrlShortener(validated_url)

	// store url and its shortened url in redis
	store.AddEncodedURL(link.short_url, validated_url)

	c.IndentedJSON(http.StatusOK, gin.H{"shortened_url": base_url + link.short_url})
}

// handler that is called when user
// gets into shortened url
func NavigateToLink(c *gin.Context) {

	shortened_url := c.Param("shortened_url")
	original_url := store.GetDecodedURL(shortened_url)

	c.Redirect(302, original_url)
}

func WelcomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "WelcomePage.html", gin.H{})
}
