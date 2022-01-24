package leetcode

import (
	pkgu "leetcode/pkg"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	test := GetTestData()
	for i := range test {
		t.Run("Round", func(t *testing.T) {
			got := longestCommonPrefix(test[i].Val1.([]string))
			want := test[i].Want.(string)

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}
		})
	}

}

func GetTestData() []pkgu.T {
	return []pkgu.T{
		{
			Val1: []string{"flower", "flow", "flight"},
			Want:     "fl",
		},
		{
			Val1: []string{"dog", "racecar", "car"},
			Want:     "",
		},
		{
			Val1: []string{"dogggg", "dog", "do"},
			Want:     "do",
		},
	}

}
