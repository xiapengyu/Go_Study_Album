package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	engine := gin.Default()

	//api分组
	userGroup := engine.Group("/user")
	{
		userGroup.GET("/ping", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "pong",
			})
		})
		userGroup.GET("/UserInfo", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"name":    "Kobe Bryant",
				"age":     23,
				"address": "China",
			})
		})
	}

	goodsGroup := engine.Group("/goods")
	{
		goodsGroup.GET("/ping", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "pong",
			})
		})
		goodsGroup.GET("/GoodsInfo", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"name":    "Kobe Bryant",
				"age":     23,
				"address": "China",
			})
		})
	}

	//engine.Run(":8888")

	srv := &http.Server{
		Addr:         ":8888",
		Handler:      engine,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	srv.ListenAndServe()
}
