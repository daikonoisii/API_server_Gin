package bigint_test

import(
	"errors"
	"math/big"
	"testing"

	bigint "API_server_Gin/api/libs/math/big_int"
)

func TestConvertFromString(t *testing.T){
	//テストケースを定義
	tests := []struct {
		name			string		// テストケース名
		numericString	string		// big.Int型に変換したい文字列
		expected		*big.Int	// 変換結果
		expectedErr		error		// エラー
	}{
		{
			name:          "SmallValue",
			numericString: "1",
			expected:      big.NewInt(1),
			expectedErr:   nil,
		},
		{
			name:          "BeyondIntMaxValue",
			numericString: "9223372036854775808",
			expected: func() *big.Int {
				n, _ := new(big.Int).SetString("9223372036854775808", 10)
				return n
			}(),
			expectedErr: nil,
		},
		{
			name:          "BeyondUintMaxValue",
			numericString: "18446744073709551616",
			expected: func() *big.Int {
				n, _ := new(big.Int).SetString("18446744073709551616", 10)
				return n
			}(),
			expectedErr: nil,
		},
		{
			name:          "NegativeValue",
			numericString: "-1",
			expected:      big.NewInt(-1),
			expectedErr:   nil,
		},
		{
			name:          "DecimalValue",
			numericString: "123.456",
			expected:      nil,
			expectedErr:   errors.New("failed to parse big integer: 123.456"),
		},
		{
			name:          "Text",
			numericString: "abc",
			expected:      nil,
			expectedErr:   errors.New("failed to parse big integer: abc"),
		},
		{
			name:          "EmptyString",
			numericString: "",
			expected:      nil,
			expectedErr:   errors.New("failed to parse big integer: "),
		},
		{
			name:          "ZeroValue",
			numericString: "0",
			expected:      big.NewInt(0),
			expectedErr:   nil,
		},
		{
			name:          "LargeNegativeValue",
			numericString: "-9223372036854775809",
			expected: func() *big.Int {
				n, _ := new(big.Int).SetString("-9223372036854775809", 10)
				return n
			}(),
			expectedErr: nil,
		},
		{
			name:          "VeryLargeValue",
			numericString: "123456789012345678901234567890",
			expected: func() *big.Int {
				n, _ := new(big.Int).SetString("123456789012345678901234567890", 10)
				return n
			}(),
			expectedErr: nil,
		},
		{
			name:          "TextAndNumeric",
			numericString: "1234abc5678",
			expected:      nil,
			expectedErr:   errors.New("failed to parse big integer: 1234abc5678"),
		},

	}

	//全てのテストケースに対してテストを実行
	//全てのテストケースに対してテストを実行
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := bigint.ConvertFromString(test.numericString)
			
			if err != nil {
				// エラーが発生すべきではないケースでエラーが発生
				if test.expectedErr == nil {
					t.Errorf("%s: expected no error, got %v", test.name, err)
				// 発生したエラーが想定と異なる
				} else if err.Error() != test.expectedErr.Error() {
					t.Errorf("%s: expected error %v, got %v", test.name, test.expectedErr, err)
				}
			// エラーが発生すべきなケースでエラーが発生しなかった
			} else if test.expectedErr != nil {
				t.Errorf("%s: expected error %v, got no error", test.name, test.expectedErr)
			}

			// 返り値が想定と異なる
			if result.Cmp(test.expected) != 0 {
				t.Errorf("%s: expected result %s, got %s", test.name, test.expected.String(), result.String())
			}
		})
	}
}
