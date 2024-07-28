package fibonacci_test

import (
	"errors"
	"math/big"
	"testing"

	fibonacci "API_server_Gin/api/libs/math/fibonacci"
	bigint "API_server_Gin/api/libs/math/big_int"
)



// TestGenerateFibonacciNumber
// GenerateFibonacciNumberのテスト
//
// GenerateFibonacciNumber:
// 	Args:
//
//		index (int) : フィボナッチ数列の順番を指定する値.1で先頭のフィボナッチ数を指定する.
//
// 	Returns:
//
//		fibonacciNumber (int) : 指定された番目のフィボナッチ数.fibonacci_index = 0 の返り値は0.
//		err (error) : エラー
func TestGenerateFibonacciNumber(t *testing.T) {



	tests := []struct {
		name        string
		index       int16
		expected    *big.Int
		expectedErr error
	}{
		{"FirstFibonacciNumber", 1, big.NewInt(1), nil},
		{"SecondFibonacciNumber", 2, big.NewInt(1), nil},
		{"ThirdFibonacciNumber", 3, big.NewInt(2), nil},
		{"FourthFibonacciNumber", 4, big.NewInt(3), nil},
		{"FifthFibonacciNumber", 5, big.NewInt(5), nil},
		{"TenthFibonacciNumber", 10, big.NewInt(55), nil},
		{"ZeroIndex", 0, big.NewInt(0), nil},
		{"NegativeIndex", -1, big.NewInt(0), errors.New("invalid index: -1")},
		{"LargeIndex", 99, bigint.ConvertFromString("218922995834555169026"), nil},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := fibonacci.GenerateFibonacciNumber(test.index)

			if err != nil {
				if test.expectedErr == nil {
					t.Errorf("%s: expected no error, got %v", test.name, err)
				} else if err.Error() != test.expectedErr.Error() {
					t.Errorf("%s: expected error %v, got %v", test.name, test.expectedErr, err)
				}
			} else if test.expectedErr != nil {
				t.Errorf("%s: expected error %v, got no error", test.name, test.expectedErr)
			}

			if result.Cmp(test.expected) != 0 {
				t.Errorf("%s: expected result %s, got %s", test.name, test.expected.String(), result.String())
			}
		})
	}
}
