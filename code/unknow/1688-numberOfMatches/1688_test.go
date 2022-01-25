package leetcode

import (
	. "leetcode/pkg"
	"testing"
)

func Test1688(t *testing.T) {
	for _, test := range Tests() {
		t.Run(test.Name(), func(t *testing.T) {
			n := test.Val.(int)

			got := numberOfMatches(n)
			want := test.Want.(int)

			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}

func Tests() []T {
	return []T{
		{
			Num:  1,
			Val:  7,
			Want: 6,
		},
		{
			Num:  2,
			Val:  14,
			Want: 13,
		},
	}
}
