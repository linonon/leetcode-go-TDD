package leetcode

import (
	. "leetcode/pkg"
	"testing"
)

func Test1414(t *testing.T) {
	for _, test := range Tests() {
		t.Run(test.Name(), func(t *testing.T) {
			k := test.Val.(int)

			got := findMinFibonacciNumbers(k)
			want := test.Want.(int)

			if got != want {
				t.Errorf("got %v, want %v\n", got, want)
			}
		})
	}
}

func Tests() []T {
	return []T{
		{
			Num:  1,
			Val:  7,
			Want: 2,
		},
		{
			Num:  2,
			Val:  10,
			Want: 2,
		},
		{
			Num:  3,
			Val:  19,
			Want: 3,
		},
	}
}
