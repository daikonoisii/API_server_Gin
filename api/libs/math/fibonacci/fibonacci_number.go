package fibonacci

import (
	"fmt"

	"math/big"
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
func GenerateFibonacciNumber(index int16) (fibonacciNumber *big.Int, err error) {
	if index > 0 {
		
		// フィボナッチ数を big.Int で計算
		fibonacciNumber = big.NewInt(1)
		nextValue := big.NewInt(1)
		tmp := big.NewInt(0) // 一時的な計算結果を格納する変数

		for i := int16(1); i < index; i++ {
			// fibonacciNumber + nextValue を tmp に保存
			tmp.Add(fibonacciNumber, nextValue)
			// nextValue を fibonacciNumber に設定
			fibonacciNumber.Set(nextValue)
			// tmp を nextValue に設定
			nextValue.Set(tmp)
		}


		return fibonacciNumber, nil

	} else if index < 0 {
		return big.NewInt(0), fmt.Errorf("invalid index: %d", index)

	} else {
		return big.NewInt(0), nil
	}
}
