package leetcode

import (
	pkgu "leetcode/pkg"
	"testing"
)

func TestStrStr(t *testing.T) {
	tests := GetTestData()
	for _, v := range tests {
		t.Run(v.GetNumName(), func(t *testing.T) {
			got := strStr(v.Value1.(string), v.Value2.(string))
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
			Value1: "hello",
			Value2: "ll",
			Want:   2,
		},
		{
			Num:    2,
			Value1: "aaaaaa",
			Value2: "bba",
			Want:   -1,
		},
	}
}
