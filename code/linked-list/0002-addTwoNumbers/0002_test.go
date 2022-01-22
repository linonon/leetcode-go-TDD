package leetcode

import (
	"fmt"
	pkgu "leetcode/pkg"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	test := GetTestData()
	for _, v := range test {
		str := fmt.Sprintf("Test %v", v.Num)
		t.Run(str, func(t *testing.T) {
			got := addTwoNumbers(v.ListNode1, v.ListNode2)
			want := v.Want
			if ok := got.IsEqual(want); !ok {
				t.Errorf("addTwoNumbers: got %v, want %v\n", got.String(), want.String())
			}
		})
	}
}

func GetTestData() []pkgu.TestListNodeStruct {
	return []pkgu.TestListNodeStruct{
		{
			Num:       0,
			ListNode1: pkgu.ToNode([]int{1, 4, 3}),
			ListNode2: pkgu.ToNode([]int{5, 6, 4}),
			Want:      pkgu.ToNode([]int{6, 0, 8}),
		},
		{
			Num:       1,
			ListNode1: pkgu.ToNode([]int{0}),
			ListNode2: pkgu.ToNode([]int{0}),
			Want:      pkgu.ToNode([]int{0}),
		},
		{
			Num:       2,
			ListNode1: pkgu.ToNode([]int{9}),
			ListNode2: pkgu.ToNode([]int{1}),
			Want:      pkgu.ToNode([]int{0, 1}),
		},
		{
			Num:       3,
			ListNode1: pkgu.ToNode([]int{9, 9}),
			ListNode2: pkgu.ToNode([]int{1}),
			Want:      pkgu.ToNode([]int{0, 0, 1}),
		},
		{
			Num:       4,
			ListNode1: pkgu.ToNode([]int{9, 9, 1}),
			ListNode2: pkgu.ToNode([]int{1}),
			Want:      pkgu.ToNode([]int{0, 0, 2}),
		},
		{
			Num:       5,
			ListNode1: pkgu.ToNode([]int{9, 9, 1}),
			ListNode2: pkgu.ToNode([]int{1, 0, 1}),
			Want:      pkgu.ToNode([]int{0, 0, 3}),
		},
		{
			Num:       6,
			ListNode1: pkgu.ToNode([]int{9, 9, 9, 9, 9, 9, 9}),
			ListNode2: pkgu.ToNode([]int{9, 9, 9, 9}),
			Want:      pkgu.ToNode([]int{8, 9, 9, 9, 0, 0, 0, 1}),
		},
		{
			Num:       7,
			ListNode1: pkgu.ToNode([]int{1, 8}),
			ListNode2: pkgu.ToNode([]int{0}),
			Want:      pkgu.ToNode([]int{1, 8}),
		},
		{
			Num:       8,
			ListNode1: pkgu.ToNode([]int{1, 8, 9}),
			ListNode2: pkgu.ToNode([]int{0}),
			Want:      pkgu.ToNode([]int{1, 8, 9}),
		},
		{
			Num:       9,
			ListNode1: pkgu.ToNode([]int{9}),
			ListNode2: pkgu.ToNode([]int{9}),
			Want:      pkgu.ToNode([]int{8, 1}),
		},
		{
			Num:       10,
			ListNode1: pkgu.ToNode([]int{1, 8, 3}),
			ListNode2: pkgu.ToNode([]int{7, 1}),
			Want:      pkgu.ToNode([]int{8, 9, 3}),
		},
		{
			Num:       11,
			ListNode1: pkgu.ToNode([]int{9}),
			ListNode2: pkgu.ToNode([]int{1, 9, 9, 9, 9, 9, 8, 9, 9, 9}),
			Want:      pkgu.ToNode([]int{0, 0, 0, 0, 0, 0, 9, 9, 9, 9}),
		},
	}
}
