package leetcode

import (
	pkgu "leetcode/pkg"
	"testing"
)

func TestStrStr(t *testing.T) {
	tests := GetTestData()
	for _, v := range tests {
		t.Run(v.GetNumName(), func(t *testing.T) {
			got := strStr(v.Val1.(string), v.Val2.(string))
			want := v.Want.(int)

			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}

func GetTestData() []pkgu.T {
	return []pkgu.T{
		{
			Num:    1,
			Val1: "hello",
			Val2: "ll",
			Want:   2,
		},
		{
			Num:    2,
			Val1: "aaaaaa",
			Val2: "bba",
			Want:   -1,
		},
	}
}
