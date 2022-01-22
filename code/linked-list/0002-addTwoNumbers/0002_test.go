package leetcode

import (
	. "leetcode/pkg"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	for _, v := range GetTestData() {
		TestRun(t, v.Num, func(t *testing.T) {

			got := addTwoNumbers(v.Value1.(*ListNode), v.Value2.(*ListNode))
			want := v.Want.(*ListNode)

			if !got.IsEqual(want) {
				t.Errorf("addTwoNumbers: got %v, want %v\n", got, want)
			}

		})
	}
}

func GetTestData() []T {
	return []T{
		{
			Num:    0,
			Value1: ToNode([]int{1, 4, 3}),
			Value2: ToNode([]int{5, 6, 4}),
			Want:   ToNode([]int{6, 0, 8}),
		},
		{
			Num:    1,
			Value1: ToNode([]int{0}),
			Value2: ToNode([]int{0}),
			Want:   ToNode([]int{0}),
		},
		{
			Num:    2,
			Value1: ToNode([]int{9}),
			Value2: ToNode([]int{1}),
			Want:   ToNode([]int{0, 1}),
		},
		{
			Num:    3,
			Value1: ToNode([]int{9, 9}),
			Value2: ToNode([]int{1}),
			Want:   ToNode([]int{0, 0, 1}),
		},
		{
			Num:    4,
			Value1: ToNode([]int{9, 9, 1}),
			Value2: ToNode([]int{1}),
			Want:   ToNode([]int{0, 0, 2}),
		},
		{
			Num:    5,
			Value1: ToNode([]int{9, 9, 1}),
			Value2: ToNode([]int{1, 0, 1}),
			Want:   ToNode([]int{0, 0, 3}),
		},
		{
			Num:    6,
			Value1: ToNode([]int{9, 9, 9, 9, 9, 9, 9}),
			Value2: ToNode([]int{9, 9, 9, 9}),
			Want:   ToNode([]int{8, 9, 9, 9, 0, 0, 0, 1}),
		},
		{
			Num:    7,
			Value1: ToNode([]int{1, 8}),
			Value2: ToNode([]int{0}),
			Want:   ToNode([]int{1, 8}),
		},
		{
			Num:    8,
			Value1: ToNode([]int{1, 8, 9}),
			Value2: ToNode([]int{0}),
			Want:   ToNode([]int{1, 8, 9}),
		},
		{
			Num:    9,
			Value1: ToNode([]int{9}),
			Value2: ToNode([]int{9}),
			Want:   ToNode([]int{8, 1}),
		},
		{
			Num:    10,
			Value1: ToNode([]int{1, 8, 3}),
			Value2: ToNode([]int{7, 1}),
			Want:   ToNode([]int{8, 9, 3}),
		},
		{
			Num:    11,
			Value1: ToNode([]int{9}),
			Value2: ToNode([]int{1, 9, 9, 9, 9, 9, 8, 9, 9, 9}),
			Want:   ToNode([]int{0, 0, 0, 0, 0, 0, 9, 9, 9, 9}),
		},
	}
}
