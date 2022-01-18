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
			fmt.Println(test[i].ListNode1.String())
		})
	}
}
func TestNodesLength(t *testing.T) {
	for i := range test {
		str := fmt.Sprintf("Round-%v", test[i].Num)
		t.Run(str, func(t *testing.T) {

			length, err := test[i].ListNode1.Length()
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println("len(node) ==", length)
		})
	}
}

func GetTestData() []TestListNodeStruct {
	return []TestListNodeStruct{
		{
			Num:       1,
			ListNode1: Node([]int{1, 2, 4}),
			ListNode2: Node([]int{1, 3, 4}),
		},
		{
			Num:       2,
			ListNode1: Node([]int{}),
			ListNode2: Node([]int{}),
		},
		{
			Num:       3,
			ListNode1: RingNode([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}),
		},
		{
			Num:       4,
			ListNode1: Node([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}),
		},
	}
}
