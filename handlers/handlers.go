package handler

import (
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
	link.short_url = shortener.UrlShortener(link.Url)

	// store url and its shortened url in redis
	store.AddEncodedURL(link.short_url, link.Url)

	c.IndentedJSON(http.StatusOK, gin.H{"shortened_url": base_url + link.short_url})
}

// handler that is called when user
// gets into shortened url
func NavigateToLink(c *gin.Context) {

	shortened_url := c.Param("shortened_url")
	original_url := store.GetDecodedURL(shortened_url)

	validated_url, err := ValidateUrl(original_url)
	if err != nil {
		c.IndentedJSON(http.StatusOK, gin.H{"error": err})
	}

	c.Redirect(302, validated_url)
}

func WelcomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "WelcomePage.html", gin.H{})
}
