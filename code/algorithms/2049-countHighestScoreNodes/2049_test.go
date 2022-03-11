package leetcode

import (
	. "leetcode/pkg"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test2049(t *testing.T) {
	for _, test := range Tests() {
		t.Run(test.Name(), func(t *testing.T) {
			parents := test.Val.([]int)

			assert.EqualValues(t,
				test.Want,
				countHighestScoreNodes(parents),
			)
		})
	}

}

func Tests() []T {
	return []T{
		{
			Num:  1,
			Val:  []int{-1, 2, 0, 2, 0},
			Want: 3,
		},
	}
}
