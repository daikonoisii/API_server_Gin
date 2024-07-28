package routers

import (
	"github.com/gin-gonic/gin"
	
	handler "API_server_Gin/api/handler/math/fibonacci"
)

func FibonacciRouter(rg *gin.RouterGroup) {
	// "/fib" パスのサブルーターを作成
	fibGroup := rg.Group("/fib")
	{
		fibGroup.GET("/", handler.GetFibonacciNumber)
	}
}
