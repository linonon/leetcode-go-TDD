package leetcode

import "testing"

func Test_maxConsecutiveAnswers(t *testing.T) {
	type args struct {
		answerKey string
		k         int
	}
	tests := []struct {
		name       string
		args       args
		wantResult int
	}{
		{
			name:       "1",
			args:       args{answerKey: "TTFF", k: 2},
			wantResult: 4,
		},
		{
			name:       "2",
			args:       args{answerKey: "TFFT", k: 1},
			wantResult: 3,
		},
		{
			name:       "3",
			args:       args{answerKey: "TTFTTFTT", k: 1},
			wantResult: 5,
		},
		{
			name:       "4",
			args:       args{answerKey: "TFTFTTTFT", k: 1},
			wantResult: 3,
		},
		{
			name:       "5",
			args:       args{answerKey: "FFFTTFTTFT", k: 1},
			wantResult: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := maxConsecutiveAnswers(tt.args.answerKey, tt.args.k); gotResult != tt.wantResult {
				t.Errorf("maxConsecutiveAnswers() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
