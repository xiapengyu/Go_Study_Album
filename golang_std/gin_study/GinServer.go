package main

import "github.com/gin-gonic/gin"

func main() {
	engine := gin.Default()
	engine.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "Pong",
		})
	})

	engine.Run(":8888")
}
