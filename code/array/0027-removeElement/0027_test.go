package leetcode

import (
	"fmt"
	pkgu "leetcode/pkg"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	tests := GetTestData()
	for _, v := range tests {
		str := fmt.Sprintf("Test %v", v.Num)
		t.Run(str, func(t *testing.T) {
			got := removeElement(v.Value1.([]int), v.Value2.(int))
			want := v.Want1.(int)
			if got != want {
				t.Errorf("got %v, want %v\n", got, want)

				fmt.Printf("[]int:")
				for i := 0; i < got; i++ {
					fmt.Printf(" %v ", v.Value1.([]int)[i])
				}
				fmt.Printf("\n")
				fmt.Println(v.Value1.([]int))
			}
		})
	}
}

func GetTestData() []pkgu.T {
	return []pkgu.T{
		{
			Num:       1,
			Value1:  []int{3, 2, 2, 3},
			Value2: 3,
			Want1:      2,
		},
		{
			Num:       2,
			Value1:  []int{0, 1, 2, 2, 3, 0, 4, 2},
			Value2: 2,
			Want1:      5,
		},
		{
			Num:       3,
			Value1:  []int{1},
			Value2: 1,
			Want1:      0,
		},
		{
			Num:       4,
			Value1:  []int{4, 5},
			Value2: 4,
			Want1:      1,
		},
	}
}
