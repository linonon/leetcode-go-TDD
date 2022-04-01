package leetcode

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
	}{
		{
			name:       "1",
			args:       args{s: "babad"},
			wantResult: "bab",
		},
		{
			name:       "2",
			args:       args{s: "sbbd"},
			wantResult: "bb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := longestPalindrome(tt.args.s); gotResult != tt.wantResult {
				t.Errorf("longestPalindrome() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
