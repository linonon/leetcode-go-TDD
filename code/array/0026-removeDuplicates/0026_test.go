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
			got := removeDuplicates(v.TestData.([]int))
			want := v.Want.([]int)
			want2 := v.Want2.(int)

			if !reflect.DeepEqual(v.TestData.([]int), want) && want2 != got {
				t.Errorf("got %v, want %v", got, want2)
			}
		})
	}

}

func GetTestData() []pkgu.TestDataStruct {
	return []pkgu.TestDataStruct{
		{
			Num:      1,
			TestData: []int{1, 1, 2},
			Want:     []int{1, 2},
			Want2:    2,
		},
		{
			Num:      2,
			TestData: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			Want:     []int{0, 1, 2, 3, 4},
			Want2:    5,
		},
	}

}
