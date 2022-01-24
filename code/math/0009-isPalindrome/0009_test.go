package leetcode

import (
	pkgu "leetcode/pkg"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	test := GetTestData()
	for i := range test {
		t.Run("Round", func(t *testing.T) {
			if isPalindrome(test[i].Val1.(int)) != test[i].Want.(bool) {
				t.Error("Want != Got")
			}
		})
	}
}

func GetTestData() []pkgu.T {
	return []pkgu.T{
		{
			Val1: 111,
			Want:     true,
		},
		{
			Val1: 121,
			Want:     true,
		},
		{
			Val1: 1221,
			Want:     true,
		},
		{
			Val1: 1271,
			Want:     false,
		},
		{
			Val1: 123234028304928321,
			Want:     false,
		},
		{
			Val1: 1111111111111111111,
			Want:     true,
		},
	}
}
