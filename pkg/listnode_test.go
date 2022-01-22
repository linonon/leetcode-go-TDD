package pkgu

import (
	"fmt"
	"testing"
)

var test = GetTestData()

func TestNodesPrint(t *testing.T) {
	for i := range test {
		str := fmt.Sprintf("Round-%v", test[i].Num)
		t.Run(str, func(t *testing.T) {
			fmt.Println(test[i].Value1.(*ListNode).String())
		})
	}
}
func TestNodesLength(t *testing.T) {
	for i := range test {
		str := fmt.Sprintf("Round-%v", test[i].Num)
		t.Run(str, func(t *testing.T) {

			length, err := test[i].Value1.(*ListNode).Length()
			if err != nil {
				fmt.Println(err)
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
			Value1: ToNode([]int{1, 2, 4}),
			Value2: ToNode([]int{1, 3, 4}),
		},
		{
			Num:    2,
			Value1: ToNode([]int{}),
			Value2: ToNode([]int{}),
		},
		{
			Num:    3,
			Value1: RingNode([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}),
		},
		{
			Num:    4,
			Value1: ToNode([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}),
		},
	}
}
