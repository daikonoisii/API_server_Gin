package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	fibonacci "API_server_Gin/api/libs/math/fibonacci"
)

func GetFibonacciNumber(responseWriter http.ResponseWriter, request *http.Request) {
	// クエリからfibonacci_indexの値を取得
	fibonacciIndexStr := request.URL.Query().Get("fibonacci_index")

	// fibonacci_indexをintへ変換
	fibonacciIndex, err := strconv.Atoi(fibonacciIndexStr)
	if err != nil {
		http.Error(responseWriter, "Invalid number format", http.StatusBadRequest)
		return
	}

	// フィボナッチ数を計算
	result, err := fibonacci.GenerateFibonacciNumber(int16(fibonacciIndex))

	// GenerateFibonacciNumberがリクセスとパラーメータを扱えなかったため、422エラーを発生させる
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	// レスポンスを作成し、Jsonへエンコード
	resp := map[string]string{"fibonacci_number": result.String()}
	responseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(responseWriter).Encode(resp)
}