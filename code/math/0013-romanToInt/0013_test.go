package leetcode

import (
	pkgu "leetcode/pkg"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	testData := GetTestData()
	for i := range testData {
		t.Run("Round", func(t *testing.T) {

			got := romanToInt(testData[i].TestData.(string))
			want := testData[i].Want.(int)
			if got != want {
				t.Errorf("Fail: TestData(%v), Want(%v)", got, want)
			}

		})
	}
}

func GetTestData() []pkgu.TestDataStruct {
	return []pkgu.TestDataStruct{
		// Normal
		{
			TestData: "I",
			Want:     1,
		},
		{
			TestData: "V",
			Want:     5,
		},
		// More
		{
			TestData: "II",
			Want:     2,
		},
		{
			TestData: "XI",
			Want:     11,
		},
		// More More
		{
			TestData: "MDCLXVI",
			Want:     1000 + 500 + 100 + 50 + 10 + 5 + 1,
		},
		// 4,9 test
		{
			TestData: "IV",
			Want:     4,
		},
		{
			TestData: "IX",
			Want:     9,
		},
	}

}
