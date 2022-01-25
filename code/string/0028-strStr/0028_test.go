package leetcode

import (
	. "leetcode/pkg"
	"testing"
)

func TestStrStr(t *testing.T) {
	tests := GetTestData()
	for _, v := range tests {
		t.Run(v.Name(), func(t *testing.T) {
			haystack, needle := v.V[0].(string), v.V[1].(string)

			got := strStr(haystack, needle)
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
			V:    SetM("hello", "ll"),
			Want: 2,
		},
		{
			Num:  2,
			V:    SetM("aaaaaa", "bba"),
			Want: -1,
		},
		{
			Num:  3,
			V:    SetM("", ""),
			Want: 0,
		},
		{
			Num:  4,
			V:    SetM("", "a"),
			Want: -1,
		},
		{
			Num:  5,
			V:    SetM("a", ""),
			Want: 0,
		},
		{
			Num:  6,
			V:    SetM("a", "a"),
			Want: 0,
		},
		{
			Num:  7,
			V:    SetM("aabaaab", "aabaaab"),
			Want: 0,
		},
		{
			Num:  8,
			V:    SetM("aabbabbaabbba", "aabbabbaabbba"),
			Want: 0,
		},
	}
}
