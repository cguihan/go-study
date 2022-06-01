package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	err := r.Run(":8081")
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080

	//v1
	fmt.Println("graham")
	fmt.Println("graham v2")
	fmt.Println("graham v3")
}
