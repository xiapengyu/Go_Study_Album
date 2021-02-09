package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()
	engine.StaticFS("/public", http.Dir("F:\\logs"))
	engine.Run(":8888")
}
