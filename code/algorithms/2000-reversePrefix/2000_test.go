package leetcode

import (
	. "leetcode/pkg"
	"testing"
)

func Test2000(t *testing.T) {
	for _, test := range Tests() {
		t.Run(test.Name(), func(t *testing.T) {
			word := test.Val1.(string)
			ch := test.Val2.(string)[0]

			got := reversePrefix(word, ch)
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
			Val1: "abcdefd",
			Val2: "d",
			Want: "dcbaefd",
		},
		{
			Num:  2,
			Val1: "xyxzxe",
			Val2: "z",
			Want: "zxyxxe",
		},
		{
			Num:  3,
			Val1: "abcd",
			Val2: "z",
			Want: "abcd",
		},
		{
			Num:  4,
			Val1: "abcd",
			Val2: "d",
			Want: "dcba",
		},
	}
}
