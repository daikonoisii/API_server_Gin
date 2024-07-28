package routers

import (
	"github.com/gin-gonic/gin"
)

func SetupMathRouter(router *gin.Engine) {
	// "/math" 用のサブルーターを作成
	mathGroup := router.Group("/math")
	{
		FibonacciRouter(mathGroup)
	}
}
