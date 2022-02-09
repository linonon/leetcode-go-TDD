package leetcode

import (
	. "leetcode/pkg"
	"testing"
)

func Test2006(t *testing.T) {
	for _, test := range Tests() {
		t.Run(test.Name(), func(t *testing.T) {
			nums := test.Val1.([]int)
			k := test.Val2.(int)

			got := countKDifference(nums, k)
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
			Val1: []int{1, 2, 2, 1},
			Val2: 1,
			Want: 4,
		},
		{
			Num:  2,
			Val1: []int{1, 3},
			Val2: 3,
			Want: 0,
		},
		{
			Num:  3,
			Val1: []int{3, 2, 1, 5, 4},
			Val2: 2,
			Want: 3,
		},
	}
}
