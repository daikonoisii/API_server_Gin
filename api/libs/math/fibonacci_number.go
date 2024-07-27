package fibonacci

import (
	"fmt"
)

// GenerateFibonacciNumber
// 指定された順番のフィボナッチ数を生成する関数
// ex. n=1 -> 1, n=5 -> 8, n=6 -> 13 n=7 -> 21
// Args:
//
//	index (int) : フィボナッチ数列の順番を指定する値.1で先頭のフィボナッチ数を指定する.
//
// Returns:
//
//	fibonacciNumber (int) : 指定された番目のフィボナッチ数.fibonacci_index = 0 の返り値は0.
//	err (error) : エラー
func GenerateFibonacciNumber(index int16) (fibonacciNumber int, err error) {
	if index > 0 {
		fibonacciNumber := 1
		nextValue := 1

		for i := int16(1); i < index; i++ {
			fibonacciNumber, nextValue = nextValue, nextValue+fibonacciNumber
		}

		return fibonacciNumber, nil

	} else if index < 0 {
		return 0, fmt.Errorf("invalid index: %d", index)

	} else {
		return 0, nil
	}
}
