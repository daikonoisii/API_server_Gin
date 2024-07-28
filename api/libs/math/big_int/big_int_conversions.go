package bigint

import (
	"fmt"
	"math/big"
)

// ConvertFromString
// String型の整数をbig.Intに変換する関数
//
// Args:
//		numeric_string (string) : String型の整数
//
// Returns:
//		bigNum (*big.Int) : big.Int型の整数
//		err (error) : エラー
// Errorf:
//		"failed to parse big integer: %s": numericStringがbig.Intに変換できない文字列だった
func ConvertFromString(numericString string) (*big.Int, error) {
	// numericStringをbig.Int型の値bigNumへ変換
    bigNum, ok := new(big.Int).SetString(numericString, 10)

	// 正常に変換が行えたかを確認
    if !ok {
        return nil, fmt.Errorf("failed to parse big integer: %s", numericString)
    }
	
    return bigNum, nil
}
