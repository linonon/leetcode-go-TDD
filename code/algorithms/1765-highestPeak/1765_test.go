package leetcode

import (
	"fmt"
	. "leetcode/pkg"
	"reflect"
	"testing"
)

func Test1765(t *testing.T) {
	for _, test := range Tests() {
		t.Run(test.Name(), func(t *testing.T) {
			isWater := test.Val.([][]int)

			got := highestPeak(isWater)
			want := test.Want.([][]int)

			if !reflect.DeepEqual(got, want) {
				t.Error("\ngot != errro")
				for i := range got {
					fmt.Println(got[i], "\t", want[i])
				}
			}
		})
	}
}

func Tests() []T {
	return []T{
		{
			Num: 1,
			Val: [][]int{
				{0, 1},
				{0, 0}},
			Want: [][]int{
				{1, 0},
				{2, 1}},
		},
		{
			Num: 2,
			Val: [][]int{
				{0, 0, 1},
				{1, 0, 0},
				{0, 0, 0}},
			Want: [][]int{
				{1, 1, 0},
				{0, 1, 1},
				{1, 2, 2}},
		},
		{
			Num: 3,
			Val: [][]int{
				{0, 1},
				{0, 1}},
			Want: [][]int{
				{1, 0},
				{1, 0}},
		},
		{
			Num: 4,
			Val: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0}},
			Want: [][]int{
				{2, 1, 2},
				{1, 0, 1},
				{2, 1, 2}},
		},
		{
			Num: 5,
			Val: [][]int{
				{0, 0},
				{1, 1},
				{1, 0},
			},
			Want: [][]int{
				{1, 1},
				{0, 0},
				{0, 1},
			},
		},
	}
}
