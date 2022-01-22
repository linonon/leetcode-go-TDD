package leetcode

import (
	"fmt"
	. "leetcode/pkg"
	"testing"
)

func Test0021(t *testing.T) {
	test := GetTestData()
	for i := range test {
		str := fmt.Sprintf("Round-%v", test[i].Num)
		t.Run(str, func(t *testing.T) {

			got := mergeTwoLists(test[i].Value1.(*ListNode), test[i].Value2.(*ListNode))
			want := test[i].Want

			if !got.IsEqual(want.(*ListNode)) {
				t.Error("got is not equal to want")
				fmt.Printf("Got:\t\t%v\n", got.String())
				fmt.Printf("Want:\t\t%v\n", want.(*ListNode).String())
			}

			t.Logf("\nPass\n")
		})
	}
}

func GetTestData() []T {
	return []T{
		{
			Num:    1,
			Value1: ToNode([]int{1, 2, 4}),
			Value2: ToNode([]int{1, 3, 4}),
			Want:   ToNode([]int{1, 1, 2, 3, 4, 4}),
		},
		{
			Num:    2,
			Value1: ToNode([]int{}),
			Value2: ToNode([]int{}),
			Want:   ToNode([]int{}),
		},
		{
			Num:    3,
			Value1: ToNode([]int{1, 3, 4}),
			Value2: ToNode([]int{1, 2, 4}),
			Want:   ToNode([]int{1, 1, 2, 3, 4, 4}),
		},
		{
			Num:    4,
			Value1: ToNode([]int{1, 3, 4, 4}),
			Value2: ToNode([]int{1, 2, 4}),
			Want:   ToNode([]int{1, 1, 2, 3, 4, 4, 4}),
		},
		{
			Num:    5,
			Value1: ToNode([]int{1, 3, 4, 5}),
			Value2: ToNode([]int{1, 2, 4}),
			Want:   ToNode([]int{1, 1, 2, 3, 4, 4, 5}),
		},
		{
			Num:    6,
			Value1: ToNode([]int{1, 3}),
			Value2: ToNode([]int{1, 2, 4}),
			Want:   ToNode([]int{1, 1, 2, 3, 4}),
		},
		{
			Num:    7,
			Value1: ToNode([]int{0}),
			Value2: ToNode([]int{1, 2, 4}),
			Want:   ToNode([]int{0, 1, 2, 4}),
		},
		{
			Num:    8,
			Value1: ToNode([]int{5}),
			Value2: ToNode([]int{1, 2, 4}),
			Want:   ToNode([]int{1, 2, 4, 5}),
		},
		{
			Num:    9,
			Value1: ToNode([]int{1, 3}),
			Value2: ToNode([]int{5, 7}),
			Want:   ToNode([]int{1, 3, 5, 7}),
		},
	}
}
