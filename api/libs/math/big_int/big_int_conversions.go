package bigint

import "math/big"

func ConvertFromString(s string) *big.Int {
	bigNum, ok := new(big.Int).SetString(s, 10)
	if !ok {
		panic("failed to parse big integer")
	}
	return bigNum
}