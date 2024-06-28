package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// gin custom middleware
func myLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get authorization header
		token := c.Request.Header.Get("Authorization")
		if len(token) == 0 {
			fmt.Println("no token")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization header is required"})
			return
		}

		c.Next()
		// Do something after request
	}
}

func main() {
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Use(myLogger())

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.Run(":8080")
}
