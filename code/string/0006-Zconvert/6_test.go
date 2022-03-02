package leetcode

import (
	. "leetcode/pkg"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test6(t *testing.T) {
	for _, test := range Tests() {
		t.Run(test.Name(), func(t *testing.T) {
			s := test.Val1.(string)
			numRows := test.Val2.(int)
			actual := convert(s, numRows)
			expected := test.Want.(string)

			assert.Equal(t, expected, actual)
		})
	}
}

func Tests() []T {

	return []T{
		{
			Num:  1,
			Val1: "PAYPALISHIRING",
			Val2: 3,
			Want: "PAHNAPLSIIGYIR",
		},
		{
			Num:  2,
			Val1: "PAYPALISHIRING",
			Val2: 4,
			Want: "PINALSIGYAHRPI",
		},
		{
			Num:  3,
			Val1: "A",
			Val2: 1,
			Want: "A",
		},
		{
			Num:  4,
			Val1: "AB",
			Val2: 1,
			Want: "AB",
		},
		{
			Num:  5,
			Val1: "ABC",
			Val2: 2,
			Want: "ACB",
		},
		{
			Num:  6,
			Val1: "ABCDE",
			Val2: 3,
			Want: "AEBDC",
		},
	}
}
