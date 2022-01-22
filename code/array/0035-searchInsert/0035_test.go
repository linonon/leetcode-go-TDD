package leetcode

import (
	. "leetcode/pkg"
	"testing"
)

func TestSearchInsert(t *testing.T) {
	for _, v := range GetTestData() {
		TestRun(t, v.Num, func(t *testing.T) {

			got := searchInsert(v.Value.([]int), v.Target)
			want := v.Want.(int)

			if got != want {
				t.Errorf("SearchInsert: got %d, want %d", got, want)
			}
		})
	}
}

func GetTestData() []T {
	return []T{
		{
			Num:    1,
			Value:  []int{1, 3, 5, 6},
			Target: 5,
			Want:   2,
		},
		{
			Num:    2,
			Value:  []int{1, 3, 5, 6},
			Target: 2,
			Want:   1,
		},
		{
			Num:    3,
			Value:  []int{1, 3, 5, 6},
			Target: 7,
			Want:   4,
		},
		{
			Num:    4,
			Value:  []int{1, 3, 5, 6},
			Target: 0,
			Want:   0,
		},
		{
			Num:    5,
			Value:  []int{1},
			Target: 0,
			Want:   0,
		},
		{
			Num:    6,
			Value:  []int{1, 3},
			Target: 3,
			Want:   1,
		},
	}
}
