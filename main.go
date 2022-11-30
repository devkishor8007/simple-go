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
	router.SetTrustedProxies([]string{"192.168.0.80"})
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

	router.GET("test/:name", func(c *gin.Context) {
		name := c.Param("name");
		c.String(http.StatusOK, "Hello %s", name);
	})

	router.GET("test/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + action
		c.String(http.StatusOK, message)
	})

	router.GET("doubleParam/:name/:address", func(c *gin.Context) {
		name := c.Param("name")
		address := c.Param("address")

		c.String(http.StatusOK, "hello %s %s", name, address);
	})

	router.GET("welcome", func(c *gin.Context) {
		name := c.Query("name");
		c.String(http.StatusOK, "Hello %s", name);
	})

	router.GET("useDefaultQuery", func(c *gin.Context) {
		name := c.DefaultQuery("name", "naames")
		c.String(http.StatusOK, "why man %s this", name)
	})

	router.GET("testy", func(c *gin.Context) {
		hello := c.DefaultQuery("hello", "hi")
		me := c.Query("me")
		c.String(http.StatusOK, "%s its %s", hello, me)
	})

	router.GET("newParam/:id", func(c *gin.Context) {
		id := c.Param("id");
		c.String(http.StatusOK, "%s", id)
	})

	router.POST("/post", func(c *gin.Context){
		id := c.Query("id")
		name := c.PostForm("name")
		// name := c.Request.Body("name")
		c.String(201, "id: %s; name: %s", id, name)
	})

	router.Run(":3605") // listen and serve on 0.0.0.0:8080
}

