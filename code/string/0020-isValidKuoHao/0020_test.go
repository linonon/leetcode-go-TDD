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
			got := isValid(v.Value1.(string))
			want := v.Want1.(bool)

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
			Value1: "()",
			Want1:     true,
		},
		{
			Num:      2,
			Value1: "(){}[]",
			Want1:     true,
		},
		{
			Num:      3,
			Value1: "({})",
			Want1:     true,
		},
		{
			Num:      4,
			Value1: "({)}",
			Want1:     false,
		},
		{
			Num:      5,
			Value1: "[",
			Want1:     false,
		},
		{
			Num:      6,
			Value1: "]",
			Want1:     false,
		},
		{
			Num:      7,
			Value1: "]]",
			Want1:     false,
		},
	}
}
