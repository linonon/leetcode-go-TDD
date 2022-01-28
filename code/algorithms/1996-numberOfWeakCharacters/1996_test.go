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
	}
}
