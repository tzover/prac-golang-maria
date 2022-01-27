package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


func main() {
	fmt.Println("Hello golang from docker!")
	fmt.Println("Hot reload")
	fmt.Println("Hot reload2")


	engine:= gin.Default()
	engine.GET("/", func(c * gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello golang from docker!",
		})
	})
	engine.Run(":8080")
}