package leetcode

import (
	. "leetcode/pkg"
	"testing"
)

func Test720(t *testing.T) {
	for _, test := range Tests() {
		t.Run(test.Name(), func(t *testing.T) {
			words := test.Val.([]string)

			got := longestWord(words)
			want := test.Want.(string)

			if got != want {
				t.Errorf("got %v, want %v\n", got, want)
			}
		})
	}
}

func Tests() []T {
	return []T{
		{
			Num:  1,
			Val:  []string{"w", "wo", "wor", "worl", "world"},
			Want: "world",
		},
		{
			Num:  2,
			Val:  []string{"a", "banana", "app", "appl", "ap", "apply", "apple"},
			Want: "apple",
		},
		{
			Num:  3,
			Val:  []string{"m", "mo", "moc", "moch", "mocha", "l", "la", "lat", "latt", "latte", "c", "ca", "cat"},
			Want: "latte",
		},
		{
			Num:  4,
			Val:  []string{"t", "ti", "tig", "tige", "tiger", "e", "en", "eng", "engl", "engli", "englis", "english", "h", "hi", "his", "hist", "histo", "histor", "history"},
			Want: "english",
		},
		{
			Num:  5,
			Val:  []string{"yo", "ew", "fc", "zrc", "yodn", "fcm", "qm", "qmo", "fcmz", "z", "ewq", "yod", "ewqz", "y"},
			Want: "yodn",
		},
	}
}
