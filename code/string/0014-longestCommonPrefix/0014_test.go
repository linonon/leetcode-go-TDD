package leetcode

import (
	pkgu "leetcode/pkg"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	test := GetTestData()
	for i := range test {
		t.Run("Round", func(t *testing.T) {
			got := longestCommonPrefix(test[i].Value1.([]string))
			want := test[i].Want1.(string)

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}
		})
	}

}

func GetTestData() []pkgu.T {
	return []pkgu.T{
		{
			Value1: []string{"flower", "flow", "flight"},
			Want1:     "fl",
		},
		{
			Value1: []string{"dog", "racecar", "car"},
			Want1:     "",
		},
		{
			Value1: []string{"dogggg", "dog", "do"},
			Want1:     "do",
		},
	}

}
