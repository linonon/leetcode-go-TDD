package leetcode

import (
	"fmt"
	. "leetcode/pkg"
	"reflect"
	"testing"
)

func TestRemoveDuplicate(t *testing.T) {
	tests := GetTestData()
	for _, v := range tests {
		str := fmt.Sprintf("Test %v", v.Num)
		t.Run(str, func(t *testing.T) {
			got := removeDuplicates(v.Val.([]int))
			want := v.W[0].([]int)
			want1 := v.W[1].(int)

			if !reflect.DeepEqual(v.Val.([]int), want) && want1 != got {
				t.Errorf("got %v, want %v", got, want1)
			}
		})
	}
}

func GetTestData() []T {
	return []T{
		{
			Num: 1,
			Val:   []int{1, 1, 2},
			W: []interface{}{
				[]int{1, 2, 3},
				2,
			},
		},
		{
			Num: 2,
			Val:   []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			W: []interface{}{
				[]int{0, 1, 2, 3, 4},
				5,
			},
		},
	}
}
