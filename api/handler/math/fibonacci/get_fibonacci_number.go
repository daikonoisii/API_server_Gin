// get_fibonacci_number.go
package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	
	fibonacci "API_server_Gin/api/libs/math/fibonacci"
)

func GetFibonacciNumber(c *gin.Context) {
	// クエリからfibonacci_indexの値を取得
	fibonacciIndexStr := c.Query("fibonacci_index")

	// fibonacci_indexをintへ変換
	fibonacciIndex, err := strconv.Atoi(fibonacciIndexStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid number format"})
		return
	}

	// フィボナッチ数を計算
	result, err := fibonacci.GenerateFibonacciNumber(int16(fibonacciIndex))

	// GenerateFibonacciNumberがリクセスとパラーメータを扱えなかったため、422エラーを発生させる
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	// レスポンスを作成し、Jsonへエンコード
	c.JSON(http.StatusOK, gin.H{"fibonacci_number": result.String()})
}
