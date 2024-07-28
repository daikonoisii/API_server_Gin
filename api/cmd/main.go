package main

import (
	"github.com/gin-gonic/gin"

	routers "API_server_Gin/api/routers"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Gin",
		})
	})

	routers.SetupMathRouter(router)
	
	// 0.0.0.0:8080 でサーバーを立てる
	router.Run()
}