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
			got := removeElement(v.Val1.([]int), v.Val2.(int))
			want := v.Want.(int)
			if got != want {
				t.Errorf("got %v, want %v\n", got, want)

				fmt.Printf("[]int:")
				for i := 0; i < got; i++ {
					fmt.Printf(" %v ", v.Val1.([]int)[i])
				}
				fmt.Printf("\n")
				fmt.Println(v.Val1.([]int))
			}
		})
	}
}

func GetTestData() []pkgu.T {
	return []pkgu.T{
		{
			Num:  1,
			Val1: []int{3, 2, 2, 3},
			Val2: 3,
			Want: 2,
		},
		{
			Num:  2,
			Val1: []int{0, 1, 2, 2, 3, 0, 4, 2},
			Val2: 2,
			Want: 5,
		},
		{
			Num:  3,
			Val1: []int{1},
			Val2: 1,
			Want: 0,
		},
		{
			Num:  4,
			Val1: []int{4, 5},
			Val2: 4,
			Want: 1,
		},
	}
}
