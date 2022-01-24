package leetcode

import (
	pkgu "leetcode/pkg"
	"strconv"
	"testing"
)

func TestIsValidKuoHao(t *testing.T) {
	test := GetTestData()
	for _, v := range test {
		t.Run(strconv.Itoa(v.Num), func(t *testing.T) {
			got := isValid(v.Val1.(string))
			want := v.Want.(bool)

			if got != want {
				t.Errorf("got: %v, want: %v", got, want)
			}
		})
	}

}

func GetTestData() []pkgu.T {
	return []pkgu.T{
		{
			Num:      1,
			Val1: "()",
			Want:     true,
		},
		{
			Num:      2,
			Val1: "(){}[]",
			Want:     true,
		},
		{
			Num:      3,
			Val1: "({})",
			Want:     true,
		},
		{
			Num:      4,
			Val1: "({)}",
			Want:     false,
		},
		{
			Num:      5,
			Val1: "[",
			Want:     false,
		},
		{
			Num:      6,
			Val1: "]",
			Want:     false,
		},
		{
			Num:      7,
			Val1: "]]",
			Want:     false,
		},
	}
}
