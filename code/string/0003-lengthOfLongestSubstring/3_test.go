package leetcode

import (
	. "leetcode/pkg"
	"testing"
)

func Test3(t *testing.T) {
	for _, test := range Tests() {
		t.Run(test.Name(), func(t *testing.T) {
			s := test.Val.(string)

			got := lengthOfLongestSubstring(s)
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
			Val:  "abcabcbb",
			Want: 3,
		},
		{
			Num:  2,
			Val:  "bbbbb",
			Want: 1,
		},
		{
			Num:  3,
			Val:  "pwwkew",
			Want: 3,
		},
		{
			Num:  4,
			Val:  "",
			Want: 0,
		},
		{
			Num:  5,
			Val:  "dvdf",
			Want: 3,
		},
	}
}
