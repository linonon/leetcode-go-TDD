package leetcode

import (
	pkgu "leetcode/pkg"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	testData := GetTestData()
	for i := range testData {
		t.Run("Round", func(t *testing.T) {

			got := romanToInt(testData[i].Value1.(string))
			want := testData[i].Want1.(int)
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
			Value1: "I",
			Want1:     1,
		},
		{
			Value1: "V",
			Want1:     5,
		},
		// More
		{
			Value1: "II",
			Want1:     2,
		},
		{
			Value1: "XI",
			Want1:     11,
		},
		// More More
		{
			Value1: "MDCLXVI",
			Want1:     1000 + 500 + 100 + 50 + 10 + 5 + 1,
		},
		// 4,9 test
		{
			Value1: "IV",
			Want1:     4,
		},
		{
			Value1: "IX",
			Want1:     9,
		},
	}

}
