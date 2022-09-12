package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		profile := os.Getenv("PROFILE")
		msg := fmt.Sprintf("Hello, PROFILE is %s, Welcome Gin World!", profile)
		c.JSON(http.StatusOK, msg)
	})
	err := router.Run(":9000")
	if err != nil {
		panic(err)
	}
}
