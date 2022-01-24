package pkgu

import (
	"fmt"
	"testing"
)

var test = GetTestData()

func TestNodesPrint(t *testing.T) {
	for _, v := range test {
		TestRun(t, v.Num, func(t *testing.T) {
			fmt.Println(v.Val1.(*ListNode))
		})
	}
}

func TestNodesLength(t *testing.T) {
	for _, v := range test {
		TestRun(t, v.Num, func(t *testing.T) {

			length, err := v.Val1.(*ListNode).Length()
			if err != nil {
				fmt.Printf("Err: %+v\n", err)
				return
			}
			fmt.Println("len(node) ==", length)
		})
	}
}

func GetTestData() []T {
	return []T{
		{
			Num:    1,
			Val1: ToNode([]int{1, 2, 4}),
			Val2: ToNode([]int{1, 3, 4}),
		},
		{
			Num:    2,
			Val1: ToNode([]int{}),
			Val2: ToNode([]int{}),
		},
		{
			Num:    3,
			Val1: RingNode([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}),
		},
		{
			Num:    4,
			Val1: ToNode([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}),
		},
	}
}
