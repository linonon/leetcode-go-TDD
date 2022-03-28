package leetcode

import "testing"

func Test_hasAlternatingBits(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		wantHas bool
	}{
		// TODO: Add test cases.
		{
			name:    "1",
			args:    args{n: 5},
			wantHas: true,
		},
		{
			name:    "2",
			args:    args{n: 7},
			wantHas: false,
		},
		{
			name:    "3",
			args:    args{n: 11},
			wantHas: false,
		},
		{
			name:    "4",
			args:    args{n: 13},
			wantHas: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotHas := hasAlternatingBits(tt.args.n); gotHas != tt.wantHas {
				t.Errorf("hasAlternatingBits() = %v, want %v", gotHas, tt.wantHas)
			}
		})
	}
}
