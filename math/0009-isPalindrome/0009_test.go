package leetcode

import (
	pkgu "leetcode/pkg"
	"testing"
)

type TestDataStruct struct {
	TestData int
	Want     bool
}

func TestIsPalindrome(t *testing.T) {
	test := GetTestData()
	for i := range test {
		t.Run("Round", func(t *testing.T) {
			p := pkgu.NewPerf()
			if isPalindrome(test[i].TestData) != test[i].Want {
				t.Error("Want != Got")
			}
			t.Log("Perf:", p.Ns())
		})
	}
}

func GetTestData() []TestDataStruct {
	return []TestDataStruct{
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
