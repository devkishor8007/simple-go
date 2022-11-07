package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
)

type todo struct {
	id string
	title string
	description string
	isShare bool
}



func getTodos(c *gin.Context) {
	var todoArray = []todo {
		{ id: "1", title: "first", description: "this could be first", isShare: true },
		{ id: "2", title: "second", description: "this could be second", isShare: false },
	}
	c.IndentedJSON(http.StatusOK, todoArray)
}



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

	router.GET("/home/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		fmt.Printf(id);
		c.JSON(http.StatusOK, gin.H{
			"params": id,
			"status": "true",
		})
	})

	router.GET("/", getTodos)

	router.Run() // listen and serve on 0.0.0.0:8080
}