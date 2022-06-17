package leetcode

import (
	"fmt"
	"testing"
)

func Test_duplicateZeros(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{
				arr: []int{1, 0, 2, 3, 0, 4, 5, 0},
			},
		},
		{
			name: "2",
			args: args{
				arr: []int{1, 0, 2, 3, 0, 4, 5, 0, 0},
			},
		},
		{
			name: "3",
			args: args{
				arr: []int{1, 0, 0, 2, 3, 0, 4, 5, 0, 0},
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(tt.args.arr)
			duplicateZeros(tt.args.arr)
			fmt.Println(tt.args.arr)
		})
	}
}
