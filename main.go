package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
)

type todo struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Description string `json:"Description"`
	IsShare bool `json:"IsShare"`
	Price  float64 `json:"price"`
}
var todoArray = []todo {
	{ ID: "1", Title: "first", Description: "this could be first", IsShare: true, Price: 15.2 },
	{ ID: "2", Title: "second", Description: "this could be second", IsShare: false, Price: 56.55 },
}


func getTodos(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, todoArray)
}

func getTodoById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range todoArray {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
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
	router.GET("/todo/:id", getTodoById)

	router.Run() // listen and serve on 0.0.0.0:8080
}

