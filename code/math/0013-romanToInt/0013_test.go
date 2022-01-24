package leetcode

import (
	pkgu "leetcode/pkg"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	testData := GetTestData()
	for i := range testData {
		t.Run("Round", func(t *testing.T) {

			got := romanToInt(testData[i].Val1.(string))
			want := testData[i].Want.(int)
			if got != want {
				t.Errorf("Fail: TestData(%v), Want(%v)", got, want)
			}

		})
	}
}

func GetTestData() []pkgu.T {
	return []pkgu.T{
		// Normal
		{
			Val1: "I",
			Want:     1,
		},
		{
			Val1: "V",
			Want:     5,
		},
		// More
		{
			Val1: "II",
			Want:     2,
		},
		{
			Val1: "XI",
			Want:     11,
		},
		// More More
		{
			Val1: "MDCLXVI",
			Want:     1000 + 500 + 100 + 50 + 10 + 5 + 1,
		},
		// 4,9 test
		{
			Val1: "IV",
			Want:     4,
		},
		{
			Val1: "IX",
			Want:     9,
		},
	}

}
