package leetcode

import (
	"fmt"
	pkgu "leetcode/pkg"
	"testing"
)

func Test0021(t *testing.T) {
	test := GetTestData()
	for i := range test {
		str := fmt.Sprintf("Round-%v", test[i].Num)
		t.Run(str, func(t *testing.T) {

			mergeTwoLists(test[i].ListNode1, test[i].ListNode1)

		})
	}
}

func GetTestData() []pkgu.TestListNodeStruct {
	return []pkgu.TestListNodeStruct{
		{
			Num:       1,
			ListNode1: pkgu.IntSliceToListNode([]int{1, 2, 4}),
			ListNode2: pkgu.IntSliceToListNode([]int{1, 3, 4}),
		},
		{
			Num:       2,
			ListNode1: pkgu.IntSliceToListNode([]int{}),
			ListNode2: pkgu.IntSliceToListNode([]int{}),
		},
	}
}
