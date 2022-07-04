package leetcode

import (
	"reflect"
	"testing"
)

func Test_minimumAbsDifference(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name       string
		args       args
		wantResult [][]int
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				arr: []int{4, 2, 1, 3},
			},
			wantResult: [][]int{{1, 2}, {2, 3}, {3, 4}},
		},
		{
			name: "2",
			args: args{
				arr: []int{1, 3, 6, 10, 15},
			},
			wantResult: [][]int{{1, 3}},
		},
		{
			name: "3",
			args: args{
				arr: []int{-57, -70, 6, -69, 5, -47, -13},
			},
			wantResult: [][]int{{-70, -69}, {5, 6}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := minimumAbsDifference(tt.args.arr); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("minimumAbsDifference() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
