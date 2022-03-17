package leetcode

import (
	. "leetcode/pkg"
	"testing"
)

func Test720(t *testing.T) {
	for _, test := range Tests() {
		t.Run(test.Name(), func(t *testing.T) {
			words := test.Val.([]string)

			got := longestWord(words)
			want := test.Want.(string)

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
			Val:  []string{"w", "wo", "wor", "worl", "world"},
			Want: "world",
		},
		{
			Num:  2,
			Val:  []string{"a", "banana", "app", "appl", "ap", "apply", "apple"},
			Want: "apple",
		},
	}
}
