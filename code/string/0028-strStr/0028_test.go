package leetcode

import (
	. "leetcode/pkg"
	"testing"
)

func TestStrStr(t *testing.T) {
	tests := GetTestData()
	for _, v := range tests {
		t.Run(v.TestName(), func(t *testing.T) {
			got := strStr(v.Val1.(string), v.Val2.(string))
			want := v.Want.(int)

			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}

func GetTestData() []T {
	return []T{
		{
			Num:  1,
			Val1: "hello",
			Val2: "ll",
			Want: 2,
		},
		{
			Num:  2,
			Val1: "aaaaaa",
			Val2: "bba",
			Want: -1,
		},
		{
			Num:  3,
			Val1: "",
			Val2: "",
			Want: 0,
		},
		{
			Num:  4,
			Val1: "",
			Val2: "a",
			Want: -1,
		},
		{
			Num:  5,
			Val1: "a",
			Val2: "",
			Want: 0,
		},
		{
			Num:  6,
			Val1: "a",
			Val2: "a",
			Want: 0,
		},
		{
			Num:  7,
			Val1: "aabaaab",
			Val2: "aabaaab",
			Want: 0,
		},
		{
			Num:  8,
			Val1: "aabbabbaabbba",
			Val2: "aabbabbaabbba",
			Want: 0,
		},
	}
}
