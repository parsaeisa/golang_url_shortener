package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	handler "github.com/parsaeisa/technical_test_golang/handlers"
	"github.com/parsaeisa/technical_test_golang/store"
)

func main() {
	//define router
	router := gin.Default()

	// load html files
	router.LoadHTMLGlob("html/*")

	// routings
	router.GET("/", handler.WelcomePage)
	router.POST("/get_shortened_url", handler.CreateShortUrl)
	router.GET("/:shortened_url", handler.NavigateToLink)

	store.ConnectToRedis()

	err := router.Run(":8080")
	if err != nil {
		panic(fmt.Sprintf("Failed"))
	}
}
