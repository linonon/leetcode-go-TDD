package leetcode

import (
	. "leetcode/pkg"
	"reflect"
	"testing"
)

func Test661(t *testing.T) {
	for _, test := range Tests() {
		t.Run(test.Name(), func(t *testing.T) {
			img := test.Val.([][]int)

			got := imageSmoother(img)
			want := test.Want.([][]int)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v, want %v\n", got, want)
			}
		})
	}
}

func Tests() []T {
	return []T{
		{
			Num:  1,
			Val:  [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}},
			Want: [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
		},
		{
			Num:  2,
			Val:  [][]int{{100, 200, 100}, {200, 50, 200}, {100, 200, 100}},
			Want: [][]int{{137, 141, 137}, {141, 138, 141}, {137, 141, 137}},
		},
		{
			Num:  3,
			Val:  [][]int{{1}},
			Want: [][]int{{1}},
		},
		{
			Num:  4,
			Val:  [][]int{{1,0,1}},
			Want: [][]int{{0,0,0}},
		},
	}
}
