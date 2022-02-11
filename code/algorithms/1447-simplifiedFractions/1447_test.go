package leetcode

import (
	. "leetcode/pkg"
	"reflect"
	"testing"
)

func Test1447(t *testing.T) {
	for _, test := range Tests() {
		t.Run(test.Name(), func(t *testing.T) {
			n := test.Val.(int)

			got := simplifiedFractions(n)
			want := test.Want.([]string)

			// HACK: []string != []string{}
			if !reflect.DeepEqual(got, want) {
				t.Errorf("got: %v, want: %v\n", got, want)
				t.FailNow()
			}
		})
	}
}

func Tests() []T {
	return []T{
		{
			Num:  1,
			Val:  2,
			Want: []string{"1/2"},
		},
		{
			Num:  2,
			Val:  3,
			Want: []string{"1/2", "1/3", "2/3"},
		},
		{
			Num:  3,
			Val:  4,
			Want: []string{"1/2", "1/3", "1/4", "2/3", "3/4"},
		},
		{
			Num:  4,
			Val:  1,
			Want: []string{},
		},
		{
			Num:  5,
			Val:  6,
			Want: []string{"1/2", "1/3", "1/4", "1/5", "1/6", "2/3", "2/5", "3/4", "3/5", "4/5", "5/6"},
		},
		{
			Num:  6,
			Val:  8,
			Want: []string{"1/2", "1/3", "1/4", "1/5", "1/6", "1/7", "1/8", "2/3", "2/5", "2/7", "3/4", "3/5", "3/7", "3/8", "4/5", "4/7", "5/6", "5/7", "5/8", "6/7", "7/8"},
		},
	}
}
