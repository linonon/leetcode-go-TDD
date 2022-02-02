package leetcode

import (
	. "leetcode/pkg"
	"testing"
)

func Test0004(t *testing.T) {
	for _, test := range Tests() {
		t.Run(test.Name(), func(t *testing.T) {
			nums1, nums2 := test.Val1.([]int), test.Val2.([]int)

			got := findMedianSortedArrays(nums1, nums2)
			want := test.Want.(float64)

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
			Val1: []int{1, 3},
			Val2: []int{2},
			Want: 2.00000,
		},
		{
			Num:  2,
			Val1: []int{1, 2},
			Val2: []int{3, 4},
			Want: 2.50000,
		},
		{
			Num:  3,
			Val1: []int{0, 0},
			Val2: []int{0, 0},
			Want: 0.00000,
		},
		{
			Num:  4,
			Val1: []int{},
			Val2: []int{1},
			Want: 1.00000,
		},
		{
			Num:  5,
			Val1: []int{1, 2, 3},
			Val2: []int{},
			Want: 2.00000,
		},
		{
			Num:  6,
			Val1: []int{1, 3, 5},
			Val2: []int{2, 4, 6},
			Want: 3.50000,
		},
		{
			Num:  7,
			Val1: []int{},
			Val2: []int{2, 3},
			Want: 2.50000,
		},
		{
			Num:  8,
			Val1: []int{3},
			Val2: []int{-2, -1},
			Want: -1.0000,
		},
	}
}
