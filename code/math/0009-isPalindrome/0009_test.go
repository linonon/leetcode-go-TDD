package leetcode

import (
	pkgu "leetcode/pkg"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	test := GetTestData()
	for i := range test {
		t.Run("Round", func(t *testing.T) {
			if isPalindrome(test[i].Value1.(int)) != test[i].Want1.(bool) {
				t.Error("Want != Got")
			}
		})
	}
}

func GetTestData() []pkgu.T {
	return []pkgu.T{
		{
			Value1: 111,
			Want1:     true,
		},
		{
			Value1: 121,
			Want1:     true,
		},
		{
			Value1: 1221,
			Want1:     true,
		},
		{
			Value1: 1271,
			Want1:     false,
		},
		{
			Value1: 123234028304928321,
			Want1:     false,
		},
		{
			Value1: 1111111111111111111,
			Want1:     true,
		},
	}
}
