package leetcode

import (
	. "leetcode/pkg"
	"testing"
)

func Test53(t *testing.T) {
	for _, test := range Tests() {
		t.Run(test.Name(), func(t *testing.T) {
			nums := test.Val.([]int)

			got := maxSubArray(nums)
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
			Val:  []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			Want: 6,
		},
		{
			Num:  2,
			Val:  []int{1},
			Want: 1,
		},
		{
			Num:  3,
			Val:  []int{5, 4, -1, 7, 8},
			Want: 23,
		},
		{
			Num:  4,
			Val:  []int{-2, -1},
			Want: -1,
		},
	}
}
