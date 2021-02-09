package main

import (
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/gin-gonic/gin"
)

//定义中间件
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func LimitHandler(lmt *limiter.Limiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		httpError := tollbooth.LimitByRequest(lmt, c.Writer, c.Request)
		if httpError != nil {
			c.Data(httpError.StatusCode, lmt.GetMessageContentType(), []byte(httpError.Message))
			c.Abort()
		} else {
			c.Next()
		}
	}
}

func main() {
	engine := gin.Default()

	//rate-limit 中间件
	lmt := tollbooth.NewLimiter(1, nil)
	lmt.SetMessage("服务繁忙，请稍后再试...")

	//全局中间件
	engine.Use(Cors())
	engine.Use(LimitHandler(lmt))

	//api分组
	userGroup := engine.Group("/user")
	{
		//群组中间件
		userGroup.Use(Cors())
		userGroup.Use(LimitHandler(lmt))

		//单个中间件
		userGroup.GET("/hello", Cors(), LimitHandler(lmt), func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "Hello World!",
			})
		})

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

	engine.Run(":8888")
}
