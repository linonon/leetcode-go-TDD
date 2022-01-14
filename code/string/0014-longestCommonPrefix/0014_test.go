package leetcode

import (
	pkgu "leetcode/pkg"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	test := GetTestData()
	for i := range test {
		t.Run("Round", func(t *testing.T) {
			got := longestCommonPrefix(test[i].TestData.([]string))
			want := test[i].Want.(string)

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}
		})
	}

}

func GetTestData() []pkgu.TestDataStruct {
	return []pkgu.TestDataStruct{
		{
			TestData: []string{"flower", "flow", "flight"},
			Want:     "fl",
		},
		{
			TestData: []string{"dog", "racecar", "car"},
			Want:     "",
		},
		{
			TestData: []string{"dogggg", "dog", "do"},
			Want:     "do",
		},
	}

}
