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
			got := removeElement(v.TestData.([]int), v.TestData2.(int))
			want := v.Want.(int)
			if got != want {
				t.Errorf("got %v, want %v\n", got, want)

				fmt.Printf("[]int:")
				for i := 0; i < got; i++ {
					fmt.Printf(" %v ", v.TestData.([]int)[i])
				}
				fmt.Printf("\n")
				fmt.Println(v.TestData.([]int))
			}
		})
	}
}

func GetTestData() []pkgu.TestDataStruct {
	return []pkgu.TestDataStruct{
		{
			Num:       1,
			TestData:  []int{3, 2, 2, 3},
			TestData2: 3,
			Want:      2,
		},
		{
			Num:       2,
			TestData:  []int{0, 1, 2, 2, 3, 0, 4, 2},
			TestData2: 2,
			Want:      5,
		},
		{
			Num:       3,
			TestData:  []int{1},
			TestData2: 1,
			Want:      0,
		},
		{
			Num:       4,
			TestData:  []int{4, 5},
			TestData2: 4,
			Want:      1,
		},
	}
}
