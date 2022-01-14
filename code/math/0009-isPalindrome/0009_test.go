package leetcode

import (
	pkgu "leetcode/pkg"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	test := GetTestData()
	for i := range test {
		t.Run("Round", func(t *testing.T) {
			if isPalindrome(test[i].TestData.(int)) != test[i].Want.(bool) {
				t.Error("Want != Got")
			}
		})
	}
}

func GetTestData() []pkgu.TestDataStruct {
	return []pkgu.TestDataStruct{
		{
			TestData: 111,
			Want:     true,
		},
		{
			TestData: 121,
			Want:     true,
		},
		{
			TestData: 1221,
			Want:     true,
		},
		{
			TestData: 1271,
			Want:     false,
		},
		{
			TestData: 123234028304928321,
			Want:     false,
		},
		{
			TestData: 1111111111111111111,
			Want:     true,
		},
	}
}
