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

			got := mergeTwoLists(test[i].ListNode1, test[i].ListNode2)
			want := test[i].Want

			if !got.IsEqual(want) {
				t.Error("got is not equal to want")
				fmt.Printf("Got:\t\t%v\n", got.String())
				fmt.Printf("Want:\t\t%v\n", want.String())
			}

			t.Logf("\nPass\n")
		})
	}
}

func GetTestData() []pkgu.TestListNodeStruct {
	return []pkgu.TestListNodeStruct{
		{
			Num:       1,
			ListNode1: pkgu.ToNode([]int{1, 2, 4}),
			ListNode2: pkgu.ToNode([]int{1, 3, 4}),
			Want:      pkgu.ToNode([]int{1, 1, 2, 3, 4, 4}),
		},
		{
			Num:       2,
			ListNode1: pkgu.ToNode([]int{}),
			ListNode2: pkgu.ToNode([]int{}),
			Want:      pkgu.ToNode([]int{}),
		},
		{
			Num:       3,
			ListNode1: pkgu.ToNode([]int{1, 3, 4}),
			ListNode2: pkgu.ToNode([]int{1, 2, 4}),
			Want:      pkgu.ToNode([]int{1, 1, 2, 3, 4, 4}),
		},
		{
			Num:       4,
			ListNode1: pkgu.ToNode([]int{1, 3, 4, 4}),
			ListNode2: pkgu.ToNode([]int{1, 2, 4}),
			Want:      pkgu.ToNode([]int{1, 1, 2, 3, 4, 4, 4}),
		},
		{
			Num:       5,
			ListNode1: pkgu.ToNode([]int{1, 3, 4, 5}),
			ListNode2: pkgu.ToNode([]int{1, 2, 4}),
			Want:      pkgu.ToNode([]int{1, 1, 2, 3, 4, 4, 5}),
		},
		{
			Num:       6,
			ListNode1: pkgu.ToNode([]int{1, 3}),
			ListNode2: pkgu.ToNode([]int{1, 2, 4}),
			Want:      pkgu.ToNode([]int{1, 1, 2, 3, 4}),
		},
		{
			Num:       7,
			ListNode1: pkgu.ToNode([]int{0}),
			ListNode2: pkgu.ToNode([]int{1, 2, 4}),
			Want:      pkgu.ToNode([]int{0, 1, 2, 4}),
		},
		{
			Num:       8,
			ListNode1: pkgu.ToNode([]int{5}),
			ListNode2: pkgu.ToNode([]int{1, 2, 4}),
			Want:      pkgu.ToNode([]int{1, 2, 4, 5}),
		},
		{
			Num:       9,
			ListNode1: pkgu.ToNode([]int{1, 3}),
			ListNode2: pkgu.ToNode([]int{5, 7}),
			Want:      pkgu.ToNode([]int{1, 3, 5, 7}),
		},
	}
}
