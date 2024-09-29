package service

import (
	"testing"

	"github.com/guhkun13/gmtg/config"
)

func TestIsMatchAssignValue(t *testing.T) {
	regexp := config.RegexAssignCurrency
	srv := NewCurrencyImpl(regexp)

	testCases := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "positive case 1",
			input: "glob is I",
			want:  true,
		},
		{
			name:  "positive case 2",
			input: "prok is V",
			want:  true,
		},
		{
			name:  "positive case 3",
			input: "pish is X",
			want:  true,
		},
		{
			name:  "positive case 4",
			input: "tegj is L",
			want:  true,
		},
		{
			name:  "negative case 1",
			input: "satu was I",
			want:  false,
		},
		{
			name:  "negative case 2",
			input: "dua is P",
			want:  false,
		},
		{
			name:  "negative case 3",
			input: "what is this",
			want:  false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := srv.IsMatchAssignValue(tt.input)
			if got != tt.want {
				t.Errorf("IsMatchAssignValue [%v], got %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
