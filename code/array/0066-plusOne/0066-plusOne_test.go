package leetcode

import (
	"testing"
)

func TestPlusOne(t *testing.T) {
	in := [][]int{
		{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6},
		{1, 2, 3},
		{9, 9, 9},
	}

	for i := range in {
		t.Run("Round", func(t *testing.T) {
			result := plusOne(in[i])
			t.Log("result:", result)
		})
	}
}
