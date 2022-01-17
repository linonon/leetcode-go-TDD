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
			got := isValid(v.TestData.(string))
			want := v.Want.(bool)

			if got != want {
				t.Errorf("got: %v, want: %v", got, want)
			}
		})
	}

}

func GetTestData() []pkgu.TestDataStruct {
	return []pkgu.TestDataStruct{
		{
			Num:      1,
			TestData: "()",
			Want:     true,
		},
		{
			Num:      2,
			TestData: "(){}[]",
			Want:     true,
		},
		{
			Num:      3,
			TestData: "({})",
			Want:     true,
		},
		{
			Num:      4,
			TestData: "({)}",
			Want:     false,
		},
		{
			Num:      5,
			TestData: "[",
			Want:     false,
		},
		{
			Num:      6,
			TestData: "]",
			Want:     false,
		},
		{
			Num:      7,
			TestData: "]]",
			Want:     false,
		},
	}
}
