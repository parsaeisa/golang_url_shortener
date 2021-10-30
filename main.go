package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	handler "github.com/parsaeisa/technical_test_golang/handlers"
	"github.com/parsaeisa/technical_test_golang/store"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "first Message",
		})
	})

	r.GET("/:shortened_url", handler.NavigateToLink)

	store.ConnectToRedis()

	err := r.Run(":8080")
	if err != nil {
		panic(fmt.Sprintf("Failed"))
	}
}
