package leetcode

import (
	. "leetcode/pkg"
	"testing"
)

func Test1996(t *testing.T) {
	for _, test := range Tests() {
		t.Run(test.Name(), func(t *testing.T) {
			properties := test.Val.([][]int)

			got := numberOfWeakCharacters(properties)
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
			Val:  [][]int{{5, 5}, {6, 3}, {3, 6}},
			Want: 0,
		},
		{
			Num:  2,
			Val:  [][]int{{2, 2}, {3, 3}},
			Want: 1,
		},
		{
			Num:  3,
			Val:  [][]int{{1, 5}, {10, 4}, {4, 3}},
			Want: 1,
		},
		{
			Num:  4,
			Val:  [][]int{{7, 7}, {1, 2}, {9, 7}, {7, 3}, {3, 10}, {9, 8}, {8, 10}, {4, 3}, {1, 5}, {1, 5}},
			Want: 6,
		},
		{
			Num:  5,
			Val:  [][]int{{1, 1}, {100000000, 100000000}},
			Want: 1,
		},
	}
}
