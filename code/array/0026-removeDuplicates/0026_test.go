package leetcode

import (
	"fmt"
	pkgu "leetcode/pkg"
	"reflect"
	"testing"
)

func TestRemoveDuplicate(t *testing.T) {
	tests := GetTestData()
	for _, v := range tests {
		str := fmt.Sprintf("Test %v", v.Num)
		t.Run(str, func(t *testing.T) {
			got := removeDuplicates(v.Value1.([]int))
			want := v.Want1.([]int)
			want2 := v.Want2.(int)

			if !reflect.DeepEqual(v.Value1.([]int), want) && want2 != got {
				t.Errorf("got %v, want %v", got, want2)
			}
		})
	}

}

func GetTestData() []pkgu.T {
	return []pkgu.T{
		{
			Num:      1,
			Value1: []int{1, 1, 2},
			Want1:     []int{1, 2},
			Want2:    2,
		},
		{
			Num:      2,
			Value1: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			Want1:     []int{0, 1, 2, 3, 4},
			Want2:    5,
		},
	}

}
