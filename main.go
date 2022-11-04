package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies([]string{"192.168.0.117"})
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	});

	router.GET("/home", func(c *gin.Context) {
		fmt.Printf("ClientIp: %s\n", c.ClientIP())
	})
	router.Run() // listen and serve on 0.0.0.0:8080
}