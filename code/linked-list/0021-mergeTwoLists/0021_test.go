package leetcode

import (
	"fmt"
	. "leetcode/pkg"
	"testing"
)

func Test0021(t *testing.T) {
	test := GetTestData()
	for _, v := range test {
		TestRun(t, v.Num, func(t *testing.T) {

			got := mergeTwoLists(v.Val1.(*ListNode), v.Val2.(*ListNode))
			want := v.Want.(*ListNode)

			if !got.IsEqual(want) {
				t.Error("got is not equal to want")
				fmt.Printf("Got:\t\t%v\n", got)
				fmt.Printf("Want:\t\t%v\n", want)
			}

			t.Logf("\nPass\n")
		})
	}
}

func GetTestData() []T {
	return []T{
		{
			Num:    1,
			Val1: ToNode([]int{1, 2, 4}),
			Val2: ToNode([]int{1, 3, 4}),
			Want:   ToNode([]int{1, 1, 2, 3, 4, 4}),
		},
		{
			Num:    2,
			Val1: ToNode([]int{}),
			Val2: ToNode([]int{}),
			Want:   ToNode([]int{}),
		},
		{
			Num:    3,
			Val1: ToNode([]int{1, 3, 4}),
			Val2: ToNode([]int{1, 2, 4}),
			Want:   ToNode([]int{1, 1, 2, 3, 4, 4}),
		},
		{
			Num:    4,
			Val1: ToNode([]int{1, 3, 4, 4}),
			Val2: ToNode([]int{1, 2, 4}),
			Want:   ToNode([]int{1, 1, 2, 3, 4, 4, 4}),
		},
		{
			Num:    5,
			Val1: ToNode([]int{1, 3, 4, 5}),
			Val2: ToNode([]int{1, 2, 4}),
			Want:   ToNode([]int{1, 1, 2, 3, 4, 4, 5}),
		},
		{
			Num:    6,
			Val1: ToNode([]int{1, 3}),
			Val2: ToNode([]int{1, 2, 4}),
			Want:   ToNode([]int{1, 1, 2, 3, 4}),
		},
		{
			Num:    7,
			Val1: ToNode([]int{0}),
			Val2: ToNode([]int{1, 2, 4}),
			Want:   ToNode([]int{0, 1, 2, 4}),
		},
		{
			Num:    8,
			Val1: ToNode([]int{5}),
			Val2: ToNode([]int{1, 2, 4}),
			Want:   ToNode([]int{1, 2, 4, 5}),
		},
		{
			Num:    9,
			Val1: ToNode([]int{1, 3}),
			Val2: ToNode([]int{5, 7}),
			Want:   ToNode([]int{1, 3, 5, 7}),
		},
	}
}
