package leetcode

import "testing"

func Test_containVirus(t *testing.T) {
	type args struct {
		isInfected [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				isInfected: [][]int{
					{0, 1, 0, 0, 0, 0, 0, 1},
					{0, 1, 0, 0, 0, 0, 0, 1},
					{0, 0, 0, 0, 0, 0, 0, 1},
					{0, 0, 0, 0, 0, 0, 0, 0},
				},
			},
			want: 10,
		},
		{
			args: args{
				isInfected: [][]int{
					{1, 1, 1},
					{1, 0, 1},
					{1, 1, 1},
				},
			},
			want: 4,
		},
		{
			args: args{
				isInfected: [][]int{
					{1, 1, 1, 0, 0, 0, 0, 0, 0},
					{1, 0, 1, 0, 1, 1, 1, 1, 1},
					{1, 1, 1, 0, 0, 0, 0, 0, 0},
				},
			},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containVirus(tt.args.isInfected); got != tt.want {
				t.Errorf("containVirus() = %v, want %v", got, tt.want)
			}
		})
	}
}
